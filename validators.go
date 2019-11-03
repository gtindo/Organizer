package main

import "io/ioutil"

// check if all dirs are directories
func AreDirectories(dirs []string) (bool, string) {
	var err error

	for _, val := range dirs {
		_, err = ioutil.ReadDir(val)

		if err != nil {
			return false, val + " Is not a directory"
		}

	}

	return true, ""
}

// check if given lang is supported
func isLangCorrect(lang string) (bool, string) {
	if !Contains(lang, LANG) {
		return false, lang + " is not supported"
	}
	return true, ""
}

// check if given method is correct
func isMethodCorrect(method string) (bool, string) {
	methods := []string{ALPHABETICAL_M, TYPE_M, EXTENSION_M, DATE_M}

	if Contains(method, methods) {
		return true, ""
	}

	return false, method + " Is incorrect, refer to help in order to see valid methods"
}

// check if given archive method is correct
func isArchiveMethodCorrect(method string) (bool, string) {
	methods := []string{TAR, ZIP}

	if Contains(method, methods) {
		return true, ""
	}
	return false, method + " Is incorrect."
}

// check if given compress method is corect
func isCompressAlgoCorrect(algo string) (bool, string) {
	methods := []string{BZIP2, FLATE, GZIP, LZW, ZLIB}

	if Contains(algo, methods) {
		return true, ""
	}
	return false, algo + " Is incorrect."
}
