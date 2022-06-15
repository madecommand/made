package main

import "testing"

func TestProjectBuild(t *testing.T) {
	p := &Project{}
	out, err := p.Build([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if out != "#!/bin/sh\n" {
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

	out, err := p.Build([]string{"say_hi"})
	if err != nil {
		t.Fatal(err)
	}

	if out != "#!/bin/sh\necho hi\n" {
		t.Fatalf("Expected out to be different than %q", out)
	}

}

func TestProjectBuild_Filter(t *testing.T) {
	p := &Project{
		Files: []*File{
			{
				Path: "madefile",
				Tasks: []*Task{
					{
						Name:   "say_hi",
						Script: []string{"echo hi"},
					},
				},
			},
		},
	}

	out, err := p.Build([]string{"say_hi"})
	if err != nil {
		t.Fatal(err)
	}
	if out != "" {
		t.Fatal(out)
	}

}
