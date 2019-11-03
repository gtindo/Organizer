package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("========== Organizer ==========")
	infos, err := ioutil.ReadDir("./folder_test")
	fmt.Println(infos)
	fmt.Println(err)

	organise := make(map[string][]string)

	organise = OrganizeDir(infos, []string{}, "en", DATE_M)

	fmt.Println(organise)

}

// organize folder by types of file (extensions)
// images, musics, videos, docs, packages, execs, books,
// archives, fonts,

// organize folder by alphabetical order

// By name looking same

// By last modification date

// archive folders, select archive type

// work on multiple folders simultanouely

// ignore some extensions

// ignore some files

// lang choice
