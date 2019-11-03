package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const OS_WINDOWS = "windows"
const BYTE_PERMS = 664

// Return extension of a file given his name as parameter
func GetFileExtension(fileName string) string {
	tabExt := strings.Split(fileName, ".")
	n := len(tabExt)
	ext := tabExt[n-1]

	return ext
}

// Check if a table contains a given string
func Contains(val string, tab []string) bool {
	check := false
	for _, value := range tab {
		if val == value {
			check = true
			break
		}
	}

	return check
}

func mv(filePath, destinationPath string) error {
	oSys := strings.Split(runtime.GOARCH, "/")[0]
	cmd := new(exec.Cmd)

	if oSys == OS_WINDOWS {
		cmd = exec.Command("move", filePath, destinationPath)
	} else {
		cmd = exec.Command("mv", filePath, destinationPath)
	}

	err := cmd.Run()

	return err
}

func MoveFiles(basePath string, organiser map[string][]string) error {
	var err error
	var workDir string
	var destPath string
	var currentPath string

	for dir, fileNames := range organiser {

		workDir = filepath.Join(basePath, dir)

		if _, err = os.Stat(workDir); os.IsNotExist(err) {
			err = os.Mkdir(workDir, BYTE_PERMS)
			if err != nil {
				return err
			}
		}

		for _, fileName := range fileNames {
			currentPath = filepath.Join(basePath, fileName)
			destPath = filepath.Join(workDir, fileName)
			err = mv(currentPath, destPath)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Return file type given his extension
func GetFileType(ext string) FileType {
	if Contains(ext, SOURCE_CODE_EXT) {
		return SOURCES_CODE
	}

	if Contains(ext, SPREADSHEETS_EXT) {
		return SPREADSHEETS
	}

	if Contains(ext, VIDEOS_EXT) {
		return VIDEOS
	}

	if Contains(ext, PUBLISHING_EXT) {
		return PUBLISHING
	}

	if Contains(ext, DOCUMENTS_EXT) {
		return DOCUMENTS
	}

	if Contains(ext, PRESENTATION_EXT) {
		return PRESENTATIONS
	}

	if Contains(ext, AUDIO_EXT) {
		return AUDIO
	}

	if Contains(ext, ARCHIVES_EXT) {
		return ARCHIVES
	}

	if Contains(ext, MEDIA_ARCHIVES_EXT) {
		return MEDIA_ARCHIVES
	}

	if Contains(ext, PICTURES_EXT) {
		return PICTURES
	}

	if Contains(ext, VECTOR_GRAPHICS_EXT) {
		return VECTOR_GRAPHICS
	}

	if Contains(ext, _3_D_GRAPHICS_EXT) {
		return _3D_GRAPHICS
	}

	if Contains(ext, SHORTCUTS_EXT) {
		return SHORTCUTS
	}

	if Contains(ext, EXECUTABLE_FILES_EXT) {
		return EXECUTABLES
	}

	if Contains(ext, COMPUTER_AID_DESIGN_EXT) {
		return CAD
	}

	if Contains(ext, ELECTRONIC_DESIGN_AUTOMATION_EXT) {
		return EDA
	}

	if Contains(ext, DATABASE_EXT) {
		return DATABASES
	}

	if Contains(ext, FINANCIAL_DATA_EXT) {
		return FINANCIAL_DATAS
	}

	if Contains(ext, FONTS_EXT) {
		return FONTS
	}

	if Contains(ext, PLAYLIST_EXT) {
		return PLAYLIST
	}

	if Contains(ext, VIRTUAL_MACHINES_EXT) {
		return VIRTUAL_MACHINES
	}

	if Contains(ext, WEB_PAGES_EXT) {
		return WEB_PAGES
	}

	return OTHERS
}
