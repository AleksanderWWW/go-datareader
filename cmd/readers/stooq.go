package reader

import (
	"errors"
	"fmt"
	"io"
	"time"
	"net/http"
    "net/url"
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
		"d1": sdr.startDate.Format("2006-01-02"),
		"d2": sdr.endDate.Format("2006-01-02"),
	}
}

func (sdr StooqDataReader) getResponse(params map[string]string, headers map[string]string) (string, error) {
	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}

    u, err := url.ParseRequestURI(sdr.baseUrl)
	if err != nil {
		return "", err  
	}

    u.RawQuery = parameters.Encode()
    urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", err  
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

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
