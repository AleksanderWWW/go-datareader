package reader

import (
	"reflect"
	"testing"

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

func TestGetData(t *testing.T) {
	mockReader := MockReader{}

	expected := dataframe.New(
		series.New([]int{1, 2}, series.Int, "Index"),
		series.New([]string{"b", "a"}, series.String, "testSymbol1"),
		series.New([]string{"b", "a"}, series.String, "testSymbol2"),
		series.New([]string{"b", "a"}, series.String, "testSymbol3"),
	)

	obtained := GetData(&mockReader)

	if !reflect.DeepEqual(expected.Records(), obtained.Records()) {
		t.Errorf("Different values:\nExpected:%v\nObtained:%v", expected.Records(), obtained.Records())
	}

	if !reflect.DeepEqual(expected.Types(), obtained.Types()) {
		t.Errorf("Different types:\nExpected:%v\nObtained:%v", expected.Types(), obtained.Types())
	}

	if mockReader.readSingleCallCount != 3 {
		t.Errorf("Expected 'readSingle' to be called 3 times. Called %d times", mockReader.readSingleCallCount)
	}

	if mockReader.concatDataframesCallCount != 1 {
		t.Errorf("Expected 'concatDataframed' to be called once. Called %d times", mockReader.concatDataframesCallCount)
	}
}
