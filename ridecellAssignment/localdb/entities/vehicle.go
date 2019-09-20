package entities

type VehicleType string

const (
	CAR  VehicleType = "car"
	BIKE VehicleType = "bike"
)

type Vehicle struct {
	RegistrationNumber string
	Type VehicleType
}