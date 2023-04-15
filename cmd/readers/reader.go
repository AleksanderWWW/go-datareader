package reader

import (
	"github.com/go-gota/gota/dataframe"
)

type SingleRecord struct {
	Symbol string
	Date string
	Open float64
	High float64
	Low float64
	Close float64
	Volume int
}

type DataReader interface {
	Read() map[string]dataframe.DataFrame
	getParams(args ...any) map[string]string
	getResponse(params map[string]string, headers map[string]string) (string, error)
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
