package tsv

import (
	"github.com/gravestench/tsv/pkg"
)

func Unmarshal(tsvData []byte, destination any) error {
	return pkg.Unmarshal(tsvData, destination)
}
