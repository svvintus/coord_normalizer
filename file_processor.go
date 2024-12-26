package main

import (
	"bufio"
	"fmt"
	"os"
)

type ImageFile struct {
	filePath string
	data     []RecordData
}

type NormalizedImageFile struct {
	filePath string
	data     []NormalizedData
}

func CreateNormImageFile(image ImageFile, normSize ImageSize) (NormalizedImageFile, error) {
	nf := NormalizedImageFile{filePath: image.filePath}
	for _, r := range image.data {
		norm, err := r.Normalize(normSize)
		if err != nil {
			return nf, err
		}
		nf.data = append(nf.data, norm)
	}
	return nf, nil
}

func CreateImageFile(filePath string) (ImageFile, error) {
	p := ImageFile{filePath: filePath}
	err := p.Read()
	return p, err
}

func (p *ImageFile) Read() error {
	var err error
	fmt.Println("Read file :", p.filePath)
	fl, err := os.OpenFile(p.filePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		var err error
		p.data, err = appendRecordData(p.data, scanner)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ImageFile) Normalize(size ImageSize) error {
	fmt.Println("Normalize file: ", p.filePath)
	fmt.Println("Image size: ", size)
	fmt.Println("Data length to normalize: ", len(p.data))
	nf, err := CreateNormImageFile(p, size)
	if err != nil {
		return err
	}
	return nf.Write()
}

func (p NormalizedImageFile) Write() error {
	fmt.Println("Write file: ", p.filePath)
	var err error
	fl, err := os.OpenFile(p.filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	fmt.Println("Lenght data to write: ", len(p.data))
	for _, r := range p.data {
		out := r.toString()
		fmt.Println(out)
		_, err := fl.WriteString(out + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func appendRecordData(records []RecordData, scanner *bufio.Scanner) ([]RecordData, error) {
	fmt.Println(scanner.Text())
	entry, err := CreateRecordData(scanner.Bytes())
	if err != nil {
		return nil, err
	}
	return append(records, entry), nil
}
