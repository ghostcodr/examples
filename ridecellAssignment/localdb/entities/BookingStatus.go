package entities

import "time"

type Status string


const (
	AVAILABLE Status ="available"
	OCCUPIED Status ="occupied"
)

func GetStatus(status string) Status {
	switch status {
	case "available":
		return AVAILABLE
	case "occupied":
		return OCCUPIED
	default:
		return OCCUPIED
		//FIXME return an error
	}
}

type BookingStatus struct {
	Status Status
	From time.Time
	To time.Time
	UserDetails *User `json:"userDetails,omitempty"`
	VehicleDetails *Vehicle `json:"vehicleDetails,omitempty"`

}
