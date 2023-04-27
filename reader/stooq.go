package reader

import (
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

type StooqDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	freq      string
	baseUrl   string
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
	params := sdr.getParams(symbol)
	data, err := getResponse(params, DefaultHeaders, sdr.baseUrl)

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

func renameDataframe(df dataframe.DataFrame, symbol string) dataframe.DataFrame {
	for _, name := range df.Names() {
		if name == "Date" {
			continue
		}
		df = df.Rename(symbol+"-"+name, name)
	}
	return df
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
