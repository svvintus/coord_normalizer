package main

import (
	"fmt"
	"os"
)

func run(args ProgArgs) error {
	size := ImageSize{uint(args.width), uint(args.height)}
	dir, err := CreateDirContent(args.workingDir, ImageSize{uint(args.width), uint(args.height)})
	if err != nil {
		return err
	}
	for _, f := range dir.files {
		err := f.Normalize(size)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fmt.Println(os.Args)
	args := new(ProgArgs)
	err := args.populate()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = run(*args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
