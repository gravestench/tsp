<!-- PROJECT LOGO -->
<h1 align="center">TSV</h1>
<p align="center">
  Parse Diablo 2 tab-separated value (TSV) files for storing spreadsheets of 
data
  <br />
  <br />
  <a href="https://github.com/gravestench/tsv/issues">Report Bug</a>
  ·
  <a href="https://github.com/gravestench/tsv/issues">Request Feature</a>
</p>

<!-- ABOUT THE PROJECT -->
## About

The Diablo 2 TSV Parser package provides a Go implementation for parsing 
tab-separated value (TSV) files used for storing spreadsheets of data inside 
Diablo 2 archives.

The main data structure used in this package is the `TsvParser`. It allows you 
to read and retrieve data from TSV files by column name, making it easy to 
access and process data in the TSV format.

### Caveats
The Diablo 2 TSV format has the following caveats:
- The separator used in the TSV files is a tab character.
- The TSV files may have rows with the string "Expansion," which is skipped 
during parsing.

## Usage

### Prerequisites
To use this Diablo 2 TSV Parser package, ensure you have Go 1.16 or a later 
version installed, and your Go environment is set up correctly.

### Installation
To install the package, you can use Go's standard `go get` command:

```shell
go get -u github.com/gravestench/tsv
```

### Parsing TSV Files
To parse a Diablo 2 TSV file from a byte slice, use the `FromBytes` function to 
create a `TsvParser`:

```golang
fileData := // Load your TSV file data here as a byte slice
parser, err := FromBytes(fileData)
if err != nil {
    // Handle error
}
// Use the parser object to access and process the TSV data
```

### Accessing Data
After parsing the TSV file, you can access and retrieve data using the 
`TsvParser` methods:

```golang
// Read the next row in the TSV file, skipping "Expansion" lines
for parser.Next() {
    // Access string data in a given column
    strValue := parser.String("ColumnName")

    // Access numerical data in a given column
    intValue := parser.Number("ColumnWithNumber")

    // Access a boolean value in a given column
    boolValue := parser.Bool("BooleanColumn")

    // Access a list of values from a comma-delimited column
    listValue := parser.List("CommaSeparatedColumn")

    // Process the data as needed
    // ...
}
```

### Features
The Diablo 2 TSV Parser package offers the following features:
- Efficiently parse Diablo 2 tab-separated value (TSV) files.
- Access data from the TSV file using column names.
- Handle "Expansion" lines automatically to skip them during parsing.

<!-- CONTRIBUTING -->
## Contributing

Contributions to the Diablo 2 TSV Parser package are welcome and encouraged. 
If you find any issues or have improvements to suggest, feel free to open an 
issue or submit a pull request.

To contribute to the project, follow these steps:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- MARKDOWN LINKS & IMAGES -->
[d2-tsv-parser]: https://github.com/gravestench/tsv
