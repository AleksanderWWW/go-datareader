package reader

import "time"

type SingleRecord struct {
	date time.Time
	open float32
	high float32
	low float32
	close float32
}

type DataReader interface {
	read() []SingleRecord
	getParams() map[string]any
}

var BaseUrlMap = map[string]string {
	"stooq": "https://stooq.com/q/d/l",
}
