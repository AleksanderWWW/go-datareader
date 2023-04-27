package reader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func parseStooqLine(line string, symbol string) (SingleRecord, error) {
	items := strings.Split(line, ",")

	open, err := strconv.ParseFloat(items[1], 32)
	if err != nil {
		return SingleRecord{}, err
	}

	high, err := strconv.ParseFloat(items[2], 32)
	if err != nil {
		return SingleRecord{}, err
	}

	low, err := strconv.ParseFloat(items[3], 32)
	if err != nil {
		return SingleRecord{}, err
	}

	close, err := strconv.ParseFloat(items[4], 32)
	if err != nil {
		return SingleRecord{}, err
	}
	volume, err := strconv.ParseInt(strings.Replace(items[5], "\r", "", 1), 10, 32)
	if err != nil {
		return SingleRecord{}, err
	}
	return SingleRecord{
		Symbol: symbol,
		Date:   items[0],
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: int(volume),
	}, nil
}

func createRequest(params map[string]string, headers map[string]string, baseUrl string) (*http.Request, error) {
	parameters := url.Values{}

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}
	}

	u, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = parameters.Encode()
	urlStr := fmt.Sprintf("%v", u)

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func getResponse(params map[string]string, headers map[string]string, baseUrl string) (string, error) {
	req, err := createRequest(params, headers, baseUrl)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
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
