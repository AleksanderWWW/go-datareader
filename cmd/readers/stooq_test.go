package reader

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-gota/gota/series"
)

func TestNewStooqDataReader(t *testing.T) {
	stooqReader, err := NewStooqDataReader([]string{"PKO"}, time.Now().AddDate(0, 0, -1), time.Now(), "d")

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	stooqUrl := BaseUrlMap["stooq"]
	if stooqReader.baseUrl != stooqUrl {
		t.Errorf("FAIL: expected %s, got %s", stooqUrl, stooqReader.baseUrl)
	}
}

func TestGetParams(t *testing.T) {
	d1 := time.Now().AddDate(0, 0, -1)
	d2 := time.Now()
	stooqReader, err := NewStooqDataReader([]string{"PKO"}, d1, d2, "d")

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	expectedParams := map[string]string {
		"s": "PKO",
		"i": "d",
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

func TestGetResponse(t *testing.T) {
	stooqReader, err := NewStooqDataReader([]string{"PKO"}, time.Now().AddDate(0, 0, -10), time.Now(), "d")

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}
	params := stooqReader.getParams("PKO")
	respText, err := stooqReader.getResponse(params, map[string]string{})
	if err != nil {
		t.Errorf("FAIL: %s", err)
	}
	
	lines := strings.Split(respText, "\n")

	if len(lines) < 2 {
		fmt.Println(lines)
		t.Error("FAILED data was not retrieved correctly")
	}
}

func TestRead(t *testing.T) {
	stooqReader, err := NewStooqDataReader([]string{"PKO"}, time.Now().AddDate(0, 0, -10), time.Now(), "d")

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	df := stooqReader.Read()["PKO"]

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
