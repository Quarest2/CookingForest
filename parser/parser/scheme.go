package parser

type Recipe struct {
	Name      string   `json:"name"`
	ImagePath string   `json:"image_path"`
	Time      string   `json:"time"`
	Persons   int      `json:"persons"`
	Steps     []string `json:"steps"`
}
