package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	file, err := ParseString(`
ENV= production   

  
task_is-yes: ## Comment   
	echo $ENV`)
	if err != nil {
		t.Fatal(err)
	}

	// Vars
	val, ok := file.Vars["ENV"]
	if !ok {
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
