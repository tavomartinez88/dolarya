package types

type MoneyDetail struct {
	Money    string  `json:"moneda"`
	Kind     string  `json:"casa"`
	Name     string  `json:"nombre"`
	Buy      float64 `json:"compra"`
	Sell     float64 `json:"venta"`
	UpdateAt string  `json:"fechaActualizacion"`
}
