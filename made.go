package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"os"
	"os/exec"

	"github.com/fatih/color"
)

func Run(t *Task) {
	f, err := os.CreateTemp("", "made")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name()) // clean up
	_, err = f.Write([]byte(strings.Join(t.Script, "\n")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/bash", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func main() {

	data, err := ioutil.ReadFile("Madefile")
	if err != nil {
		panic("Can't read madefile")
	}
	file := ParseString(string(data))

	if len(os.Args) <= 1 {
		printHelp(file)
	}

	for _, arg := range os.Args[1:] {
		for _, t := range file.Tasks {
			if t.Name == arg {
				Run(t)
			}
		}
	}

}

func printHelp(file *File) {
	if len(file.Tasks) == 0 {
		color.Yellow("There are not tasks defined in your Madefile")
		return
	}

	var maxTaskNameSize int
	for _, t := range file.Tasks {
		if len(t.Name) > maxTaskNameSize {
			maxTaskNameSize = len(t.Name)
		}
	}

	taskColor := color.New(color.Bold, color.FgHiGreen)
	commentColor := color.New(color.FgBlue)

	for _, task := range file.Tasks {
		taskColor.Print(task.Name)
		for i := 0; i < maxTaskNameSize-len(task.Name); i++ {
			fmt.Print(" ")
		}
		commentColor.Println(task.Comment)
	}
}
