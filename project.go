package main

import (
	"fmt"
	"strings"
)

type Project struct {
	Dir   string
	Files []*File
}

type Var struct {
	Key, Value string
}

type File struct {
	Path  string
	Vars  []Var
	Tasks []*Task
}

func (f *File) GetVar(name string) (string, error) {
	for _, v := range f.Vars {
		if v.Key == name {
			return v.Value, nil
		}
	}
	return "", fmt.Errorf("Var not found")

}

type Task struct {
	File    *File
	Name    string
	Comment string
	Script  []string
	Deps    []string
}

func (t *Task) ScriptString() string {
	s := fmt.Sprintf("# %s\n", t.Name)
	for _, line := range t.Script {
		if strings.Index(line, "  ") == 0 {
			s += line[2:] + "\n"
			continue
		}
		if strings.Index(line, "\t") == 0 {
			s += line[1:] + "\n"
			continue
		}
		s += line + "\n"
	}

	s += "\n"
	return s
}

func (p *Project) FindTask(name string) (*Task, *File) {
	for _, f := range p.Files {
		for _, t := range f.Tasks {
			if t.Name == name {
				return t, f
			}
		}
	}
	return nil, nil

}
