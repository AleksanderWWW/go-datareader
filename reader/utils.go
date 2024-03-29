/*
Copyright © 2023 Aleksander Wojnarowicz <alwojnarowicz@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func createRequest(params map[string]string, headers map[string]string, baseUrl string) (*http.Request, error) {
	parameters := url.Values{}

	var urlStr string

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}

		u, err := url.ParseRequestURI(baseUrl)
		if err != nil {
			return nil, err
		}

		u.RawQuery = parameters.Encode()
		urlStr = fmt.Sprintf("%v", u)
	} else {
		urlStr = baseUrl
	}

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

func filterDates(startDate time.Time, endDate time.Time) func(el series.Element) bool {
	return func(el series.Element) bool {
		if valStr, ok := el.Val().(string); ok {
			val, err := time.Parse("2006-01-02", valStr)
			if err != nil {
				return false
			}
			return (val.After(startDate) || val.Equal(startDate)) &&
				(val.Before(endDate) || val.Equal(endDate))
		}
		return false
	}
}

func renameDataframe(df dataframe.DataFrame, symbol string) dataframe.DataFrame {
	for _, name := range df.Names() {
		if strings.ToLower(name) == "date" {
			continue
		}
		df = df.Rename(symbol+"-"+name, name)
	}
	return df
}
