package reader

import (
	"time"

	"github.com/go-gota/gota/dataframe"
)

type SingleRecord struct {
	date time.Time
	open float64
	high float64
	low float64
	close float64
	volume int64
}

type DataReader interface {
	read() dataframe.DataFrame
	getParams(args ...any) map[string]string
}

var BaseUrlMap = map[string]string {
	"stooq": "https://stooq.com/q/d/l",
}

var DefaultHeaders = map[string]string {
	"Connection": "keep-alive",
    "Expires": "-1",
    "Upgrade-Insecure-Requests": "1",
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
}
