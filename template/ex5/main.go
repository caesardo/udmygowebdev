package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date                                     time.Time
	Open, High, Low, Close, Volume, AdjClose float64
}

func main() {
	http.HandleFunc("/", handlerfunc)
	http.ListenAndServe(":8080", nil)
}

func handlerfunc(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := parsecsv("table.csv")

	// parse template
	tpl, err := template.ParseFiles("stock.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func parsecsv(path string) []Record {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))
	for i, v := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", v[0])
		open, _ := strconv.ParseFloat(v[1], 64)
		high, _ := strconv.ParseFloat(v[2], 64)
		low, _ := strconv.ParseFloat(v[3], 64)
		close, _ := strconv.ParseFloat(v[4], 64)
		volume, _ := strconv.ParseFloat(v[5], 64)
		adjclose, _ := strconv.ParseFloat(v[6], 64)

		records = append(records, Record{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			AdjClose: adjclose,
		})
	}

	return records
}
