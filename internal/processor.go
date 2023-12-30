package internal

import (
	"dollar-bot/internal/strategy"
	"dollar-bot/internal/types"
	"dollar-bot/internal/utils"
	"errors"
)

type Processor struct{}

func (p *Processor) GetCotizations(moneyType string) ([]types.MoneyDetail, error) {
	var api strategy.DolarApiInterface

	if moneyType == utils.Dollar {
		api = &strategy.Dollar{}
	}

	if moneyType == utils.Euro {
		api = &strategy.Euro{}
	}

	if moneyType == utils.Real {
		api = &strategy.Real{}
	}

	if moneyType == utils.PesoChilean {
		api = &strategy.PesoChilean{}
	}

	if moneyType == utils.PesoUruguayan {
		api = &strategy.PesoUruguayan{}
	}

	if api == nil {
		return nil, errors.New("money type not supported")
	}

	return api.GetMoneyDetail()
}
