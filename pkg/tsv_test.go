package pkg

import (
	"fmt"
	"testing"
)

const exampleTsv = `Foo	Bar	Baz
1	2	3
expansion
4	5	6
`

type exampleTsvStruct struct {
	Foo int `csv:"Foo"`
	Bar int `csv:"Bar"`
	Baz int `csv:"Baz"`
}

func TestUnmarshal(t *testing.T) {
	var dst []exampleTsvStruct
	if err := Unmarshal([]byte(exampleTsv), &dst); err != nil {
		t.Errorf("unmarshalling error: %v", err)
	}

	fmt.Println(dst)
}
