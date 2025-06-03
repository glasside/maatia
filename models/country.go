package models

type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Region     string `json:"region"`
	Population int    `json:"population"`
}
