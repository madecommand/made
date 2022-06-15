package main

import (
	"io/ioutil"
)

func LoadProject(dir string) (*Project, error) {

	prj := &Project{
		Dir: dir,
	}

	f, err := loadFile("Madefile")
	if err != nil {
		return nil, err
	}
	prj.Files = []*File{f}

	return prj, nil

}

func loadFile(path string) (*File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	p, err := ParseString(string(data))
	if err != nil {
		return nil, err
	}

	f := &File{
		Path:  path,
		Tasks: p.Tasks,
		Vars:  p.Vars,
	}

	return f, nil
}
