package resource

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	"github.com/mrtazz/checkmake/parser"
)

type Action struct {
	Name        string
	Description string
}

type Resource struct {
	URL     *url.URL
	Actions []Action
}

func Open(locator string) (*Resource, error) {

	u, err := url.Parse(locator)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "file":
		return openFile(u)
	}
	return nil, fmt.Errorf("Unknown Scheme: %q for %q", u.Scheme, u)
}

func openFile(url *url.URL) (*Resource, error) {

	r := &Resource{URL: url}
	makefile, err := parser.Parse(url.Path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open Resource: %w", err)
	}

	for _, rule := range makefile.Rules {
		desc := "Runs " + rule.Target

		deps := strings.Join(rule.Dependencies, " ")
		parts := strings.Split(deps, "##")
		if len(parts) > 1 {
			desc = strings.TrimSpace(parts[1])
		}

		r.Actions = append(r.Actions, Action{
			Name:        rule.Target,
			Description: desc,
		})
	}

	return r, nil
}

func (r *Resource) Run(target string) (string, error) {
	// TODO: Ensure Working directory is correct
	out, err := r.Cmd(target).Output()
	return string(out), err
}

func (r *Resource) Cmd(target string) *exec.Cmd {
	return exec.Command("make", target)
}
