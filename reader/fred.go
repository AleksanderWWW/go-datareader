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
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type FredDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
}

func NewFredDataReader(config FredReaderConfig) (*FredDataReader, error) {
	// defaults
	if config.StartDate.IsZero() {
		config.StartDate = time.Now().AddDate(-5, 0, 0)
	}

	if config.EndDate.IsZero() {
		config.EndDate = time.Now()
	}

	return &FredDataReader{
		symbols:   config.Symbols,
		startDate: config.StartDate,
		endDate:   config.EndDate,
	}, nil
}

func (fdr FredDataReader) getName() string {
	return "fred"
}

func (fdr *FredDataReader) getSymbols() []string {
	return fdr.symbols
}

func (fdr *FredDataReader) url(symbol string) string {
	return fmt.Sprintf(FredBaseUrl, symbol)
}

func (fdr *FredDataReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	data, err := getResponse(nil, nil, fdr.url(symbol))

	if err != nil {
		return dataframe.DataFrame{}, err
	}

	df := dataframe.ReadCSV(strings.NewReader(data),
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.Float),
		dataframe.WithTypes(map[string]series.Type{
			"DATE": series.String,
		}))

	if df.Error() != nil {
		return dataframe.DataFrame{}, df.Error()
	}

	df = df.Filter(
		dataframe.F{
			Colidx:     0,
			Colname:    "DATE",
			Comparator: series.CompFunc,
			Comparando: filterDates(fdr.startDate, fdr.endDate),
		},
	)
	if df.Error() != nil {
		return dataframe.DataFrame{}, df.Error()
	}

	return df, nil
}

func (fdr *FredDataReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	if len(dfs) == 0 {
		log.Printf("[WARNING] Returning empty data frame for %s", fdr.getName())
		return dataframe.DataFrame{}
	}

	combined := dfs[0]
	if len(dfs) > 1 {

		for _, df := range dfs[1:] {
			combined = combined.InnerJoin(df, "DATE")
		}
	}

	return combined
}
