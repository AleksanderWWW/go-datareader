package reader

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-gota/gota/series"
)

func TestGetParams(t *testing.T) {
	d1 := time.Now().AddDate(0, 0, -1)
	d2 := time.Now()
	stooqReader, err := NewStooqDataReader(
		StooqReaderConfig{
			Symbols:   []string{"PKO"},
			StartDate: d1,
			EndDate:   d2,
			Freq:      "d",
		},
	)

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	expectedParams := map[string]string{
		"s":  "PKO",
		"i":  "d",
		"d1": d1.Format("20060102"),
		"d2": d2.Format("20060102"),
	}
	params := stooqReader.getParams("PKO")

	cmp := reflect.DeepEqual(params, expectedParams)

	if !cmp {
		fmt.Println("Expected: ", expectedParams)
		fmt.Println("Got: ", params)
		t.Error("Params don't match")
	}
}

func TestRead(t *testing.T) {
	stooqReader, err := NewStooqDataReader(
		StooqReaderConfig{
			Symbols: []string{"PKO"},
			Freq:    "d",
		},
	)

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	df := GetData(stooqReader)

	if len(df.Records()) == 0 {
		t.Error("FAILED empty dataframe")
	}
	types := df.Types()
	expectedTypes := []series.Type{"string", "float", "float", "float", "float", "int"}

	comp := reflect.DeepEqual(types, expectedTypes)

	if !comp {
		fmt.Println(types)
		t.Error("FAILED types missmatch")
	}
}
