package main

import (
	"os"
	"os/exec"
)

func (p *Project) Run(tasks []*Task, args []string) error {
	f, err := os.CreateTemp("", "made")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name()) // clean up

	s, err := p.BuildScript(tasks)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	args = append([]string{f.Name()}, args...)

	cmd := exec.Command("/bin/sh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
