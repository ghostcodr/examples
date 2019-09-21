package entities

/*
A street within a city, having multiple parking spots
 */
type Street struct {
	Id string
	Name string
	ParkingSpots []map[string]*ParkingSpot
}
