package reader

import (
	"fmt"
	"os"
	"time"
)

const TIINGO_API_KEY = "TIINGO_API_KEY"

type TiingoReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	baseUrl   string
	apiKey    string
}

func NewTiingoReader(symbols []string,
	startDate *time.Time,
	endDate *time.Time,
	apiKey *string) (*TiingoReader, error) {

	var startDateVal time.Time
	var endDateVal time.Time
	var apiKeyVal string

	// defaults
	if startDate == nil {
		startDateVal = time.Now().AddDate(-5, 0, 0)
	} else {
		startDateVal = *startDate
	}

	if endDate == nil {
		endDateVal = time.Now()
	} else {
		endDateVal = *endDate
	}

	if apiKey == nil || len(*apiKey) == 0 {
		apiKeyVal = os.Getenv(TIINGO_API_KEY)
	}

	if len(apiKeyVal) == 0 {
		return &TiingoReader{}, fmt.Errorf("API token not found")
	}

	return &TiingoReader{
		symbols:   symbols,
		startDate: startDateVal,
		endDate:   endDateVal,
		baseUrl:   "https://api.tiingo.com/tiingo/daily/%s/prices",
		apiKey:    apiKeyVal,
	}, nil
}

func (tdr *TiingoReader) getName() string {
	return "tiingo"
}

func (tdr *TiingoReader) getSymbols() []string {
	return tdr.symbols
}

func (tdr *TiingoReader) getParams() map[string]string {
	return map[string]string{
		"startDate": tdr.startDate.Format("2006-01-02"),
		"endDate":   tdr.endDate.Format("2006-01-02"),
		"format":    "json",
	}
}

func (tdr *TiingoReader) getHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Token %s", tdr.apiKey),
	}
}

func (tdr *TiingoReader) getUrl(symbol string) string {
	return fmt.Sprintf(tdr.baseUrl, symbol)
}
