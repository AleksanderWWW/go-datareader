package reader

import (
	"log"
	"time"
)


type StooqDataReader struct {
	startDate time.Time
	endDate time.Time
	baseUrl string
}

func NewStooqReader(startDate time.Time, endDate time.Time) *StooqDataReader {
	baseUrl, ok := BaseUrlMap["stooq"]

	if !ok {
		log.Fatal("Could not find base url for stooq")
	}

	return &StooqDataReader{
		startDate: startDate,
		endDate: endDate,
		baseUrl: baseUrl,
	}
}
