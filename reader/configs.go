package reader

import "time"

type FredReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
}

type BOCReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
}

type StooqReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	Freq      string
}

type TiingoReaderConfig struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	ApiKey    string
}
