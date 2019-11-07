package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const OS_WINDOWS = "windows"
const BYTE_PERMS = 0775

// Return extension of a file given his name as parameter
func GetFileExtension(fileName string) string {
	tabExt := strings.Split(fileName, ".")
	n := len(tabExt)
	ext := tabExt[n-1]

	return strings.ToLower(ext)
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
	oSys := strings.Split(runtime.GOOS, "/")[0]
	cmd := new(exec.Cmd)

	if oSys == OS_WINDOWS {
		cmd = exec.Command("move", filePath, destinationPath)
	} else {
		cmd = exec.Command("mv", filePath, destinationPath)
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // standard output
	cmd.Stderr = &stderr // errors output

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

	extensions := map[FileType][]string{
		SOURCES_CODE:     SOURCE_CODE_EXT,
		SPREADSHEETS:     SPREADSHEETS_EXT,
		VIDEOS:           VIDEOS_EXT,
		PUBLISHING:       PUBLISHING_EXT,
		DOCUMENTS:        DOCUMENTS_EXT,
		PRESENTATIONS:    PRESENTATION_EXT,
		MEDIA_ARCHIVES:   MEDIA_ARCHIVES_EXT,
		AUDIO:            AUDIO_EXT,
		ARCHIVES:         ARCHIVES_EXT,
		PICTURES:         PICTURES_EXT,
		VECTOR_GRAPHICS:  VECTOR_GRAPHICS_EXT,
		_3D_GRAPHICS:     _3_D_GRAPHICS_EXT,
		SHORTCUTS:        SHORTCUTS_EXT,
		EXECUTABLES:      EXECUTABLE_FILES_EXT,
		CAD:              COMPUTER_AID_DESIGN_EXT,
		EDA:              ELECTRONIC_DESIGN_AUTOMATION_EXT,
		DATABASES:        DATABASE_EXT,
		FINANCIAL_DATAS:  FINANCIAL_DATA_EXT,
		FONTS:            FONTS_EXT,
		PLAYLIST:         PLAYLIST_EXT,
		VIRTUAL_MACHINES: VIRTUAL_MACHINES_EXT,
		WEB_PAGES:        WEB_PAGES_EXT,
	}

	for fileType, extList := range extensions {
		if Contains(ext, extList) {
			return fileType
		}
	}

	return OTHERS
}
