package reader

import (
	"os"
	"testing"
	"time"
)

func TestTiingoReaderDefaultInit(t *testing.T) {
	apiKey := "test-key"
	os.Setenv(TIINGO_API_KEY, apiKey)
	tdr, err := NewTiingoReader(
		[]string{},
		nil,
		nil,
		nil,
	)

	if err != nil {
		t.Error("Error during default initialization")
	}

	if tdr.apiKey != apiKey {
		t.Errorf("Wrong value of API key attribute.\n Expected: \t%s\n Actual: \t%s", apiKey, tdr.apiKey)
	}
}

func TestTiingoReaderCustomInit(t *testing.T) {
	startDate := time.Now().AddDate(-2, 0, 0)
	endDate := time.Now()
	apiKey := "MySecretAPIKey"
	tdr, err := NewTiingoReader(
		[]string{"sym1", "sym2"},
		&startDate,
		&endDate,
		&apiKey,
	)

	if err != nil {
		t.Error("Error during custom intialization")
	}

	if tdr.apiKey != apiKey {
		t.Errorf("Wrong value of API key attribute.\n Expected: \t%s\n Actual: \t%s", apiKey, tdr.apiKey)
	}
}
