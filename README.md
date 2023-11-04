# go-datareader: financial datareader

`go-datareader` is a library for downloading financial data in a tabular form. It's written in [Go](https://go.dev/) and therefore aims to be a more performant counterpart of the Python's [pandas-datareader](https://github.com/pydata/pandas-datareader).

The project currently supports the following data providers:

- [Stooq](https://stooq.com)
- [FRED](https://fred.stlouisfed.org)
- [Bank of Canada](https://www.bankofcanada.ca/)


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
		[]string{"PKO", "KGH", "PZU"},  // stooq tickers
		time.Now().AddDate(0, 0, -100),  // start date
		time.Now(),  // end date
		"d",  // daily
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

---

### FRED
```
fredReader, err := reader.NewFredDataReader(
		[]string{"SP500", "DJIA", "VIXCLS"},
		time.Now().AddDate(0, 0, -100),
		time.Now(),
	)

// error handling
// ...

data := reader.GetData(fredReader)
```

---

### Bank of Canada
```
bocReader, err := reader.NewBOCDataReader(
		[]string{"FXUSDCAD", "FXCADIDR", "FXCADPEN"},
		time.Now().AddDate(0, 0, -100),
		time.Now(),
	)

// error handling
// ...

data := reader.GetData(bocReader)
```

The list of available symbols can be found [here](https://www.bankofcanada.ca/valet/lists/series).

---


### Tiingo
```
	startDate := time.Now().AddDate(0, 0, -4)
	endDate := time.Now()
	apiKey := "my-secret-api-key"
	os.Setenv(TIINGO_API_KEY, apiKey)  // either export the key as env variable...

	tiingoReader, err := reader.NewTiingoDailyReader(
		[]string{"ZZZOF", "000001"},
		&startDate,  // default is 5 yrs before current date (if nil)
		&endDate,  // default is current date (if nil)
		nil,  // ...or pass it here as *apiKey

	)

	// error handling
	// ...

	data := reader.GetData(tiingoReader)
```

The list of available symbols can be found [here](https://apimedia.tiingo.com/docs/tiingo/daily/supported_tickers.zip).
