package strategy

import "dollar-bot/internal/types"

type DolarApiInterface interface {
	GetMoneyDetail() ([]types.MoneyDetail, error)
}
