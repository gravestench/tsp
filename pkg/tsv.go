package pkg

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
)

const (
	separator = '\t'
	expansion = "Expansion" // the diablo2 devs were smoking crack
)

// FromBytes loads a TsvParser from bytes
func Unmarshal(tsvData []byte, destination *any) error {
	// set the pipe as the delimiter for writing
	gocsv.TagSeparator = "\t"
	// set the pipe as the delimiter for reading
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	if err := gocsv.Unmarshal(bytes.NewReader(tsvData), &destination); err != nil {
		return fmt.Errorf("unmarshaling tsv data: %v", err)
	}

	return nil
}
