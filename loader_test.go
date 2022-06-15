package main

import (
	"testing"
)

func TestLoader(t *testing.T) {

	p, err := LoadProject("tests/dot-files")
	if err != nil {
		t.Fatal(err)
	}

	if len(p.Files) != 2 {
		t.Fatal("missing files", len(p.Files))
	}

	task, file := p.FindTask("say_hi")
	if task == nil {
		t.Fatal("Task should exits.")
	}
	if file.Path != "tests/dot-files/Madefile" {
		t.Fatal("say_hi should be from the Madefile", file.Path)
	}

	task, file = p.FindTask("say_ho")
	if task == nil {
		t.Fatal("Task should exists")
	}
	if file.Path != "tests/dot-files/.made/tasks.made" {
		t.Fatal("say_ho should be from the tasks.made", file.Path)
	}

}
