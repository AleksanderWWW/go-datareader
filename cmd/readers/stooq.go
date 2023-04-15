package reader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func (sdr StooqDataReader) Read() []dataframe.DataFrame {
	results := make([]dataframe.DataFrame, 0, len(sdr.symbols))
	for _, symbol := range sdr.symbols {
		params := sdr.getParams(symbol)

		data, err := sdr.getResponse(params, DefaultHeaders)

		if err != nil {
			continue
		}

		records, err := sdr.parseResponse(data, symbol)
		if err != nil {
			continue
		}
		results = append(results, dataframe.LoadStructs(records))
	}
	return results
}


func NewStooqDataReader(symbols []string, startDate time.Time, endDate time.Time, freq string) (*StooqDataReader, error) {
	baseUrl, ok := BaseUrlMap["stooq"]

	if !ok {
		return &StooqDataReader{}, errors.New("Could not find stooq base url")
	}

	if _, ok := frequenciesAvailable[freq]; !ok {
		errMsg := fmt.Sprintf("Incorrect frequency chosen: %s", freq)
		return &StooqDataReader{}, errors.New(errMsg)
	}

	return &StooqDataReader{
		symbols: symbols,
		freq: freq,
		startDate: startDate,
		endDate: endDate,
		baseUrl: baseUrl,
	}, nil
}
