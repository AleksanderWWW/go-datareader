package reader

import (
	"strings"
	"sync"
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

func (sdr StooqDataReader) getParams(symbol string) map[string]string {
	return map[string]string{
		"s":  symbol,
		"i":  sdr.freq,
		"d1": strings.Replace(sdr.startDate.Format("2006-01-02"), "-", "", -1),
		"d2": strings.Replace(sdr.endDate.Format("2006-01-02"), "-", "", -1),
	}
}

func (sdr StooqDataReader) Read() dataframe.DataFrame {
	results := make([]dataframe.DataFrame, 0, len(sdr.symbols))
	var wg sync.WaitGroup

	for _, symbol := range sdr.symbols {

		wg.Add(1)

		go func(symbol string) {
			defer wg.Done()
			params := sdr.getParams(symbol)
			data, err := getResponse(params, DefaultHeaders, sdr.baseUrl)

			if err != nil {
				return
			}

			df := dataframe.ReadCSV(strings.NewReader(data))
			if df.Err != nil {
				return
			}

			df = renameDataframe(df, symbol)

			results = append(results, df)
		}(symbol)
	}

	wg.Wait()

	return concatDataframes(results)
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

func concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	combined := dfs[0]
	if len(dfs) > 1 {

		for _, df := range dfs[1:] {
			combined = combined.InnerJoin(df, "Date")
		}
	}

	return combined
}
