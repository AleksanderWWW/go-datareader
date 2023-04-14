package reader

import (
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
