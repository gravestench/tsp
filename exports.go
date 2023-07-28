package tsv

import (
	"github.com/gravestench/tsv/pkg"
)

type TsvParser = pkg.TsvParser

func FromBytes(data []byte) (*TsvParser, error) {
	return pkg.FromBytes(data)
}
