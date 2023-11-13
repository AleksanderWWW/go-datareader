package e2e

import (
	"testing"
	"time"

	"github.com/AleksanderWWW/go-datareader/reader"
)

func TestE2e(t *testing.T) {
	stooqReader, err := reader.NewStooqDataReader(
		reader.StooqReaderConfig{
			Symbols:   []string{"PKO", "KGH", "PZU"},
			StartDate: time.Now().AddDate(0, 0, -100),
			EndDate:   time.Now(),
			Freq:      "d",
		},
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(stooqReader)
	// ---------------------------------------------------

	fredReader, err := reader.NewFredDataReader(
		reader.FredReaderConfig{
			Symbols:   []string{"SP500", "DJIA", "VIXCLS"},
			StartDate: time.Now().AddDate(0, 0, -100),
			EndDate:   time.Now(),
		},
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(fredReader)
	// ---------------------------------------------------

	bocReader, err := reader.NewBOCDataReader(
		reader.BOCReaderConfig{
			Symbols:   []string{"FXUSDCAD", "FXCADIDR", "FXCADPEN"},
			StartDate: time.Now().AddDate(0, 0, -100),
			EndDate:   time.Now(),
		},
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(bocReader)
	// ---------------------------------------------------

	startDate := time.Now().AddDate(0, 0, -100)
	endDate := time.Now()
	tiingoReader, _ := reader.NewTiingoDailyReader(
		reader.TiingoReaderConfig{
			Symbols:   []string{"ZZZOF", "000001"},
			StartDate: startDate,
			EndDate:   endDate,
		},
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(tiingoReader)
	// ---------------------------------------------------
}
