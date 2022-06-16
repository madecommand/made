package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestProjectBuild(t *testing.T) {
	p := &Project{
		Files: []*File{
			{
				Tasks: []*Task{
					{Name: "say_hi"},
				},
			},
		},
	}
	out, err := p.BuildScript(p.Files[0].Tasks)
	if err != nil {
		t.Fatal(err)
	}
	if out != "#!/bin/sh\n\n# say_hi\n\n" {
		t.Fatalf("Expecting an empty script. Got %q", out)
	}
}

func TestProjectBuild_Task(t *testing.T) {
	p := &Project{
		Files: []*File{
			{
				Path: "madefile",
				Tasks: []*Task{
					{Name: "say_hi", Script: []string{"echo hi"}},
				},
			},
		},
	}

	out, err := p.BuildScript(p.Files[0].Tasks)
	if err != nil {
		t.Fatal(err)
	}

	if out != "#!/bin/sh\n\n# say_hi\necho hi\n" {
		t.Fatalf("Expected out to be different than %q", out)
	}

}

func testDepResolution(t *testing.T, p *Project, input, output string, shouldError bool) {

	tasks := strings.Fields(input)
	order := strings.Fields(output)

	name := fmt.Sprintf("(%s) => [%s]", input, output)

	t.Run(name, func(t *testing.T) {

		ts := []*Task{}
		for _, name := range tasks {
			task, _ := p.FindTask(name)
			if task == nil {
				t.Fatalf("Task %q not found", name)
			}
			ts = append(ts, task)

		}
		sb := &ScriptBuilder{
			p:     p,
			tasks: ts,
		}

		result, err := sb.taskOrder()
		if shouldError && err == nil {
			t.Errorf("Expected to have error. But there was none ")
		}
		if !shouldError && err != nil {
			t.Errorf("Got error %s", err)

		}

		// Compare output
		if len(result) != len(order) {
			t.Errorf("Got %v", result)
		} else {

			for i := range result {
				if result[i] != order[i] {
					t.Errorf("Got unexpected output: %v", result)
					break
				}
			}
		}
	})

}

func TestScriptBuilder_taskOrder(t *testing.T) {
	p := &Project{
		Files: []*File{
			{
				Path: "madefile",
				Tasks: []*Task{
					{Name: "A", Deps: []string{"B"}},
					{Name: "B", Deps: []string{"C"}},
					{Name: "C"},
					{Name: "D"},
					{Name: "E", Deps: []string{"D"}},
					{Name: "F", Deps: []string{"G"}},
					{Name: "G", Deps: []string{"F"}},
				},
			},
		},
	}

	testDepResolution(t, p, "A", "C B A", false)
	testDepResolution(t, p, "B", "C B", false)
	testDepResolution(t, p, "C", "C", false)
	testDepResolution(t, p, "G", "", true)
}
