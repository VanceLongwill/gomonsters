package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestCSVReaderAndWriter(t *testing.T) {
	csvData := `
Denalmo north=Agixo-A south=Amolusnisnu east=Elolesme west=Migina
Asnu north=Ago-Mo south=Emexisno east=Dinexe west=Amiximine
Esmosno north=Emexege south=Anegu east=Axesminilmo west=Dosmolixo
`
	reader := strings.NewReader(csvData)
	csvReader := NewCSVReader(reader)
	testChannel := csvReader.ReadAll()

	var b bytes.Buffer
	output := bufio.NewWriter(&b)
	csvWriter := NewCSVWriter(output)
	csvWriter.WriteAll(testChannel)

	if strings.TrimSpace(b.String()) != strings.TrimSpace(csvData) {
		t.Errorf("Expected input and output csv to match")
	}
}
