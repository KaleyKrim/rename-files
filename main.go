package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "path/filepath"
	// "regexp"
	// "os/exec"
)

func readDir(dirname string) {

}

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

// rename (match pattern) prefix/suffix (newinput)
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
			pathsToRename = append(pathsToRename, getOldToNew(file.Name(), file.Name()+"-"+args[2]))
		}
	default:
		for i, file := range matchedFiles {
			fileNum := string(i + 1)
			pathsToRename = append(pathsToRename, getOldToNew(file.Name(), args[1]+"-"+fileNum))
		}
	}

	for _, path := range pathsToRename {
		os.Rename(path.Old, path.New)
	}

}
