package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	file, err := ParseString(`
ENV= production   

  
task_is-yes: # Comment   
	echo $ENV`)
	if err != nil {
		t.Fatal(err)
	}

	// Vars
	val, err := file.GetVar("ENV")
	if err != nil {
		t.Error("ENV was not defined")
	} else if val != "production" {
		t.Error("ENV was not production. Was ", val)
	}

	//Tasks
	if len(file.Tasks) != 1 {
		t.Fatal("task was not define")
	}
	task := file.Tasks[0]

	if task.Comment != "Comment" {
		t.Errorf("Comment was not defined. It was: %q", task.Comment)
	}
	if strings.Join(task.Script, "\n") != "\techo $ENV" {
		t.Errorf("Script was not defined. It was: %q", strings.Join(task.Script, "\n"))
	}

}

func TestParser_Dependencies(t *testing.T) {
	p, err := ParseString(`
a: # Hola
  echo hola
b: a ## Comment   
  echo b`)
	if err != nil {
		t.Fatal(err)
	}

	b := &Task{}
	for _, task := range p.Tasks {
		if task.Name == "b" {
			b = task
		}
	}

	if b == nil {
		t.Fatal("Task not found")
	}

	if len(b.Deps) != 1 {
		t.Fatal("Expecting b to have dependencies")
	}
	if b.Deps[0] != "a" {
		t.Error("Expecting b to depend on a. Got", b.Deps)
	}

	if b.Comment != "Comment" {
		t.Fatal("expecting b comment. Got", b.Comment)
	}

}
