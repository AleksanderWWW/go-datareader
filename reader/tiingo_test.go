package reader

import (
	"os"
	"testing"
	"time"
)

func TestTiingoReaderDefaultInit(t *testing.T) {
	apiKey := "test-key"
	os.Setenv(TIINGO_API_KEY, apiKey)
	tdr, err := NewTiingoDailyReader(
		[]string{},
		TiingoReaderConfig{},
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
	tdr, err := NewTiingoDailyReader(
		[]string{"sym1", "sym2"},
		TiingoReaderConfig{
			startDate: startDate,
			endDate:   endDate,
			apiKey:    apiKey,
		},
	)

	if err != nil {
		t.Error("Error during custom intialization")
	}

	if tdr.apiKey != apiKey {
		t.Errorf("Wrong value of API key attribute.\n Expected: \t%s\n Actual: \t%s", apiKey, tdr.apiKey)
	}
}

func TestTiingoReaderEmptyAPIKeyAndNoEnv(t *testing.T) {
	startDate := time.Now().AddDate(-2, 0, 0)
	endDate := time.Now()
	apiKey := ""

	err := os.Unsetenv(TIINGO_API_KEY)
	if err != nil {
		t.Error(err)
	}

	_, err = NewTiingoDailyReader(
		[]string{"sym1", "sym2"},
		TiingoReaderConfig{
			startDate: startDate,
			endDate:   endDate,
			apiKey:    apiKey,
		},
	)

	if err == nil {
		t.Error("Expected error not raised")
	}
}
