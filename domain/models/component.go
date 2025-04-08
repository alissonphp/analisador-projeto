package models

type Component struct {
	ID          string    `json:"id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Qualifier   string    `json:"qualifier"`
	Measures    []Measure `json:"measures"`
}
