package main

import (
	"fmt"
	"os"
	"strconv"
)

const argsLength = 4

type ProgArgs struct {
	workingDir string
	width      uint64
	height     uint64
}

func (p *ProgArgs) populate() error {
	if len(os.Args) < argsLength {
		return fmt.Errorf("should be %d arguments", argsLength)
	}
	var err error = nil
	p.workingDir = os.Args[1]
	p.width, err = strconv.ParseUint(os.Args[2], 10, 64)
	if err == nil {
		p.height, err = strconv.ParseUint(os.Args[3], 10, 64)
	}
	return err
}
