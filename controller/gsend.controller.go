package controller

import(
	"gjek/service"
	"gjek/config"
	"gjek/model"

	"github.com/labstack/echo/v4"
)

var gsendPriceService service.GsendPriceCalculationService = service.GsendPriceCalculationServiceImpl{}

func SetGsend(eg *echo.Group) {
	eg.GET("/gsend", getPrice)
}

func getPrice(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := new(model.GsendTrxInput)
	if err := c.Bind(data); err != nil{
		return resErr(c, err)
	}

	result, err := gsendPriceService.CalculatePriceGsend(data)
	if err == nil {
		return res(c, result)	
	}
	return resErr(c, err)

}