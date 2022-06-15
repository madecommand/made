package main

import (
	"fmt"
	"strings"
)

func (p *Project) BuildScript(tasks []*Task) (string, error) {

	header := "#!/bin/sh\n"

	script := ""
	env := ""

	//Set env
	for _, f := range p.Files {
		for k, v := range f.Vars {
			env += fmt.Sprintf("%s=%s\n", k, v)
		}
	}

	for _, t := range tasks {
		script += fmt.Sprintf("# %s\n", t.Name)
		script += strings.Join(t.Script, "\n")
		script += "\n"
	}

	out := strings.Join([]string{header, env, script}, "\n")

	return out, nil

}
