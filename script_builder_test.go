package main

import "testing"

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
	if out != "#!/bin/sh\n\n\n# say_hi\n\n" {
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

	if out != "#!/bin/sh\n\n\n# say_hi\necho hi\n" {
		t.Fatalf("Expected out to be different than %q", out)
	}

}
