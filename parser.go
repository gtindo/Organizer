package main

import (
	"flag"
)

var Dirs string
var Lang string
var Method string
var Archive string
var Compress string
var IgnoresFiles string
var IgnoresExt string

func InitArgs() {
	flag.StringVar(&Dirs, "dirs", "",
		`List of directories to organize
		Example for single dir : -dirs /home/username/Downloads
		For multiple dirs : -dirs "/home/username/Downloads, docs/files,"`)

	flag.StringVar(&Lang, "lang", "en",
		`Language of created folders
		Suported languages are en (english), fr (french), es (spanish)
		By default lang is set to english
		Example : -lang en
		`)

	flag.StringVar(&Method, "method", "type",
		`Method used to organize files inside directories
		This params take one of these four values : type, date, ext or alph
			type : Files are regrouped regarding their types (videos, audios, docs, pictures,...)
			date : Files are regrouped by their last update date (yyyy-mm-dd)
			ext : Files are regrouped by their extension (mp3, docs, psd, ...)
			alph : Files are regrouped using alphabetical order (A, B, C, ...)

		Example : -method type
		`)

	flag.StringVar(&IgnoresFiles, "ef", "",
		`Exclude some files while organising directories
		Example for one file : --ignore-file file.txt
		Example for mulitple files : --ignore-file "file1.txt, file2.txt, file3.txt"
	`)

	flag.StringVar(&IgnoresExt, "ex", "",
		`Excludes files with given extensions
		Example for one extension : --ignore-ext txt
		Example for mulitple extensions : --ignore-file "mp3, txt, json"
	`)

	flag.StringVar(&Archive, "archive", "zip",
		`(Not working yet)
		Archive created folders with specified method
		Suported methods are zip and tar
		Example : -archive zip
	`)

	flag.StringVar(&Compress, "compress", "gzip",
		`(Not working yet)
		Compress created folders with specified algorithm
		Suported algorithms are : bzip2, flate, gzip, lzw, zlib 
		Example : -archive gzip
	`)

	flag.Parse()
}
