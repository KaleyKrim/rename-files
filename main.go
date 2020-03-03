package main

import (
	"os"
	"strings"
	"io/ioutil"
	// "path/filepath"
	// "regexp"
	// "os/exec"
)

func readDir(dirname string) {

}

func getMatchedFiles(files io.FileInfo, stringToMatch string) []os.FileInfo {
	var matchedFiles []os.FileInfo
	for _, file := range files {
		matched, err := regexp.MatchString(strings.ToLower(file.Name), strings.ToLower(stringToMatch))
		if (err != nil) {
			continue
		}

		matchedFiles = append(matchedFiles, file)
	}

	return matchedFiles
}

// rename (match pattern) prefix/suffix (newinput)
func main() {
	args := os.Args[1:]

	files, err := ioutil.ReadDir(".")
	if (err != nil) {
		fmt.Println("couldn't read current directory")
		return
	}

	matchedFiles := getMatchedFiles(files, args[0])

	switch strings.ToLower(args[1]) {
	case "prefix":
		// add args[2] to start of each
	case "suffix";
		// add args[2] to end of each
	default:
		// get amount of matching files
		// name them all with args[1] and 
		// classic rename with numbers

	}
}

// rename all with new name, numbered 1~length
// rename all with new prefix
// rename all with new suffix