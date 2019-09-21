package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"example.com/ridecellAssignment/localdb/entities"
)

func TestHandleGetParkingSpots(t *testing.T) {
	spots, err := HandleGetParkingSpots("PN001", "MG001","")
	assert.Nil(t,err,"Error should be nil")
	b, e := json.Marshal(spots)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	assert.Contains(t,spots[0],"spot001","Response should contain cities")
}

func TestHandleGetParkingSpotsInvalidCity(t *testing.T) {
	_, err := HandleGetParkingSpots("invalid", "MG001","")
	assert.NotNil(t,err,"Error should be not nil")
	assert.Equal(t,err.Code,"PL-1")
	assert.Equal(t,err.HTTPCode,404)

}

func TestHandleGetParkingSpotsInvalidStreet(t *testing.T) {
	_, err := HandleGetParkingSpots("PN001", "invalid","")
	assert.NotNil(t,err,"Error should be not nil")
	assert.Equal(t,err.Code,"PL-2")
	assert.Equal(t,err.HTTPCode,404)

}
func TestHandleGetNearByParkingSpots(t *testing.T) {
	spots, err := HandleGetNearByParkingSpots("PN001", "MG001", 18.520430, 73.856743, 15)
	assert.Nil(t,err,"Error should be nil")
	b, e := json.Marshal(spots)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
}

func TestHandleReserveParkingSpot(t *testing.T) {
	response, err := HandleReserveParkingSpot("PN001", "MG001", "spot001")
	assert.Nil(t,err,"Error should be nil")
	b, e := json.Marshal(response)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	assert.Equal(t,response.Message,"Successfully booked the parking")

	spots, err := HandleGetParkingSpots("PN001", "MG001","")
	assert.Nil(t,err,"Error should be nil")
	b, e = json.Marshal(spots)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	parkingSpots := spots[0]
	spot := parkingSpots["spot001"]
	assert.Equal(t, spot.Id,"spot001","parking spot not found")
	assert.Equal(t,spot.BookingStatus.Status,entities.OCCUPIED)
}

func TestHandleCancelParkingSpot(t *testing.T) {
	response, err := HandleReserveParkingSpot("PN001", "MG001", "spot001")
	assert.Nil(t,err,"Error should be nil")
	b, e := json.Marshal(response)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	assert.Equal(t,response.Message,"Successfully booked the parking")

	spots, err := HandleGetParkingSpots("PN001", "MG001","")
	assert.Nil(t,err,"Error should be nil")
	b, e = json.Marshal(spots)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	parkingSpots := spots[0]
	spot := parkingSpots["spot001"]
	assert.Equal(t, spot.Id,"spot001","parking spot not found")
	assert.Equal(t,spot.BookingStatus.Status,entities.OCCUPIED)

	spotResponse, err := HandleCancelParkingSpot("PN001", "MG001", "spot001")
	assert.Nil(t,err,"Error should be nil")
	b, e = json.Marshal(spotResponse)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))

	spots, err = HandleGetParkingSpots("PN001", "MG001","")
	assert.Nil(t,err,"Error should be nil")
	b, e = json.Marshal(spots)
	assert.Nil(t,e,"marshall error should be nil")
	t.Logf("%v",string(b))
	parkingSpots = spots[0]
	spot = parkingSpots["spot001"]
	assert.Equal(t, spot.Id,"spot001","parking spot not found")
	assert.Equal(t,spot.BookingStatus.Status,entities.AVAILABLE)
}