package service

import(
	"gjek/model"
)

type GsendPriceCalculationService interface {
	CalculatePriceGsend(data *model.GsendTrxInput) (int, error)
}