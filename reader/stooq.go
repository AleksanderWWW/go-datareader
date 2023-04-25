package reader

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-gota/gota/dataframe"
)

var frequenciesAvailable = map[string]bool {
	"d": true,
	"w": true, 
	"m": true,
	"q": true,
	"y": true,
}

type StooqDataReader struct {
	symbols []string
	startDate time.Time
	endDate time.Time
	freq string
	baseUrl string
}

func (sdr StooqDataReader) getParams(symbol string) map[string]string{
	return map[string]string {
		"s": symbol,
		"i": sdr.freq,
		"d1": strings.Replace(sdr.startDate.Format("2006-01-02"), "-", "", -1),
		"d2": strings.Replace(sdr.endDate.Format("2006-01-02"), "-", "", -1),
	}
}

func (sdr StooqDataReader) getResponse(params map[string]string, headers map[string]string) (string, error) {
	req, err := createRequest(params, headers, sdr.baseUrl)
	if err != nil {
		return "", err  
	}

	client := &http.Client{}
    resp, err := client.Do(req)

	if err != nil {
		return "", err  
	}

	defer resp.Body.Close()
	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err  
	}

	return string(respText), nil
}

func (srd StooqDataReader) parseResponse(respText string, symbol string) ([]SingleRecord, error) {
	lines := strings.Split(respText, "\n")
	records := make([]SingleRecord, 0, len(lines))

	for _, line := range lines[1:len(lines)-1] {
		record, err := parseStooqLine(line, symbol)
		if err != nil {
			return []SingleRecord{}, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (sdr StooqDataReader) Read() dataframe.DataFrame {
	results := make([]dataframe.DataFrame, 0, len(sdr.symbols))
	var wg sync.WaitGroup
	
	for _, symbol := range sdr.symbols {
		
		wg.Add(1)

		go func(symbol string) {
			defer wg.Done()
			params := sdr.getParams(symbol)
			data, err := sdr.getResponse(params, DefaultHeaders)

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
		df = df.Rename(symbol + "-" + name, name)
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
