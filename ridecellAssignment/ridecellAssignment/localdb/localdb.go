package localdb

import (
	"example.com/ridecellAssignment/localdb/entities"
	"net/http"
	"sync"
	"os/user"
)

type db struct {
	parkingLocations map[string]*entities.City
	plans map[string]*entities.Plan
	users map[string]*user.User
	mux              sync.Mutex
}

var dbInstance *db
var once sync.Once

func GetInstance() *db {
	once.Do(func() {
		dbInstance = &db{}
		dbInstance.parkingLocations = make(map[string]*entities.City)
		dbInstance.plans=make(map[string]*entities.Plan)
		plan := &entities.Plan{Id: "per_hour", Name: "Per hour plan", Rate: 20, Unit: 1}
		dbInstance.plans["per_hour"]= plan
		city := &entities.City{
			Id:   "PN001",
			Name: "Pune",
			Streets: []map[string]*entities.Street{
				{
					"MG001": {
						Id:   "MG001",
						Name: "MG Road",
						ParkingSpots: []map[string]*entities.ParkingSpot{
							{
								"spot001": {
									Id:       "spot001",
									Type:     entities.SMALL,
									Lat:      18.520430,
									Lng:      73.856743,
									Rate:     "per_hour",
									Currency: entities.RS,
									BookingStatus: entities.BookingStatus{
										Status: entities.AVAILABLE,
									},
								},
								"spot002": {
									Id:       "spot002",
									Type:     entities.SMALL,
									Lat:      18.530430,
									Lng:      73.866743,
									Rate:     "per_hour",
									Currency: entities.RS,
									BookingStatus: entities.BookingStatus{
										Status: entities.AVAILABLE,
									},
								},
								"spot003": {
									Id:       "spot003",
									Type:     entities.SMALL,
									Lat:      18.540430,
									Lng:      73.876743,
									Rate:     "per_hour",
									Currency: entities.RS,
									BookingStatus: entities.BookingStatus{
										Status: entities.AVAILABLE,
									},
								},
								"spot004": {
									Id:       "spot004",
									Type:     entities.LARGE,
									Lat:      18.550430,
									Lng:      73.886743,
									Rate:     "per_hour",
									Currency: entities.RS,
									BookingStatus: entities.BookingStatus{
										Status: entities.AVAILABLE,
									},
								},
							},
						},
					},
				},
			},
		}
		dbInstance.parkingLocations[city.Id] = city
	})
	return dbInstance
}

func (dbInst *db) GetParkingSpots(cityId, streetId,status string) ([]map[string]*entities.ParkingSpot, *entities.Error) {
	if city, ok := dbInst.parkingLocations[cityId]; ok {
		for _, streets := range city.Streets {
			if street, streetOk := streets[streetId]; streetOk {
				return street.ParkingSpots, nil //FIXME handle status if empty return all spots
			}
		}
		// if street not found return error
		return nil, &entities.Error{
			Code:     "PL-2", //FIXME: externalize to constants file
			Message:  "Street Not found",
			HTTPCode: http.StatusNotFound,
		}

	} else {
		// if city not found return error
		return nil, &entities.Error{
			Code:     "PL-1", //FIXME: externalize to constants file
			Message:  "City Not found",
			HTTPCode: http.StatusNotFound,
		}
	}
}

func (dbInst *db) GetParkingSpotsByLatLngWithIn(cityId, streetId string, lat, lng, radius float64, status entities.Status) ([]map[string]*entities.ParkingSpot, *entities.Error) {
	spots, err := dbInst.GetParkingSpots(cityId, streetId,"")
	if err != nil {
		return nil, err
	}
	//FIXME needs to calculate the spots according to the radius and lat, lng
	return spots, nil
}

func (dbInst *db) IsParkingSpotAvailable(cityId, streetId, parkingSpotId string) (bool, *entities.Error) {
	spot, err := dbInst.getParkingSpot(cityId, streetId, parkingSpotId)
	if err != nil {
		return false, err
	}
	if spot.BookingStatus.Status == entities.AVAILABLE {
		return true, nil
	}
	return false, nil
}

func (dbInst *db) getParkingSpot(cityId, streetId, parkingSpotId string) (*entities.ParkingSpot, *entities.Error) {
	spots, err := dbInst.GetParkingSpots(cityId, streetId,"")
	if err != nil {
		return nil, err
	}
	for _, parkingSpot := range spots {
		if spot, ok := parkingSpot[parkingSpotId]; ok {
			return spot, nil
		}
	}
	return nil, &entities.Error{
		Code:     "PS-1", //FIXME: externalize to constants file
		Message:  "Parking spot Not found",
		HTTPCode: http.StatusNotFound,
	}
}
func (dbInst *db) ReserveParking(cityId, streetId, parkingSpotId string) ( *entities.Error) {
	spot, err := dbInst.getParkingSpot(cityId, streetId, parkingSpotId)
	if err != nil {
		return err
	}
	if spot.BookingStatus.Status != entities.AVAILABLE {
		return  &entities.Error{
			Code:     "PS-11",
			Message:  "Parking spot is not available",
			HTTPCode: http.StatusPreconditionFailed,
		}
	}
	spot.BookingStatus.Status=entities.OCCUPIED
	return nil
}

func (dbInst *db) CancelParking(cityId, streetId, parkingSpotId string) ( *entities.Error) {
	spot, err := dbInst.getParkingSpot(cityId, streetId, parkingSpotId)
	if err != nil {
		return err
	}
	if spot.BookingStatus.Status != entities.OCCUPIED {
		return  &entities.Error{
			Code:     "PS-11",
			Message:  "Parking spot is already available",
			HTTPCode: http.StatusPreconditionFailed,
		}
	}
	spot.BookingStatus.Status=entities.AVAILABLE
	return nil
}

func (dbInst *db) AcquireLock() {
	dbInstance.mux.Lock()
}

func (dbInst *db) ReleaseLock() {
	dbInstance.mux.Unlock()
}
