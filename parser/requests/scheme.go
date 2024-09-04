package requests

type Recipe struct {
	Name     string
	HasImage bool
	Time     string
	Persons  int
	Steps    []string
}
