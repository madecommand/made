package resource

import (
	"os"
	"testing"
)

func TestResource(t *testing.T) {

	wd, _ := os.Getwd()
	file := "file://" + wd + "/Makefile"
	r, err := Open(file)
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Actions) != 5 {
		t.Error(len(r.Actions))
	}

	out, err := r.Run("patatas")
	if err != nil {
		t.Fatal(err)
	}
	if out != "Muy ricas\n" {
		t.Fatalf("Expecting output to be %q. Got %q", "Muy ricas\n", out)
	}

}
