package reader

import (
	"errors"
	"time"
)


type StooqDataReader struct {
	symbols []string
	startDate time.Time
	endDate time.Time
	baseUrl string
}

func NewStooqDataReader(symbols []string, startDate time.Time, endDate time.Time) (*StooqDataReader, error) {
	baseUrl, ok := BaseUrlMap["stooq"]

	if !ok {
		return &StooqDataReader{}, errors.New("Could not find stooq base url")
	}

	return &StooqDataReader{
		symbols: symbols,
		startDate: startDate,
		endDate: endDate,
		baseUrl: baseUrl,
	}, nil
}
