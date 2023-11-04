package reader

import (
	"os"
	"testing"
)

func TestTiingoReaderDefaultInit(t *testing.T) {
	os.Setenv(TIINGO_API_KEY, "test-key")
	_, err := NewTiingoReader(
		[]string{},
		nil,
		nil,
		nil,
	)

	if err != nil {
		t.Error("Error while default initialization")
	}
}
