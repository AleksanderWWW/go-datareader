# go-datareader: financial datareader

`go-datareader` is a library for downloading financial data in a tabular form. It's written in [Go](https://go.dev/) and therefore aims to be a more performant counterpart of the Python's [pandas-datareader](https://github.com/pydata/pandas-datareader).

The project currently supports the following data providers:

- [Stooq](https://stooq.com)
- [FRED](https://fred.stlouisfed.org)
- [Bank of Canada](https://www.bankofcanada.ca/)
- [Tiingo](https://www.tiingo.com/)


## Advantages over pandas-datareader

The two main advantages of `go-datareader` over it's counterpart are:

- better overall performance due to strong typing and a compiled nature of the Go programming langauge, compared to the interpreted, dynamically-typed Python
- faster data extraction due to the usage of `goroutines` to send the requests concurrently.

## Getting started

Run the following command to install the `go-datareader`:

```sh
$ go get -u github.com/AleksanderWWW/go-datareader
```


## Example usage
Gather quotes for a couple of tickers from the last 100 days.
The returned data is in the form of the [gota](https://github.com/go-gota/gota) dataframe. Symbols for which the data could not be obtained are omitted.

### Stooq
```
stooqReader, err := reader.NewStooqDataReader(
	reader.StooqReaderConfig{
		Symbols:   []string{"PKO", "KGH", "PZU"},
		StartDate: time.Now().AddDate(0, 0, -100),
		EndDate:   time.Now(),
		Freq:      "d",
	},
)

// error handling
// ...

data := reader.GetData(stooqReader)  // returns a DataFrame object
```

In this example the quotes are obtained in a "daily" mode. Other available options are:
- "w": weekly
- "m": monthly
- "q": quarterly
- "y": yearly

If no frequency is provided, the reader will default to "d".

If `startDate` or `endDate` are not provided, the reader will default to '5 years ago' and 'now' respectively.

---

### FRED
```
fredReader, err := reader.NewFredDataReader(
		reader.FredReaderConfig{
			Symbols:   []string{"SP500", "DJIA", "VIXCLS"},
			StartDate: time.Now().AddDate(0, 0, -100),
			EndDate:   time.Now(),
		},
	)

// error handling
// ...

data := reader.GetData(fredReader)
```

If `startDate` or `endDate` are not provided, the reader will default to '5 years ago' and 'now' respectively.

---

### Bank of Canada
```
bocReader, err := reader.NewBOCDataReader(
		reader.BOCReaderConfig{
			Symbols:   []string{"FXUSDCAD", "FXCADIDR", "FXCADPEN"},
			StartDate: time.Now().AddDate(0, 0, -100),
			EndDate:   time.Now(),
		},
	)

// error handling
// ...

data := reader.GetData(bocReader)
```

The list of available symbols can be found [here](https://www.bankofcanada.ca/valet/lists/series).

If `startDate` or `endDate` are not provided, the reader will default to '5 years ago' and 'now' respectively.

---

### Tiingo
```
startDate := time.Now().AddDate(0, 0, -4)
endDate := time.Now()
apiKey := "my-secret-api-key"
os.Setenv("TIINGO_API_KEY", apiKey)  // either export the key as env variable...

tiingoReader, err := reader.NewTiingoDailyReader(
	[]string{"ZZZOF", "000001"},
	tiingoReader, _ := reader.NewTiingoDailyReader(
	[]string{"ZZZOF", "000001"},
	reader.TiingoReaderConfig{
		StartDate: startDate,
		EndDate:   endDate,
		ApiKey:    apiKey  // ... or pass it here.
	},
)

)

// error handling
// ...

data := reader.GetData(tiingoReader)
```

The list of available symbols can be found [here](https://apimedia.tiingo.com/docs/tiingo/daily/supported_tickers.zip).

There are two ways to pass the `Tiingo` API token - either explicitly in the `TiingoReaderConfig` (takes precedence),
or via a `TIINGO_API_KEY` envirionment variable (recommended option).

If `startDate` or `endDate` are not provided, the reader will default to '5 years ago' and 'now' respectively.
