package parser

type Recipe struct {
	Name    string
	Image   []byte
	Time    string
	Persons int
	Steps   []string
}
