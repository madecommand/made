package main

import "strings"

func (p *Project) Build(tasks []string) (string, error) {

	out := "#!/bin/sh\n"

	for _, taskName := range tasks {

		for _, file := range p.Files {
			for _, task := range file.Tasks {
				if task.Name == taskName {

					if len(task.Filters) > 0 {

						for 

					}

					out += strings.Join(task.Script, "\n")
					out += "\n"

				}
			}
		}
	}

	return out, nil

}
