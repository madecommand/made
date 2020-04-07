package project

import (
	"errors"
	"os/exec"
	"path/filepath"

	"github.com/guillermo/made/resource"
)

type Project struct {
	base      string
	Resources []*resource.Resource
}

func Open(base string) (*Project, error) {

	p := &Project{
		base: base,
	}

	r, err := resource.Open("file://" + filepath.Join(base, "Makefile"))
	if err != nil {
		return nil, err
	}

	p.Resources = append(p.Resources, r)

	return p, nil
}

func (p *Project) Run(target string) (string, error) {
	resource, err := p.getResourceForAction(target)
	if err != nil {
		return "", err
	}
	return resource.Run(target)
}

func (p *Project) Cmd(target string) (*exec.Cmd, error) {
	resource, err := p.getResourceForAction(target)
	if err != nil {
		return nil, err
	}
	return resource.Cmd(target), nil
}

func (p *Project) getResourceForAction(target string) (*resource.Resource, error) {
	for _, r := range p.Resources {
		for _, action := range r.Actions {
			if action.Name == target {
				return r, nil
			}
		}
	}

	return nil, ErrActionNotFound
}

var ErrActionNotFound = errors.New("Action not found")
