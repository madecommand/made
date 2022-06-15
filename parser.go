package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Parser struct {
	Tasks    []*Task
	Vars     map[string]string
	lastTask *Task
}

func ParseString(madefile string) (*Parser, error) {
	p := &Parser{
		Tasks: make([]*Task, 0),
		Vars:  make(map[string]string),
	}

	err := p.parse(madefile)
	return p, err
}

func (p *Parser) parse(s string) error {
	leftString := s

	end := strings.IndexRune(leftString, '\n')

	for end >= 0 {
		if end != 0 {
			err := p.parseLine(leftString[:end])
			if err != nil {
				return err
			}
		}
		leftString = leftString[end+1:]
		end = strings.IndexRune(leftString, '\n')
	}
	p.parseLine(leftString)
	return nil
}

func (p *Parser) parseLine(line string) error {
	firstChar, _ := utf8.DecodeRuneInString(line)
	if unicode.IsSpace(firstChar) {
		p.parseSpacedLine(line)
	} else if unicode.IsLetter(firstChar) {
		p.parseLetterLine(line)
	} else {
		return fmt.Errorf("don't know how to parse line %q", line)
	}
	return nil
}

func (p *Parser) parseSpacedLine(line string) error {
	trimmed := strings.Trim(line, " \t\r\n")
	if len(trimmed) == 0 {
		return nil
	}
	if p.lastTask == nil {
		return fmt.Errorf("missing task")
	}
	p.lastTask.Script = append(p.lastTask.Script, line)

	return nil

}

func (p *Parser) parseTaskDefinition(task, rest string) {
	t := new(Task)
	t.Name = task
	p.Tasks = append(p.Tasks, t)
	p.lastTask = t

	if index := strings.Index(rest, "##"); index > -1 {
		t.Comment = strings.Trim(rest[index+2:], " \t\r\n")
	}

}

func (p *Parser) parseError(err string, line string) {
	panic(fmt.Sprintf("%s: %q", err, line))
}

func (p *Parser) parseLetterLine(line string) {
	for i, r := range line {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' || r == '-' {
			continue
		} else if r == ':' {
			p.parseTaskDefinition(line[:i], line[i:])
			break
		} else if r == '=' {
			p.Vars[line[:i]] = strings.Trim(line[i+1:], " ")
			break
		} else {
			p.parseError("Expecting a task defition", line)
		}
	}
}
