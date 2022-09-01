package main

import (
	"fmt"
	"log"
	"strings"
	"runtime"
	"os"

	"github.com/creativeprojects/go-selfupdate"
	"github.com/fatih/color"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Lshortfile)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("can't get current directory", err)
	}

	dir, err := FindProjectDir(wd)
	if err != nil {
		log.Fatal("No Madefile or .made directory found")
	}

	p, err := LoadProject(dir)
	if err != nil {
		log.Fatal("can't load the project", err)
	}

	if len(os.Args) <= 1 {
		printTasks(p, false)
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

		case "--autocomplete":
			printAutocomplete()
			return

		case "--global", "-g":
			printTasks(p, true)
		case "--update":
			if ! strings.HasPrefix(version,"v") {
				log.Println("--update option is only available in versioned release")
				log.Println("  current version:", version)
				return
			}
			err = update()
			if err != nil {
				log.Println(err)
			}
			return
		case "--version":
			fmt.Printf("made. Version(v%s) COMMIT: %10s from %s\n", version, commit, date)
			return
		case "-h", "--help":
			printHelp()
			return
		case "-t", "--tasks":
			printTasks(p, false)
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

func printAutocomplete() {
	fmt.Println(`
_made() { COMPREPLY=($(compgen -W "$(made -t | cut -f1 -d' '  | grep -v '.made' | grep -v 'Madefile')" -- "$2")); }
complete -F _made made
	`)
}

func printHelp() {
	fmt.Println(`made [OPTIONS] task
	
OPTIONS:
	--show -s    Show the generated script
	--help -h    Show the help
	--tasks -t   List the current tasks`)
}

func printTasks(p *Project, showGlobal bool) {

	tasksToDisplay := make(map[*File][]*Task)

	for _, f := range p.Files {
		for _, t := range f.Tasks {
			if t.Global == showGlobal {
				ts, ok := tasksToDisplay[f]
				if !ok {
					ts = make([]*Task, 0)
				}
				tasksToDisplay[f] = append(ts, t)
			}
		}
	}

	if len(tasksToDisplay) == 0 {
		color.Yellow("There are not tasks defined")
		color.Blue("Create some in Madefile or inside .made/my_tasks.made")
		return
	}

	showFilePrefix := showGlobal
	if len(tasksToDisplay) > 1 {
		showFilePrefix = true
	}

	for f, tasks := range tasksToDisplay {
		if showFilePrefix || showGlobal {
			if showGlobal {
				home, err := os.UserHomeDir()
				if err != nil {
					color.HiMagenta(f.Path)
				} else {
					color.HiMagenta("~/" + f.Path[len(home)+1:])
				}
			} else {
				wd, _ := os.Getwd()
				color.HiMagenta(f.Path[len(wd)+1:])
			}
		}
		var maxTaskNameSize int
		for _, t := range tasks {
			if t.Comment == "" {
				continue
			}
			if len(t.Name) > maxTaskNameSize {
				maxTaskNameSize = len(t.Name)
			}
		}
		taskColor := color.New(color.Bold, color.FgHiGreen)
		commentColor := color.New(color.FgBlue)

		for _, t := range tasks {
			if t.Comment == "" {
				continue
			}
			taskColor.Print(t.Name + " ")
			for i := 0; i < maxTaskNameSize-len(t.Name); i++ {
				fmt.Print(" ")
			}
			commentColor.Println(t.Comment)
		}
	}

}

var version string
var commit string
var date string

const github_repo = "madecommand/made"

func update() error {
	latest, found, err := selfupdate.DetectLatest(github_repo)
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %v", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not locate executable path")
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %v", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil

}
