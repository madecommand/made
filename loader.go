package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// FindProjectDir scans current directory and goes up until find a project
func FindProjectDir(dir string) (string, error) {
	// Convert to abs
	if !path.IsAbs(dir) {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			return "", fmt.Errorf("project not found: %w", err)
		}
		dir = absDir
	}

	if isDir(dir, ".made") || isFile(dir, "Madefile") {
		return dir, nil
	}

	if dir == "/" {
		return "", fmt.Errorf("project not found")
	}

	return FindProjectDir(filepath.Dir(dir))

}

func isDir(paths ...string) bool {
	p := path.Join(paths...)

	dir, err := os.Stat(p)
	return err == nil && dir.IsDir()
}

func isFile(paths ...string) bool {
	p := path.Join(paths...)

	file, err := os.Stat(p)
	return err == nil && !file.IsDir()
}

func LoadProject(dir string) (*Project, error) {

	prj := &Project{
		Dir: dir,
	}

	f, err := loadFile(path.Join(dir, "Madefile"))
	if err == nil {
		prj.Files = []*File{f}
	}

	// Load .made files
	files, err := loadDirectory(path.Join(dir, ".made"))
	if err != nil {
		return nil, err
	}
	prj.Files = append(prj.Files, files...)

	// Load ~/.made files
	config, err := os.UserConfigDir()
	if err != nil {
		log.Println("Can't load global tasks:", err)
	} else {
		files, err = loadDirectory(path.Join(config, "made"))
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			for _, t := range f.Tasks {
				t.Global = true
			}
		}
	}

	prj.Files = append(prj.Files, files...)

	return prj, nil

}

func loadDirectory(dir string) ([]*File, error) {
	d, err := os.Open(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	entries, err := d.Readdir(0)
	if err != nil {
		return nil, err
	}

	files := []*File{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".made") {
			f, err := loadFile(path.Join(dir, entry.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, f)
		}
	}

	return files, nil

}

func loadFile(path string) (*File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	p, err := ParseString(string(data))
	if err != nil {
		return nil, err
	}

	f := &File{
		Path:  path,
		Tasks: p.Tasks,
		Vars:  p.Vars,
	}

	for _, t := range f.Tasks {
		t.File = f
	}

	return f, nil
}
