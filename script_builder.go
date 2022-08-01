package main

import (
	"fmt"

	"github.com/stevenle/topsort"
)

func (p *Project) BuildScript(tasks []*Task) (string, error) {
	sb := &ScriptBuilder{
		tasks: tasks,
		p:     p,
	}

	return sb.Render()
}

type ScriptBuilder struct {
	tasks []*Task
	p     *Project
}

func (sb *ScriptBuilder) taskOrder() ([]string, error) {
	if len(sb.tasks) == 0 {
		return []string{}, nil
	}

	graph := topsort.NewGraph()
	const root = "  --root"

	var addDeps func(t *Task) error

	addDeps = func(t *Task) error {
		for _, dep := range t.Deps {
			err := graph.AddEdge(t.Name, dep)
			if err != nil {
				return err
			}

			task, _ := sb.p.FindTask(dep)
			if task == nil {
				return fmt.Errorf("task %q, required by %q not found", dep, t.Name)
			}

			// Check if there are not ciruclar dependencies
			_, err = graph.TopSort(root)
			if err != nil {
				return err
			}
			err = addDeps(task)
			if err != nil {
				return err
			}
		}

		return nil
	}

	graph.AddNode(root)
	for i, task := range sb.tasks {
		graph.AddEdge(root, task.Name)
		if i != 0 {
			err := graph.AddEdge(task.Name, sb.tasks[i-1].Name)
			if err != nil {
				return []string{}, err
			}
		}
		err := addDeps(task)
		if err != nil {
			return []string{}, err
		}
	}

	result, err := graph.TopSort(root)
	return result[:len(result)-1], err
}

func (sb *ScriptBuilder) genTasksCode() (string, error) {
	runDeps, err := sb.taskOrder()
	if err != nil {
		return "", err
	}
	script := ""
	// Add scripts
	for _, dep := range runDeps {
		task, _ := sb.p.FindTask(dep)
		script += task.ScriptString()
	}

	return script, nil
}

func (sb *ScriptBuilder) Render() (string, error) {
	out := "#!/bin/sh\n"

	env, err := sb.genEnv()
	if err != nil {
		return "", err
	}
	out += env + "\n"

	tasks, err := sb.genTasksCode()
	if err != nil {
		return "", err
	}

	out += tasks

	return out, nil
}

func (sb *ScriptBuilder) genEnv() (string, error) {
	env := ""
	//Set env
	for _, f := range sb.p.Files {
		for _, v := range f.Vars {
			env += fmt.Sprintf("%s=%s\n", v.Key, v.Value)
		}
	}

	return env, nil
}
