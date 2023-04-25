# go-datareader: financial datareader

`go-datareader` is a library for downloading financial data in a tabular form. It's written in [Go](https://go.dev/) and therefore aims to be a more performant counterpart of the Python's [pandas-datareader](https://github.com/pydata/pandas-datareader).

The project currently support the following data providers:

- [Stooq](https://stooq.com)


## Advantages over pandas-datareader

The two main advantages of `go-datareader` over it's counterpart are:

- better overall performance due to string typing and a compiled nature of the Go programming langauge, compared to interpreted, dynamically-typed Python
- faster data extraction due to the usage of `goroutines`



## Example usage
Gather Stooq daily quotes for a couple of tickers from the last 100 days.

```
stooqReader, err := reader.NewStooqDataReader(
		[]string{"PKO", "KGH", "PZU"},  // stooq tickers
		time.Now().AddDate(0, 0, -100),  // start date
		time.Now(),  // end date
		"d",  // daily
	)

// error handling
// ...

data := stooqReader.Read()  // returns a DataFrame object
```

The data is in the form of the `gota dataframe.DataFrame`. 
