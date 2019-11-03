package main

import (
	"fmt"
	"os"
	"strings"
)

const VALIDATION_ERROR = -1

func main() {
	var check bool
	var msg string

	InitArgs()

	dirs := strings.Split(Dirs, ", ")

	// directories validation
	check, msg = AreDirectories(dirs)
	if !check {
		fmt.Println(msg)
		os.Exit(VALIDATION_ERROR)
	}

	// lang validation
	check, msg = isLangCorrect(Lang)
	if !check {
		fmt.Println(msg)
		os.Exit(VALIDATION_ERROR)
	}

	// method validation
	check, msg = isMethodCorrect(Method)
	if !check {
		fmt.Println(msg)
		os.Exit(VALIDATION_ERROR)
	}

	// TODO for future release
	// Archive validation
	// Compress validation

	for _, dir := range dirs {
		err := MoveFiles(dir, OrganizeDir(
			dir,
			strings.Split(IgnoresFiles, ", "),
			strings.Split(IgnoresExt, ", "),
			Lang, Method))

		if err != nil {
			fmt.Println("An error occured while processing ", err)
		}
	}
}
