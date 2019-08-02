package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// CSVReader reads CSV records in the specified format
type CSVReader struct{ reader io.Reader }

// ReadAll reads CSV records one by one, converts them to a WorldRecord and sends them to an ouput channer
func (csvReader *CSVReader) ReadAll() (ch chan *WorldRecord) {
	ch = make(chan *WorldRecord)
	go func() {
		defer close(ch)
		r := csv.NewReader(csvReader.reader)
		// Space separates values
		r.Comma = ' '
		// Allow variable number of fields per record
		r.FieldsPerRecord = -1
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			cityName := CityName(rec[0])
			record := &WorldRecord{City: cityName}

			roads := make([]*Road, len(rec)-1)
			for i, edge := range rec[1:] {
				// Directions and their respective destinations are separated by '='
				// e.g. North=Edinburgh
				roadTuple := strings.Split(edge, "=")
				direction := roadTuple[0]
				destCityName := CityName(roadTuple[1])
				roads[i] = NewRoad(direction, cityName, destCityName)
			}
			record.Roads = roads

			ch <- record
		}
	}()
	return
}

// NewCSVReader creates a new CSVReader which will read from the reader param
func NewCSVReader(r io.Reader) *CSVReader {
	return &CSVReader{reader: r}
}

// CSVWriter writes CSV records in the specified format
type CSVWriter struct{ writer io.Writer }

// WriteAll converts WorldRecord records one by one from the input channel to a row of CSV in the specified format and writes them
func (csvWriter *CSVWriter) WriteAll(ch <-chan *WorldRecord) {
	writer := csv.NewWriter(csvWriter.writer)
	writer.Comma = ' '
	defer writer.Flush()
	for {
		if record, ok := <-ch; ok {
			csvRecord := make([]string, len(record.Roads)+1)
			// CityName is the first column
			csvRecord[0] = string(record.City)
			// Following columns are the roads leading out of the city in the format "north=Edinburgh"
			for i, road := range record.Roads {
				csvRecord[i+1] = fmt.Sprintf("%s=%s", road.Direction, road.Destination)
			}
			writer.Write(csvRecord)
		} else {
			break
		}
	}
}

// NewCSVWriter creates a new CSVWriter which will write WorldRecord records to CSV format
func NewCSVWriter(w io.Writer) *CSVWriter {
	return &CSVWriter{writer: w}
}
