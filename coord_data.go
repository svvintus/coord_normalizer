package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const elemsLimit = 5

type RecordData struct {
	class   int64
	x_point int64
	y_point int64
	width   int64
	height  int64
}

type NormalizedData struct {
	class   int64
	x_point float32
	y_point float32
	width   float32
	height  float32
}

type ImageSize struct {
	width  uint
	height uint
}

func CreateRecordData(line []byte) (RecordData, error) {
	var r RecordData
	err := r.parse(line)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (p RecordData) Normalize(size ImageSize) (NormalizedData, error) {
	if size.width == 0 || size.height == 0 {
		return NormalizedData{}, errors.New("width or height should not be 0")
	}
	fmt.Printf("Data to normalize: %d %d %d %d %d \n", p.class, p.x_point, p.y_point, p.width, p.height)
	return NormalizedData{p.class,
			float32(p.x_point) / float32(size.width),
			float32(p.y_point) / float32(size.height),
			float32(p.width) / float32(size.width),
			float32(p.height) / float32(size.height)},
		nil
}

func (p *RecordData) parse(line []byte) error {
	re, err := regexp.Compile(`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	if err != nil {
		return err
	}

	elements := re.FindStringSubmatch(string(line))
	fmt.Println("Parsed groups: ", elements)
	if len(elements) < elemsLimit {
		return fmt.Errorf("the line %s has wrong format", line)
	}
	return p.populate(elements[1:])
}

func (p *RecordData) populate(elements []string) error {
	if len(elements) != elemsLimit {
		return fmt.Errorf("there should be %d elements per line. Current line len is: %d", elemsLimit, len(elements))
	}
	var err error
	p.class, err = strconv.ParseInt(elements[0], 10, 64)
	p.x_point, err = strconv.ParseInt(elements[1], 10, 64)
	p.y_point, err = strconv.ParseInt(elements[2], 10, 64)
	p.width, err = strconv.ParseInt(elements[3], 10, 64)
	p.height, err = strconv.ParseInt(elements[4], 10, 64)

	if err != nil {
		return fmt.Errorf("Wrong element class %s", elements[0])
	}

	fmt.Printf("Parsed data: %d %d %d %d %d \n", p.class, p.x_point, p.y_point, p.width, p.height)
	return nil
}

func (p NormalizedData) toString() string {
	return fmt.Sprintf("%d %f %f %f %f", p.class, p.x_point, p.y_point, p.width, p.height)
}
