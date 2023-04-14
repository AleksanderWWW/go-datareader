package reader

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
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
