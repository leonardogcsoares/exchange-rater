package exchanger

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ExchangeRateGetterImpl TODO
type ExchangeRateGetterImpl struct {
	// currencies currencyDetails
}

// ErrInvalidCurrencyCode TODO
type ErrInvalidCurrencyCode struct {
	Code string
}

func (err ErrInvalidCurrencyCode) Error() string {
	return fmt.Sprintf("Invalid currency code provided: %s", err.Code)
}

// GetCurrencyBuyRate TODO
func (e *ExchangeRateGetterImpl) GetCurrencyBuyRate(code string, client http.Client) (float64, error) {

	res, err := client.Get("https://www.bankofamerica.com/foreign-exchange/get-currdetails.go")
	if err != nil {
		return float64(0), err
	}

	var cs currencies
	err = json.NewDecoder(res.Body).Decode(&cs)
	if err != nil {
		return float64(0), err
	}

	curDetails := findCurrency(cs, code)
	if curDetails.CurrencyCode != code {
		return float64(0), ErrInvalidCurrencyCode{Code: code}
	}

	return curDetails.CurrencyBuyRateNumeric, nil
}

// GetCurrencySellRate TODO
func (e *ExchangeRateGetterImpl) GetCurrencySellRate(code string, client http.Client) (float64, error) {

	res, err := client.Get("https://www.bankofamerica.com/foreign-exchange/get-currdetails.go")
	if err != nil {
		return float64(0), err
	}

	var cs currencies
	err = json.NewDecoder(res.Body).Decode(&cs)
	if err != nil {
		return float64(0), err
	}

	curDetails := findCurrency(cs, code)
	if curDetails.CurrencyCode != code {
		return float64(0), ErrInvalidCurrencyCode{Code: code}
	}

	return curDetails.CurrencySellRateNumeric, nil
}

func findCurrency(cs currencies, code string) currencyDetails {
	for _, c := range cs {
		if c.CurrencyCode == code {
			return c
		}
	}

	return currencyDetails{}
}

type currencies []currencyDetails

type currencyDetails struct {
	CheckBuyRate            interface{} `json:"checkBuyRate"`
	CheckBuyRateNumeric     float64     `json:"checkBuyRateNumeric"`
	CheckSellRate           interface{} `json:"checkSellRate"`
	CheckSellRateNumeric    float64     `json:"checkSellRateNumeric"`
	CheckSmallestDenom      interface{} `json:"checkSmallestDenom"`
	ConversionChartType     interface{} `json:"conversionChartType"`
	CountryName             string      `json:"countryName"`
	CurrencyBuyRate         string      `json:"currencyBuyRate"`
	CurrencyBuyRateNumeric  float64     `json:"currencyBuyRateNumeric"`
	CurrencyCode            string      `json:"currencyCode"`
	CurrencyID              string      `json:"currencyId"`
	CurrencyName            string      `json:"currencyName"`
	CurrencyNamePlural      string      `json:"currencyNamePlural"`
	CurrencySellRate        string      `json:"currencySellRate"`
	CurrencySellRateNumeric float64     `json:"currencySellRateNumeric"`
	CurrencySiteID          interface{} `json:"currencySiteId"`
	CurrencySmallestDenom   int         `json:"currencySmallestDenom"`
	DeleteDate              interface{} `json:"deleteDate"`
	Image                   string      `json:"image"`
	IsResponseNull          interface{} `json:"isResponseNull"`
	Rank                    interface{} `json:"rank"`
	Region                  interface{} `json:"region"`
}
