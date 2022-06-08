package parser

import (
	"testing"
)

var input = `

# This is a comment 
TASSDF=fdasf

target: dependency ## Comment
	execution

`

func TestParser(t *testing.T) {

	_, items := Parse(input)
	for i := range items {
		t.Log(i)
	}

	t.Fatal("That's all")

}
