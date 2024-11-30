package main

import (
	"errors"
	"fmt"
	"os"
)

func getWorkingDir() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("working dir is not specified")
	}
	return os.Args[1], nil
}

func readWorkingDir() ([]os.DirEntry, error) {
	var result []os.DirEntry
	workingDir, err := getWorkingDir()
	if err != nil {
		return result, err
	}
	entries, err := os.ReadDir(workingDir)
	if err != nil {
		return result, err
	}
	for _, e := range entries {
		if !e.IsDir() {
			result = append(result, e)
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return result, errors.New("working dir is empty")
}

func processFile(f os.File) error {
	var b []byte
	_, err := f.Read(b)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println(len(os.Args), os.Args)
}
