package pkg

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	separator = '\t'
	expansion = "Expansion" // the diablo2 devs were smoking crack
)

// TsvParser represents a file that contains tab-separated values.
type TsvParser struct {
	lookup map[string]int
	r      *csv.Reader
	record []string
	Err    error
}

// FromBytes loads a TsvParser from bytes
func FromBytes(buf []byte) (*TsvParser, error) {
	cr := csv.NewReader(bytes.NewReader(buf))
	cr.Comma = separator
	cr.ReuseRecord = true

	fieldNames, err := cr.Read()
	if err != nil {
		return nil, fmt.Errorf("reading field names: %v", err)
	}

	parser := &TsvParser{
		lookup: make(map[string]int, len(fieldNames)),
		r:      cr,
	}

	for index, name := range fieldNames {
		parser.lookup[name] = index
	}

	return parser, nil
}

// Next reads the next row, skips "Expansion" lines or
// returns false when the end of a file is reached or an error occurred
func (d *TsvParser) Next() bool {
	var err error
	d.record, err = d.r.Read()

	if err == io.EOF {
		return false
	} else if err != nil {
		d.Err = err
		return false
	}

	if d.record[0] == expansion {
		return d.Next()
	}

	return true
}

// String gets a string from the given column
func (d *TsvParser) String(field string) string {
	return d.record[d.lookup[field]]
}

// Number gets a number for the given column
func (d *TsvParser) Number(field string) int {
	n, err := strconv.Atoi(d.String(field))
	if err != nil {
		return 0
	}

	return n
}

// List splits a delimited list from the given column
func (d *TsvParser) List(field string) []string {
	str := d.String(field)
	return strings.Split(str, ",")
}

// Bool gets a bool value for the given column
func (d *TsvParser) Bool(field string) bool {
	return d.Number(field) == 1
}
