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

type BOCDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	baseUrl   string
}

func NewBOCDataReader(symbols []string, startDate time.Time, endDate time.Time) (*BOCDataReader, error) {
	baseUrl := "http://www.bankofcanada.ca/valet/observations"

	return &BOCDataReader{
		symbols:   symbols,
		startDate: startDate,
		endDate:   endDate,
		baseUrl:   baseUrl,
	}, nil
}

func (bdr BOCDataReader) getName() string {
	return "bank-of-canada"
}

func (bdr BOCDataReader) getSymbols() []string {
	result := strings.Join(bdr.symbols, ",")

	return []string{result}
}

func (bdr BOCDataReader) getParams() map[string]string {
	return map[string]string{
		"start_date": bdr.startDate.Format("2006-01-02"),
		"end_date":   bdr.endDate.Format("2006-01-02"),
	}
}

func (bdr BOCDataReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	params := bdr.getParams()
	url := fmt.Sprintf("%s/%s/csv", bdr.baseUrl, symbol)
	data, err := getResponse(params, DefaultHeaders, url)

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
	return dataframe.DataFrame{}
}
