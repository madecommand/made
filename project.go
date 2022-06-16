package main

import (
	"fmt"
	"strings"
)

type Project struct {
	Dir   string
	Files []*File
}

type File struct {
	Path  string
	Vars  map[string]string
	Tasks []*Task
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
	s += strings.Join(t.Script, "\n")
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
