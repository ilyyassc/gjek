package service

import(
	// "errors"
	// "math"
	"gjek/model"
	"gjek/dao"
	"gjek/config"
	"fmt"

)

type GsendPriceCalculationServiceImpl struct{}

var locationDao dao.LocationDao = dao.LocationDaoImpl{}

func (GsendPriceCalculationServiceImpl) CalculatePriceGsend(data *model.GsendTrxInput) (res int, e error){
	defer config.CatchError(&e)
	fmt.Println("here")
	pinpointPickup, err := locationDao.GetLocationByName(data.Pickup)
	destinationPickup, err := locationDao.GetLocationByName(data.Destination)

	fmt.Println("here2")

	if err != nil{
		return 0, err
	}

	fmt.Println("here3")

	fmt.Println(pinpointPickup.Coordinate)
	fmt.Println(destinationPickup.Coordinate)

	res = (pinpointPickup.Coordinate - destinationPickup.Coordinate) * data.Weight * 1

	if res < 0{
		res = -res
	}

	return res, err
}
