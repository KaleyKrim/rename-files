package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getMatchedFiles(files []os.FileInfo, stringToMatch string) []os.FileInfo {
	var matchedFiles []os.FileInfo
	for _, file := range files {
		containsTarget := strings.Contains(strings.ToLower(file.Name()), strings.ToLower(stringToMatch))
		if containsTarget == false {
			continue
		}

		matchedFiles = append(matchedFiles, file)
	}

	if len(matchedFiles) == 0 {
		return nil
	}

	return matchedFiles
}

type oldToNew struct {
	Old string
	New string
}

func getOldToNew(old string, new string) oldToNew {
	return oldToNew{"./" + old, "./" + new}
}

func main() {
	args := os.Args[1:]

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("couldn't read current directory")
		return
	}

	matchedFiles := getMatchedFiles(files, args[0])
	if matchedFiles == nil {
		fmt.Println("no matching files to rename")
		return
	}

	var pathsToRename []oldToNew
	switch strings.ToLower(args[1]) {
	case "prefix":
		for _, file := range matchedFiles {
			pathsToRename = append(pathsToRename, getOldToNew(file.Name(), args[2]+"-"+file.Name()))
		}
	case "suffix":
		for _, file := range matchedFiles {
			extensionIndex := strings.Index(file.Name(), ".")
			fileName := file.Name()
			extension := ""
			if extensionIndex > -1 {
				extension = fileName[extensionIndex:len(fileName)]
				fileName = fileName[0:extensionIndex]
			}
			pathsToRename = append(pathsToRename, getOldToNew(fileName+extension, fileName+"-"+args[2]+extension))
		}
	default:
		for i, file := range matchedFiles {
			fileNum := strconv.Itoa(i + 1)
			extensionIndex := strings.Index(file.Name(), ".")
			fileName := file.Name()
			extension := ""
			if extensionIndex > -1 {
				extension = fileName[extensionIndex:len(fileName)]
				fileName = fileName[0:extensionIndex]
			}
			pathsToRename = append(pathsToRename, getOldToNew(fileName+extension, args[1]+"-"+fileNum+extension))
		}
	}

	for _, path := range pathsToRename {
		os.Rename(path.Old, path.New)
		fmt.Println("renamed " + path.Old + " to " + path.New)
	}

}
