package entities

type User struct {
	Id string
	Name string
	Phone string
	Vehicles []*Vehicle
}
