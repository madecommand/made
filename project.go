package main

import (
	"fmt"
	"path"
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
	Global  bool
}

func (t *Task) ScriptString() string {
	s := fmt.Sprintf("# %s\n", t.Name)
	s += strings.Join(t.Script, "\n")
	s += "\n"
	return s
}

func (p *Project) FindTask(name string) (*Task, *File) {
	var globalTask, madeTask, dotMadeTask *Task
	for _, f := range p.Files {
		for _, t := range f.Tasks {
			if t.Name == name {
				switch {
				case path.Base(f.Path) == "Madefile":
					madeTask = t
				case t.Global:
					globalTask = t
				default:
					dotMadeTask = t
				}
			}
		}
	}
	if madeTask != nil {
		return madeTask, madeTask.File
	}
	if dotMadeTask != nil {
		return dotMadeTask, dotMadeTask.File
	}
	if globalTask != nil {
		return globalTask, globalTask.File
	}
	return nil, nil

}
