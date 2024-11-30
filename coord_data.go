package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type SourceData struct {
	class   int64
	x_point int64
	y_point int64
	width   int64
	height  int64
}

func (p *SourceData) parse(line []byte) error {
	re, err := regexp.Compile(`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	if err != nil {
		return err
	}

	elements := re.FindAllString(string(line), -1)
	if len(elements) < 1 {
		return fmt.Errorf("the line %s has wrong format", line)
	}
	p.transform(elements[1:])

	return nil
}

func (p *SourceData) transform(elements []string) error {
	const elemsLimit = 5
	if len(elements) < elemsLimit {
		return fmt.Errorf("there should be %d elements per line", elemsLimit)
	}
	var err error
	p.class, err = strconv.ParseInt(elements[0], 10, 64)
	p.x_point, err = strconv.ParseInt(elements[1], 10, 64)
	p.y_point, err = strconv.ParseInt(elements[2], 10, 64)
	p.width, err = strconv.ParseInt(elements[4], 10, 64)

	if err != nil {
		return fmt.Errorf("Wrong element class %s", elements[0])
	}

	return nil
}

func (p *SourceData) parseElement(element string) error {
	i, err := strconv.ParseInt(element, 10, 8)
	if err != nil {
		return fmt.Errorf("Wrong element class %s", element)
	}
	p.class = i
	return nil
}
