package requests

type ReserveParkingSpotRequest struct {
	UserId string `json:"user_id"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
}
