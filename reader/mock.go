package reader

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type MockReader struct {
	readSingleCallCount       int
	concatDataframesCallCount int
}

func (mr MockReader) getName() string {
	return "mock-reader"
}

func (mr MockReader) getSymbols() []string {
	return []string{"testSymbol1", "testSymbol2", "testSymbol3"}
}

func (mr *MockReader) readSingle(symbol string) (dataframe.DataFrame, error) {
	mr.readSingleCallCount += 1
	return dataframe.New(
		series.New([]int{1, 2}, series.Int, "Index"),
		series.New([]string{"b", "a"}, series.String, symbol),
	), nil
}

func (mr *MockReader) concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame {
	combined := dfs[0]

	for _, df := range dfs[1:] {
		combined = combined.InnerJoin(df, "Index")
	}

	// to assert stable order of columns
	combined = combined.Select([]string{"Index", "testSymbol1", "testSymbol2", "testSymbol3"})

	mr.concatDataframesCallCount += 1
	return combined
}
