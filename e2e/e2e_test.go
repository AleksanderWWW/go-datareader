package e2e

import (
	"testing"
	"time"

	"github.com/AleksanderWWW/go-datareader/reader"
)

func TestE2e(t *testing.T) {
	stooqReader, err := reader.NewStooqDataReader(
		[]string{"PKO", "KGH", "PZU"},
		time.Now().AddDate(0, 0, -100),
		time.Now(),
		"d",
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(stooqReader)
}
