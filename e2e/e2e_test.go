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
	// ---------------------------------------------------

	fredReader, err := reader.NewFredDataReader(
		[]string{"SP500", "DJIA", "VIXCLS"},
		time.Now().AddDate(0, 0, -100),
		time.Now(),
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(fredReader)
	// ---------------------------------------------------

	bocReader, err := reader.NewBOCDataReader(
		[]string{"FXUSDCAD", "FXCADIDR", "FXCADPEN"},
		time.Now().AddDate(0, 0, -100),
		time.Now(),
	)

	if err != nil {
		t.Error(err)
	}

	reader.GetData(bocReader)
	// ---------------------------------------------------

	startDate := time.Now().AddDate(0, 0, -100)
	endDate := time.Now()
	tiingoReader, _ := reader.NewTiingoDailyReader(
		[]string{"ZZZOF", "000001"},
		reader.TiingoReaderConfig{
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
