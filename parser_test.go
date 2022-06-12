package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	file := ParseString(`
ENV= production   
>strict:
  set -$1
  echo $SCRIPT
task_is-yes: ## Comment   
^strict eux
	echo $ENV`)

	// Vars
	val, ok := file.Vars["ENV"]
	if !ok {
		t.Error("ENV was not defined")
	} else if val != "production" {
		t.Error("ENV was not production. Was ", val)
	}

	// Filters
	if len(file.Filters) == 0 {
		t.Fatal("A filter should have been defined")
	}
	filter := file.Filters[0]
	if filter.Name != "strict" {
		t.Error("Filter should be 'strict', got:", filter.Name)
	}
	if strings.Join(filter.Script, "\n") != "  set -$1\n  echo $SCRIPT" {
		t.Errorf("Filter Script is incorrect: %q", strings.Join(filter.Script, "\n"))
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

	if len(task.Filters) == 0 {
		t.Fatal("task should have a filter")
	}

	taskFilter := task.Filters[0]

	if taskFilter.Name != "strict" {
		t.Error("Expected strict as taskfilter. Got:", taskFilter.Name)
	}

	if len(taskFilter.Args) > 0 && taskFilter.Args[0] != "eux" {
		t.Error("Expected strict not to be:", taskFilter.Args)
	}

}
