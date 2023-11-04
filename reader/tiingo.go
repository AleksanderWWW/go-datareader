package reader

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

const TIINGO_API_KEY = "TIINGO_API_KEY"

type TiingoRecord struct {
	Date        string  `json:"Date"`
	Close       float32 `json:"Close"`
	High        float32 `json:"High"`
	Low         float32 `json:"Low"`
	Open        float32 `json:"Open"`
	Volume      float32 `json:"Volume"`
	AdjClose    float32 `json:"AdjClose"`
	AdjHigh     float32 `json:"AdjHigh"`
	AdjLow      float32 `json:"AdjLow"`
	AdjOpen     float32 `json:"AdjOpen"`
	AdjVolume   float32 `json:"AdjVolume"`
	DivCash     float32 `json:"DivCash"`
	SplitFactor float32 `json:"SplitFactor"`
}

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
	} else {
		apiKeyVal = *apiKey
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

func (tdr *TiingoReader) params() map[string]string {
	return map[string]string{
		"startDate": tdr.startDate.Format("2006-01-02"),
		"endDate":   tdr.endDate.Format("2006-01-02"),
		"format":    "csv",
	}
}

func (tdr *TiingoReader) headers() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Token %s", tdr.apiKey),
	}
}

func (tdr *TiingoReader) url(symbol string) string {
	return fmt.Sprintf(tdr.baseUrl, symbol)
}

func (tdr *TiingoReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	data, err := getResponse(tdr.params(), tdr.headers(), tdr.url(symbol))
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	df := dataframe.ReadCSV(strings.NewReader(data))
	return renameDataframe(df, symbol), nil
}

func (tdr *TiingoReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	combined := dfs[0]
	if len(dfs) > 1 {
		for _, df := range dfs[1:] {
			combined = combined.OuterJoin(df, "date")
		}
	}

	return combined
}
