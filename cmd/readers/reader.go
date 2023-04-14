package reader

import (
	"time"

	"github.com/go-gota/gota/dataframe"
)

type SingleRecord struct {
	date time.Time
	open float32
	high float32
	low float32
	close float32
}

type DataReader interface {
	read() dataframe.DataFrame
	getParams() map[string]any
}

var BaseUrlMap = map[string]string {
	"stooq": "https://stooq.com/q/d/l",
}
