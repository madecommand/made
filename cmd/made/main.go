package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/guillermo/made/project"
	. "github.com/logrusorgru/aurora"
)

func main() {
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p, err := project.Open(wd)
	if err != nil {
		panic(err)
	}

	args := flag.Args()
	if len(args) == 0 {
		showActions(p)
		os.Exit(0)
	}

	action := args[0]
	cmd, err := p.Cmd(action)
	if err != nil {
		if err == project.ErrActionNotFound {
			fmt.Println()
			fmt.Println(Sprintf(Red("Action %s not found."), Red(Bold(action))))
			os.Exit(-1)
		}
		panic(err)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			panic(err)
		}
	}

	os.Exit(cmd.ProcessState.ExitCode())

}

func showActions(p *project.Project) {
	maxLen := 0
	for _, r := range p.Resources {
		for _, a := range r.Actions {
			if len(a.Name) > maxLen {
				maxLen = len(a.Name)
			}

		}
	}
	strLen := strconv.Itoa(maxLen + 1)
	strFmt := "%-" + strLen + "s"

	for _, r := range p.Resources {
		for _, a := range r.Actions {
			fmt.Printf("%s %s\n",
				Magenta(fmt.Sprintf(strFmt, a.Name)),
				Blue(a.Description),
			)
		}
	}

}
