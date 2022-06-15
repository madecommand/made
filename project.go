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

type Task struct {
	File    *File
	Name    string
	Comment string
	Script  []string
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
