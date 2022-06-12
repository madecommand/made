package main

type Project struct {
	Dir   string
	Files []*File
}

type File struct {
	Name    string
	Vars    map[string]string
	Filters []*Filter
	Tasks   []*Task

	lastTask         *Task
	lastTaskOrFilter Scripter
}

type Scripter interface {
	AddLineToScript(string)
}

type Filter struct {
	Name   string
	Script []string
}

func (f *Filter) AddLineToScript(line string) {
	f.Script = append(f.Script, line)
}

type Task struct {
	Name    string
	Comment string
	Script  []string
	Filters []*TaskFilter
}

func (t *Task) AddLineToScript(line string) {
	t.Script = append(t.Script, line)
}

type TaskFilter struct {
	Name string
	Args []string
}
