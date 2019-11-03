package main

import (
	"os"
	"strings"
)

// List of organizers methods
const ALPHABETICAL_M = "alphabetical" // -
const TYPE_M = "type"                 // -
const DATE_M = "date method"
const LOOKING_SAME = "looking same"

func OrganizeDir(filesInfos []os.FileInfo, ignore []string, lang string, method string) map[string][]string {
	var ext string
	var fileType FileType
	var ftName string
	var uName string
	var firstL string
	var dateStr string
	var isIgnore bool

	organizer := make(map[string][]string)

	for _, info := range filesInfos {
		isIgnore = Contains(info.Name(), ignore)

		if !info.IsDir() && !isIgnore {
			ext = GetFileExtension(info.Name())
			fileType = GetFileType(ext)
			ftName = fileType.val(lang)

			switch method {

			case ALPHABETICAL_M:
				uName = strings.ToUpper(info.Name())
				firstL = string(uName[0])
				organizer[firstL] = append(organizer[firstL], info.Name())

			case DATE_M:
				dateStr = strings.Split(info.ModTime().String(), " ")[0]
				organizer[dateStr] = append(organizer[dateStr], info.Name())

			default: //TYPE_M method
				organizer[ftName] = append(organizer[ftName], info.Name())
			}
		}
	}

	return organizer
}
