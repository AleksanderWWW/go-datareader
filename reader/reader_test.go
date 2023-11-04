package reader

import (
	"reflect"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

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
