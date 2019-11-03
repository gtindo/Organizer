package main

import (
	"os"
	"strings"
)

// List of organizers methods
const ALPHABETICAL_M = "alph"
const TYPE_M = "type"
const DATE_M = "date"
const EXTENSION_M = "ext"
const LOOKING_SAME = "same" // For future release

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

			case EXTENSION_M:
				organizer[ext] = append(organizer[ext], info.Name())

			default: //TYPE_M method
				organizer[ftName] = append(organizer[ftName], info.Name())
			}
		}
	}

	return organizer
}
