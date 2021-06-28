package main

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dir := "test"
	skipDir := "skip"
	if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		// Fail accessing a path
		if err != nil {
			return err
		}
		// Skip "skip" directory
		if info.IsDir() && info.Name() == skipDir {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			err := rename(path, info.Name())
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		panic(err)
	}

}

func rename(path string, filename string) error {
	p := strings.Split(path, filename)
	new := Pretty(filename)
	var b bytes.Buffer
	b.WriteString(p[0])
	b.WriteString(new)
	if err := os.Rename(path, b.String()); err != nil {
		return err
	}
	return nil
}

// Returns filename with a format: Title NNN.ext
// E.g. test_001.txt -> Test 001.txt
func Pretty(filename string) string {
	// Case name_NNN.ext
	re := regexp.MustCompile(`(_[0-9]+)`)
	s := re.Split(filename, -1)
	name := strings.Title(s[0])
	ext := s[1]
	order := re.Find([]byte(filename))
	n := strings.Split(string(order), "_")

	var b bytes.Buffer
	b.WriteString(name) // new filename
	b.WriteString(" ")
	b.WriteString(n[1]) // ordering number
	b.WriteString(ext)  // extension

	return b.String()
}
