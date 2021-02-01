package model

type GsendTrxInput struct {
	Destination	string `json:"destination"`
	Pickup		string `json:"pickup"`
	Weight		int `json:"weight"`
}