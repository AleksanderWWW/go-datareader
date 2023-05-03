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

func NewBOCDataReader(symbols []string, startDate time.Time, endDate time.Time) (*BOCDataReader, error) {
	baseUrl, ok := BaseUrlMap["boc"]

	if !ok {
		return &BOCDataReader{}, errors.New("Could not find Bank of Canada base url")
	}

	return &BOCDataReader{
		symbols:   symbols,
		startDate: startDate,
		endDate:   endDate,
		baseUrl:   baseUrl,
	}, nil
}
