package main

import (
	"fmt"
	"log"
	"strings"

	"os"

	"github.com/fatih/color"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Lshortfile)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("can't get current directory", err)
	}
	p, err := LoadProject(wd)
	if err != nil {
		log.Fatal("can't load the project", err)
	}

	if len(os.Args) <= 1 {
		printTasks(p)
		return
	}

	show := false
	tasks := []*Task{}
	args := []string{}
	FOR:
	for i, arg := range os.Args[1:] {
		switch arg {
		case "--":
			args = os.Args[2+i:]
			break FOR
		case "--show", "-s":
			show = true
		case "-h", "--help":
			printHelp()
			return
		case "-t", "--tasks":
			printTasks(p)
			return
		default:
			if strings.HasPrefix(arg, "-") {
				log.Fatalf("option %s not recognized", arg)
			}

			t, _ := p.FindTask(arg)
			if t == nil {
				log.Fatalf("task %q not found", arg)
			}
			tasks = append(tasks, t)
		}
	}

	if show {
		script, err := p.BuildScript(tasks)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(script)
	} else {
		err = p.Run(tasks, args)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func printHelp() {
	fmt.Println(`made [OPTIONS] task
	
OPTIONS:
	--show -s    Show the generated script
	--help -h    Show the help
	--tasks -t   List the current tasks`)
}

func printTasks(p *Project) {
	for _, f := range p.Files {
		if len(f.Tasks) == 0 {
			color.Yellow("There are not tasks defined in your Madefile")
			continue
		}

		if len(p.Files) > 1 {
			wd, _ := os.Getwd()
			color.Blue(f.Path[len(wd)+1:])
		}

		var maxTaskNameSize int
		for _, t := range f.Tasks {
			if len(t.Name) > maxTaskNameSize {
				maxTaskNameSize = len(t.Name)
			}
		}

		taskColor := color.New(color.Bold, color.FgHiGreen)
		commentColor := color.New(color.FgBlue)

		for _, task := range f.Tasks {
			taskColor.Print(task.Name + " ")
			for i := 0; i < maxTaskNameSize-len(task.Name); i++ {
				fmt.Print(" ")
			}
			commentColor.Println(task.Comment)
		}
	}
}
