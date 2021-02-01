package dao

import(
	"gjek/model"
	"gjek/config"
	"fmt"
)

type LocationDaoImpl struct{}

func (LocationDaoImpl) GetLocationByName(name string) (l model.Location, e error) {
	defer config.CatchError(&e)
	var data model.Location
	fmt.Println("here21")
	fmt.Println(name)
	fmt.Println("here22")
	res := g.Where("name = ?", name).Find(&data)

	fmt.Println(res.Error)

	if res.Error == nil {
		return data, nil
	}
	return data, res.Error
}