package exchanger

import "net/http"

// ExchangeRateGetter TODO
type ExchangeRateGetter interface {
	GetCurrencyBuyRate(code string, client http.Client) (float64, error)
	GetCurrencySellRate(code string, client http.Client) (float64, error)
}
