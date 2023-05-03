package reader

import (
	"testing"
	"time"
)

func TestNewFredReader(t *testing.T) {
	fredReader, err := NewFredDataReader([]string{"PKO"}, time.Now().AddDate(0, 0, -1), time.Now())

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	fredUrl := BaseUrlMap["fred"]
	if fredReader.baseUrl != fredUrl {
		t.Errorf("FAIL: expected %s, got %s", fredUrl, fredReader.baseUrl)
	}
}

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

func TestNewBOCDataReader(t *testing.T) {
	bocReader, err := NewBOCDataReader([]string{"FXUSDCAD"}, time.Now().AddDate(0, 0, -1), time.Now())

	if err != nil {
		t.Errorf("FAIL: %s", err)
	}

	bocUrl := BaseUrlMap["boc"]
	if bocReader.baseUrl != bocUrl {
		t.Errorf("FAIL: expected %s, got %s", bocUrl, bocReader.baseUrl)
	}
}
