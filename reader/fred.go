package reader

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type FredDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	baseUrl   string
}

func NewFredDataReader(symbols []string, startDate time.Time, endDate time.Time) (*FredDataReader, error) {
	baseUrl, ok := BaseUrlMap["fred"]

	if !ok {
		return &FredDataReader{}, errors.New("Could not find fred base url")
	}

	return &FredDataReader{
		symbols:   symbols,
		startDate: startDate,
		endDate:   endDate,
		baseUrl:   baseUrl,
	}, nil
}

func (fdr FredDataReader) getName() string {
	return "fred"
}

func (fdr *FredDataReader) getSymbols() []string {
	return fdr.symbols
}

func (fdr *FredDataReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	data, err := getResponse(nil, nil, fmt.Sprintf("%s?id=%s", fdr.baseUrl, symbol))

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
	combined := dfs[0]
	if len(dfs) > 1 {

		for _, df := range dfs[1:] {
			combined = combined.InnerJoin(df, "DATE")
		}
	}

	return combined
}
