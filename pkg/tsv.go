package pkg

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	separator = '\t'
)

func Unmarshal(tsvData []byte, destination any) (err error) {
	tsvData, err = filterCSVFile(tsvData)
	if err != nil {
		return fmt.Errorf("filtering csv: %v", err)
	}

	// set the pipe as the delimiter for writing
	gocsv.TagSeparator = string(separator)
	// set the pipe as the delimiter for reading
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	if err = gocsv.Unmarshal(bytes.NewReader(tsvData), destination); err != nil {
		return fmt.Errorf("unmarshaling tsv data: %v", err)
	}

	return nil
}

func filterCSVFile(input []byte) ([]byte, error) {
	// Create a new buffer to store the filtered CSV data
	var output bytes.Buffer

	// Create a CSV reader using the input []byte
	reader := csv.NewReader(bytes.NewReader(input))
	reader.Comma = separator

	// Read the first line (header) separately
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %v", err)
	}

	// Write the header to the output buffer
	output.WriteString(strings.Join(header, string(separator)) + "\n")

	// Loop through the remaining lines
	for {
		// Read each row
		row, err := reader.Read()

		// Check for end of file or any other error
		if err == io.EOF {
			break
		}

		// Check if the number of columns in the row is equal to the number in the header
		if len(row) == len(header) {
			// If the row has the correct number of columns, write it to the output buffer
			output.WriteString(strings.Join(row, string(separator)) + "\n")
		}
	}

	// Return the filtered data as a []byte
	return output.Bytes(), nil
}
