/*
Copyright © 2023 Aleksander WOjnarowicz <alwojnarowicz@gmail.com>

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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

var frequenciesAvailable = map[string]bool{
	"d": true,
	"w": true,
	"m": true,
	"q": true,
	"y": true,
}

type StooqReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	Freq      string
}

type StooqDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	freq      string
}

func NewStooqDataReader(config StooqReaderConfig) (*StooqDataReader, error) {
	if _, ok := frequenciesAvailable[config.Freq]; !ok {
		errMsg := fmt.Sprintf("Incorrect frequency chosen: %s", config.Freq)
		return &StooqDataReader{}, errors.New(errMsg)
	}

	// defaults
	if config.StartDate.IsZero() {
		config.StartDate = time.Now().AddDate(-5, 0, 0)
	}

	if config.EndDate.IsZero() {
		config.EndDate = time.Now()
	}

	if config.Freq == "" {
		config.Freq = "d"
	}

	return &StooqDataReader{
		symbols:   config.Symbols,
		freq:      config.Freq,
		startDate: config.StartDate,
		endDate:   config.EndDate,
	}, nil
}

func (sdr StooqDataReader) getName() string {
	return "stooq"
}

func (sdr StooqDataReader) getSymbols() []string {
	return sdr.symbols
}

func (sdr StooqDataReader) getParams(symbol string) map[string]string {
	return map[string]string{
		"s":  symbol,
		"i":  sdr.freq,
		"d1": strings.Replace(sdr.startDate.Format("2006-01-02"), "-", "", -1),
		"d2": strings.Replace(sdr.endDate.Format("2006-01-02"), "-", "", -1),
	}
}

func (sdr StooqDataReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	baseUrl := "https://stooq.com/q/d/l"
	params := sdr.getParams(symbol)
	data, err := getResponse(params, DefaultHeaders, baseUrl)

	if err != nil {
		return dataframe.DataFrame{}, err
	}

	df := dataframe.ReadCSV(strings.NewReader(data))
	if df.Error() != nil {
		return dataframe.DataFrame{}, df.Error()
	}

	df = renameDataframe(df, symbol)

	return df, nil
}

func (sdr StooqDataReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	combined := dfs[0]
	if len(dfs) > 1 {

		for _, df := range dfs[1:] {
			combined = combined.InnerJoin(df, "Date")
		}
	}

	return combined
}
