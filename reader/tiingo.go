/*
Copyright © 2023 Aleksander Wojnarowicz <alwojnarowicz@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reader

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

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
	apiKey    string
}

type TiingoReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	ApiKey    string
}

func NewTiingoDailyReader(config TiingoReaderConfig) (*TiingoDailyReader, error) {

	// defaults
	if config.StartDate.IsZero() {
		config.StartDate = time.Now().AddDate(-5, 0, 0)
	}

	if config.EndDate.IsZero() {
		config.EndDate = time.Now()
	}

	if len(config.ApiKey) == 0 {
		apiKey, ok := os.LookupEnv(TIINGO_API_KEY)
		if !ok {
			return &TiingoDailyReader{}, fmt.Errorf("API token not found")
		}

		config.ApiKey = apiKey
	}

	return &TiingoDailyReader{
		symbols:   config.Symbols,
		startDate: config.StartDate,
		endDate:   config.EndDate,
		apiKey:    config.ApiKey,
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
	return fmt.Sprintf(TiingoBaseUrl, symbol)
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
	if len(dfs) == 0 {
		log.Printf("[WARNING] Returning empty data frame for %s", tdr.getName())
		return dataframe.DataFrame{}
	}

	combined := dfs[0]
	if len(dfs) > 1 {
		for _, df := range dfs[1:] {
			combined = combined.OuterJoin(df, "date")
		}
	}

	return combined
}
