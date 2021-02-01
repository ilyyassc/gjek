package model

type Location struct {
	Name string `json:"name"`
	Coordinate int `json:"coordinate"`
}

func (Location) TableName() string {
	return "locations"
}