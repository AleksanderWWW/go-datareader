package reader

import (
	"errors"
	"time"
)


type StooqDataReader struct {
	startDate time.Time
	endDate time.Time
	baseUrl string
}

func NewStooqDataReader(startDate time.Time, endDate time.Time) (*StooqDataReader, error) {
	baseUrl, ok := BaseUrlMap["stooq"]

	if !ok {
		return &StooqDataReader{}, errors.New("Could not find stooq base url")
	}

	return &StooqDataReader{
		startDate: startDate,
		endDate: endDate,
		baseUrl: baseUrl,
	}, nil
}
