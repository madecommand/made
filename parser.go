package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ParseString(madefile string) (f *File) {
	f = &File{
		Tasks:   make([]*Task, 0),
		Vars:    make(map[string]string),
		Filters: make([]*Filter, 0),
	}

	f.parse(madefile)
	return f
}

func (f *File) parse(s string) {
	leftString := s

	end := strings.IndexRune(leftString, '\n')

	for end >= 0 {
		if end != 0 {
			f.parseLine(leftString[:end])
		}
		leftString = leftString[end+1:]
		end = strings.IndexRune(leftString, '\n')
	}
	f.parseLine(leftString)
}

func (f *File) parseLine(line string) {
	firstChar, _ := utf8.DecodeRuneInString(line)
	if unicode.IsSpace(firstChar) {
		f.parseSpacedLine(line)
	} else if unicode.IsLetter(firstChar) {
		f.parseLetterLine(line)
	} else if line[0] == '>' {
		f.parseFilterDefinition(line)
	} else if line[0] == '^' {
		f.parseTaskFilter(line)
	} else {
		f.parseError("Don't know how to parse line", line)
	}
}

func (f *File) parseTaskFilter(line string) {
	if f.lastTask == nil {
		f.parseError("Can't use a filter without a task.", line)
		return
	}
	fields := strings.Fields(line[1:])
	if len(fields) == 0 {
		f.parseError("Can't have an empty filter", line)
		return
	}
	tf := &TaskFilter{
		Name: fields[0],
	}
	if len(fields) > 1 {
		tf.Args = fields[1:]
	}

	f.lastTask.Filters = append(f.lastTask.Filters, tf)
}

func (f *File) parseFilterDefinition(line string) {
	name := ""
	for _, ch := range line[1:] {
		switch {
		case unicode.IsLetter(ch):
			name += string(ch)
		case ch == ':':
			break
		default:
			f.parseError("can't understand this filter: ", line)
			return
		}
	}

	filter := &Filter{
		Name:   name,
		Script: make([]string, 0),
	}
	f.Filters = append(f.Filters, filter)
	f.lastTaskOrFilter = filter

}

func (f *File) parseSpacedLine(line string) {
	trimmed := strings.Trim(line, " \t\r\n")
	if len(trimmed) == 0 {
		return
	}
	if f.lastTaskOrFilter == nil {
		f.parseError("Missing Task or Filter definition", line)
	}
	f.lastTaskOrFilter.AddLineToScript(line)
}

func (f *File) parseTaskDefinition(task, rest string) {
	t := new(Task)
	t.Name = task
	f.Tasks = append(f.Tasks, t)
	f.lastTask = t
	f.lastTaskOrFilter = t

	if index := strings.Index(rest, "##"); index > -1 {
		t.Comment = strings.Trim(rest[index+2:], " \t\r\n")
	}

}

func (f *File) parseError(err string, line string) {
	panic(fmt.Sprintf("%s: %q", err, line))
}

func (f *File) parseLetterLine(line string) {
	for i, r := range line {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' || r == '-' {
			continue
		} else if r == ':' {
			f.parseTaskDefinition(line[:i], line[i:])
			break
		} else if r == '=' {
			f.Vars[line[:i]] = strings.Trim(line[i+1:], " ")
			break
		} else {
			f.parseError("Expecting a task defition", line)
		}
	}
}
