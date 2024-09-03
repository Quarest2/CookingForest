package storage

var recipes = make([]Recipe, 0)

type Recipe struct {
	Name     string
	HasImage bool
	Time     string
	Persons  int
	Steps    []string
}
