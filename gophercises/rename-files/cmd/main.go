package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)


var re = regexp.MustCompile(`^(.+?) (\d{4}) \((\d+) of (\d+)\)\.(.+?)$`)

const replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "Whether or not to perform a dry run.")
	flag.Parse()

	var toRename []string
	filepath.Walk("sample", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		return nil
	})

	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, _ := match(filename)
		newPath := filepath.Join(dir, newFilename)
		if dry {
			fmt.Printf("mv %s -> %s\n", oldPath, newPath)
		} else {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}
}

func match(fn string) (string, error) {
	if !re.MatchString(fn) {
		return "", fmt.Errorf("%s didn't match the pattern", fn)
	}
	return re.ReplaceAllString(fn, replaceString), nil
}
