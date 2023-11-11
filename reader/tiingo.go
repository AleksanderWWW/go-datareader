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

type TiingoDailyReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	baseUrl   string
	apiKey    string
}

type TiingoReaderConfig struct {
	startDate time.Time
	endDate   time.Time
	apiKey    string
}

func NewTiingoDailyReader(symbols []string,
	config TiingoReaderConfig) (*TiingoDailyReader, error) {

	// defaults
	if config.startDate.IsZero() {
		config.startDate = time.Now().AddDate(-5, 0, 0)
	}

	if config.endDate.IsZero() {
		config.endDate = time.Now()
	}

	if len(config.apiKey) == 0 {
		config.apiKey = os.Getenv(TIINGO_API_KEY)
	}

	if len(config.apiKey) == 0 {
		return &TiingoDailyReader{}, fmt.Errorf("API token not found")
	}

	return &TiingoDailyReader{
		symbols:   symbols,
		startDate: config.startDate,
		endDate:   config.endDate,
		baseUrl:   "https://api.tiingo.com/tiingo/daily/%s/prices",
		apiKey:    config.apiKey,
	}, nil
}

func (tdr *TiingoDailyReader) getName() string {
	return "tiingo"
}

func (tdr *TiingoDailyReader) getSymbols() []string {
	return tdr.symbols
}

func (tdr *TiingoDailyReader) params() map[string]string {
	return map[string]string{
		"startDate": tdr.startDate.Format("2006-01-02"),
		"endDate":   tdr.endDate.Format("2006-01-02"),
		"format":    "csv",
	}
}

func (tdr *TiingoDailyReader) headers() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Token %s", tdr.apiKey),
	}
}

func (tdr *TiingoDailyReader) url(symbol string) string {
	return fmt.Sprintf(tdr.baseUrl, symbol)
}

func (tdr *TiingoDailyReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	data, err := getResponse(tdr.params(), tdr.headers(), tdr.url(symbol))
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	df := dataframe.ReadCSV(strings.NewReader(data))
	return renameDataframe(df, symbol), nil
}

func (tdr *TiingoDailyReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	combined := dfs[0]
	if len(dfs) > 1 {
		for _, df := range dfs[1:] {
			combined = combined.OuterJoin(df, "date")
		}
	}

	return combined
}
