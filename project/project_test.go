package project

import (
	"os"
	"testing"
)

func TestProject(t *testing.T) {

	wd, _ := os.Getwd()

	p, err := Open(wd)
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.Run("patatas")
	if err != nil {
		t.Fatal(err)
	}
	if out != "Muy ricas\n" {
		t.Fatalf("Expecting output to be %q. Got %q", "Muy ricas\n", out)
	}

}
