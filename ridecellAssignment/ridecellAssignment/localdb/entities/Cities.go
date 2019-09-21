package entities

/*
	A city contains multiple streets
 */
type City struct {
	Id string
	Name string
	Streets []map[string]*Street
}
