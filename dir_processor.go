package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirContent struct {
	dir   string
	files []ImageFile
}

func CreateDirContent(dir string, size ImageSize) (DirContent, error) {
	files, err := getDirFileEntries(dir)
	if err != nil {
		return DirContent{}, err
	}
	content := DirContent{dir: dir}
	content.AddFiles(files, size)
	return content, nil
}

func (p *DirContent) AddFiles(files []os.DirEntry, imageSize ImageSize) {
	for _, fl := range files {
		err := p.appendFile(fl)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *DirContent) appendFile(file os.DirEntry) error {
	fileData, err := CreateImageFile(filepath.Join(p.dir, file.Name()))
	if err != nil {
		return err
	}
	p.files = append(p.files, fileData)
	return nil
}

func (p DirContent) Normalize(size ImageSize) error {
	for _, fl := range p.files {
		err := fl.Normalize(size)
		if err != nil {
			return err
		}
	}
	return nil
}

func appendFile(files []os.DirEntry, entry os.DirEntry) []os.DirEntry {
	if entry.IsDir() {
		return files
	}
	return append(files, entry)
}

func getFiles(entries []os.DirEntry) []os.DirEntry {
	var files []os.DirEntry
	for _, en := range entries {
		files = appendFile(files, en)
	}
	return files
}

func getDirFileEntries(dir string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []os.DirEntry{}, err
	}
	return getFiles(entries), nil
}
