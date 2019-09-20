package entities

type ParkingSpotType string
type Currency string


const (
	SMALL ParkingSpotType = "small"
	LARGE ParkingSpotType = "large"
)

const (
	RS Currency ="rs"
)

type ParkingSpot struct {
	Id   string
	Lat  float64
	Lng  float64
	Type ParkingSpotType
	Rate float64
	Currency Currency
	BookingStatus BookingStatus //?? Keep booking status separate when we update we only need to change the booking status.
}
