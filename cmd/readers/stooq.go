package reader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	// "github.com/go-gota/gota/dataframe"
)

var frequenciesAvailable = map[string]bool {
	"d": true,
	"w": true, 
	"m": true,
	"q": true,
	"y": true,
}

type StooqDataReader struct {
	symbols []string
	startDate time.Time
	endDate time.Time
	freq string
	baseUrl string
}

func (sdr StooqDataReader) getParams(symbol string) map[string]string{
	return map[string]string {
		"s": symbol,
		"i": sdr.freq,
		"d1": strings.Replace(sdr.startDate.Format("2006-01-02"), "-", "", -1),
		"d2": strings.Replace(sdr.endDate.Format("2006-01-02"), "-", "", -1),
	}
}

func (sdr StooqDataReader) getResponse(params map[string]string, headers map[string]string) (string, error) {
	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}

    u, err := url.ParseRequestURI(sdr.baseUrl)
	if err != nil {
		return "", err  
	}

    u.RawQuery = parameters.Encode()
    urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", err  
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

    resp, err := client.Do(req)

	if err != nil {
		return "", err  
	}

	defer resp.Body.Close()
	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err  
	}

	return string(respText), nil
}

func (srd StooqDataReader) ParseResponse(respText string) ([]SingleRecord, error) {
	lines := strings.Split(respText, "\n")
	records := make([]SingleRecord, 0, len(lines))

	for _, line := range lines[1:len(lines)-1] {
		items := strings.Split(line, ",")

		date, err := time.Parse("2006-01-02", items[0])
		if err != nil {
			return []SingleRecord{}, err
		}

		open, err := strconv.ParseFloat(items[1], 32)
		if err != nil {
			return []SingleRecord{}, err
		}

		high, err := strconv.ParseFloat(items[2], 32)
		if err != nil {
			return []SingleRecord{}, err
		}

		low, err := strconv.ParseFloat(items[3], 32)
		if err != nil {
			return []SingleRecord{}, err
		}

		close, err := strconv.ParseFloat(items[4], 32)
		if err != nil {
			return []SingleRecord{}, err
		}
		volume, err := strconv.ParseInt(strings.Replace(items[5], "\r", "", 1), 10, 32)
		if err != nil {
			return []SingleRecord{}, err
		}
		record := SingleRecord{
			date: date,
			open: open,
			high: high,
			low: low,
			close: close,
			volume: volume,
		}
		records = append(records, record)
	}
	return records, nil
}

// func (sdr StooqDataReader) read() dataframe.DataFrame {
// 	for _, symbol := range sdr.symbols {
// 		params := sdr.getParams(symbol)

// 		data, err := sdr.getResponse(params, DefaultHeaders)

// 		if err != nil {
// 			continue
// 		}


// 	}
// }


func NewStooqDataReader(symbols []string, startDate time.Time, endDate time.Time, freq string) (*StooqDataReader, error) {
	baseUrl, ok := BaseUrlMap["stooq"]

	if !ok {
		return &StooqDataReader{}, errors.New("Could not find stooq base url")
	}

	if _, ok := frequenciesAvailable[freq]; !ok {
		errMsg := fmt.Sprintf("Incorrect frequency chosen: %s", freq)
		return &StooqDataReader{}, errors.New(errMsg)
	}

	return &StooqDataReader{
		symbols: symbols,
		freq: freq,
		startDate: startDate,
		endDate: endDate,
		baseUrl: baseUrl,
	}, nil
}
