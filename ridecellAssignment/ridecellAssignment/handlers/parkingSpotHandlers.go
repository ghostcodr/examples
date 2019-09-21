package handlers

import (
	"example.com/ridecellAssignment/handlers/responses"
	"example.com/ridecellAssignment/localdb"
	"example.com/ridecellAssignment/localdb/entities"
	"net/http"
	"strings"
)

func HandleGetParkingSpots(cityId, streetId, status string) ([]map[string]*entities.ParkingSpot, *entities.Error) {
	if strings.TrimSpace(cityId) == "" || strings.TrimSpace(streetId) == "" {
		return nil, &entities.Error{
			Code:     "PL-3",
			HTTPCode: http.StatusNotFound,
			Message:  "required parameter is missing",
		}
	}
	spots, err := localdb.GetInstance().GetParkingSpots(cityId, streetId, status)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(status) != "" {
		res := make([]map[string]*entities.ParkingSpot, 0)
		for _, parkingSpots := range spots {
			for _, spot := range parkingSpots {
				if spot.BookingStatus.Status == entities.GetStatus(status) {
					res = append(res, map[string]*entities.ParkingSpot{spot.Id: spot})
				}
			}
		}
		return res, nil
	}
	return spots, nil
}

func HandleGetNearByParkingSpots(cityId, streetId string, lat, lng, radius float64) ([]map[string]*entities.ParkingSpot, *entities.Error) {
	spots, err := localdb.GetInstance().GetParkingSpotsByLatLngWithIn(cityId, streetId, lat, lng, radius, entities.AVAILABLE)
	if err != nil {
		return nil, err
	}
	return spots, nil
}

func HandleReserveParkingSpot(cityId, streetId, parkingSpotId string) (*responses.ReserveParkingSpotResponse, *entities.Error) {
	/*
		validations
		1. spot should be available before booking.
		2. concurrency handling Lock the parking spot.
		3. Keep reservation and payment in one transaction.
		4. ??Cancel booking if payment failed
	*/
	db := localdb.GetInstance()
	db.AcquireLock()
	defer db.ReleaseLock()

	err := db.ReserveParking(cityId, streetId, parkingSpotId)
	if err != nil {
		return nil, err
	}
	return &responses.ReserveParkingSpotResponse{
		Message: "Successfully booked the parking",
	}, nil
}

func HandleCancelParkingSpot(cityId, streetId, parkingSpotId string) (*responses.CancelParkingSpotResponse, *entities.Error) {
	db := localdb.GetInstance()
	db.AcquireLock()
	defer db.ReleaseLock()

	err := db.CancelParking(cityId, streetId, parkingSpotId)
	if err != nil {
		return nil, err
	}
	return &responses.CancelParkingSpotResponse{
		Message: "Successfully canceled the booking",
	}, nil
}
