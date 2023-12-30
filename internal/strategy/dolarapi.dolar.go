package strategy

import (
	"dollar-bot/internal/types"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"os"
)

type Dollar struct{}

type Euro struct{}

type Real struct{}

type PesoChilean struct{}

type PesoUruguayan struct{}

func (d *Dollar) GetMoneyDetail() ([]types.MoneyDetail, error) {
	return RetrieveData("DOLAR_API_DOLAR")
}

func (e *Euro) GetMoneyDetail() ([]types.MoneyDetail, error) {
	return RetrieveData("DOLAR_API_EURO")
}

func (e *Real) GetMoneyDetail() ([]types.MoneyDetail, error) {
	return RetrieveData("DOLAR_API_REAL")
}

func (e *PesoChilean) GetMoneyDetail() ([]types.MoneyDetail, error) {
	return RetrieveData("DOLAR_API_CH")
}

func (e *PesoUruguayan) GetMoneyDetail() ([]types.MoneyDetail, error) {
	return RetrieveData("DOLAR_API_UY")
}

func RetrieveData(url string) ([]types.MoneyDetail, error) {
	client := resty.New()

	resp, err := client.R().EnableTrace().Get(os.Getenv(url))

	if url == "DOLAR_API_DOLAR" {
		var md []types.MoneyDetail
		err = json.Unmarshal(resp.Body(), &md)
		return md, nil
	}

	var md types.MoneyDetail
	err = json.Unmarshal(resp.Body(), &md)

	if err != nil {
		return nil, err
	}

	result := make([]types.MoneyDetail, 0)
	return append(result, md), nil
}
