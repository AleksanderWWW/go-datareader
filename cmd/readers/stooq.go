package reader

import (
	"errors"
	"fmt"
	"time"
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
