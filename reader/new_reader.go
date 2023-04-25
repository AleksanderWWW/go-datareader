package reader

import (
	"errors"
	"fmt"
	"time"
)

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
		symbols:   symbols,
		freq:      freq,
		startDate: startDate,
		endDate:   endDate,
		baseUrl:   baseUrl,
	}, nil
}
