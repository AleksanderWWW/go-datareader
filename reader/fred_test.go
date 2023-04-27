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
