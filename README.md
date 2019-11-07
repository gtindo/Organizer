# Organizer

## Overview
An utilitairy command to organize files inside directories.

It organize by :
* File type (Audio, video, pictures, documents,...)
* Date : files with same last update date are placed in same folder
* Extensions : files with same extensions are regrouped in a folder
* Alphabetical order

## How to install

### Get a package release for your operating system

* Go to release page : https://github.com/gtindo/Organizer/releases
* Download appropriate file for your system
* Extract file and add binary to your system path
* Check installation  with  :
```shell
org -help
```

### Clone repository and build

* Make sure you have golang environment install in your system (https://golang.org/doc/install)

* Clone repository
````shell
git clone https://github.com/gtindo/Organizer
````

* Build project
````
cd Organizer
go build -o org
````

## How to use

Check help
````shell
org -help
````

Organize directory by type (audio, video,...) 
````shell
org -dirs <dir relative or absolute path> 
or
org -dirs <dir relative or absolute path> -method type
````

Organize multiple folders simulaneously
````
org -dirs "<dir1 path>, <dir2 path>, ..., <dirn path>"
````

Organize folder with alphabetical order
````shell
org -dirs <dir path> -method alph
````

Organize dirirectory by extensions
````shell
org -dirs <dir path> -method ext
````

Organize directory by date
````shell
org -dirs <dir path> -method date
````

Exclude some files
````shell
org -dirs <dir path> -ef "file1.txt, file2.txt, ..., filen.txt"
````

Exclude files with some extensions
````shell
org -dirs <dir path> -ex "mp3, mp4, png"
````

Translate created directories (only english, french and spanish are supported). English is default language.
````shell
org -dirs <dir path> -lang fr
org -dirs <dir path> -lang en
org -dirs <dir path> -lang es
````