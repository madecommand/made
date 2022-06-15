package main

type Project struct {
	Dir   string
	Files []*File
}

type File struct {
	Path  string
	Vars  map[string]string
	Tasks []*Task
}

type Scripter interface {
	AddLineToScript(string)
}

type Task struct {
	Name    string
	Comment string
	Script  []string
}

func (t *Task) AddLineToScript(line string) {
	t.Script = append(t.Script, line)
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
