package dao

import (
	"gjek/model"
)

type LocationDao interface {
	GetLocationByName(name string) (model.Location, error)
}