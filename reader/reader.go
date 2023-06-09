package reader

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"

	"github.com/go-gota/gota/dataframe"
)

const LogsDirpath string = "logs"

const LogsFilePermission fs.FileMode = 0777

type DataReader interface {
	getName() string
	getSymbols() []string
	readSingle(symbol string) (dataframe.DataFrame, error)
	concatDataframes(dfs []dataframe.DataFrame) dataframe.DataFrame
}

var BaseUrlMap = map[string]string{
	"stooq": "https://stooq.com/q/d/l",
	"fred":  "https://fred.stlouisfed.org/graph/fredgraph.csv",
	"boc":   "http://www.bankofcanada.ca/valet/observations",
}

var DefaultHeaders = map[string]string{
	"Connection":                "keep-alive",
	"Expires":                   "-1",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
}

func GetData(reader DataReader) dataframe.DataFrame {
	if _, err := os.Stat(LogsDirpath); os.IsNotExist(err) {
		os.Mkdir(LogsDirpath, LogsFilePermission)
	}

	loggerName := getLoggerName(reader.getName())

	loggerPath, err := os.OpenFile(LogsDirpath+"/"+loggerName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, LogsFilePermission)
	if err != nil {
		fmt.Printf("Error setting up logging %s. Logs will not be saved", err)
	}

	defer loggerPath.Close()

	errorLogger := log.New(loggerPath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Println("SOme messae")

	symbols := reader.getSymbols()
	results := make([]dataframe.DataFrame, 0, len(symbols))
	var wg sync.WaitGroup

	for _, symbol := range symbols {

		wg.Add(1)

		go func(symbol string) {
			defer wg.Done()

			singleDf, err := reader.readSingle(symbol)
			if err != nil {
				errorLogger.Println(symbol, err)
				return
			}

			results = append(results, singleDf)
		}(symbol)
	}

	wg.Wait()

	return reader.concatDataframes(results)
}
