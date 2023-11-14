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
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

type BOCReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
}

type BOCDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
}

func NewBOCDataReader(config BOCReaderConfig) (*BOCDataReader, error) {

	// defaults
	if config.StartDate.IsZero() {
		config.StartDate = time.Now().AddDate(-5, 0, 0)
	}

	if config.EndDate.IsZero() {
		config.EndDate = time.Now()
	}

	return &BOCDataReader{
		symbols:   config.Symbols,
		startDate: config.StartDate,
		endDate:   config.EndDate,
	}, nil
}

func (bdr *BOCDataReader) getName() string {
	return "bank-of-canada"
}

func (bdr *BOCDataReader) getSymbols() []string {
	result := strings.Join(bdr.symbols, ",")

	return []string{result}
}

func (bdr *BOCDataReader) getParams() map[string]string {
	return map[string]string{
		"start_date": bdr.startDate.Format("2006-01-02"),
		"end_date":   bdr.endDate.Format("2006-01-02"),
	}
}

func (bdr *BOCDataReader) url(symbol string) string {
	return fmt.Sprintf(BOCBaseUrl, symbol)
}

func (bdr *BOCDataReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	params := bdr.getParams()
	data, err := getResponse(params, DefaultHeaders, bdr.url(symbol))

	data_splitted := strings.Split(data, "OBSERVATIONS")

	if len(data_splitted) < 2 {
		err = errors.New(data)
		return dataframe.DataFrame{}, err
	}

	data = data_splitted[1][2:]

	if err != nil {
		return dataframe.DataFrame{}, err
	}

	df := dataframe.ReadCSV(strings.NewReader(data))
	if df.Error() != nil {
		return dataframe.DataFrame{}, df.Error()
	}

	return df, nil
}

func (bdr BOCDataReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	if len(dfs) > 0 {
		return dfs[0]
	}
	log.Printf("[WARNING] Returning empty data frame for %s", bdr.getName())
	return dataframe.DataFrame{}
}
