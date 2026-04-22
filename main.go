package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	CheckArgs() // Make sure dirname was passed in

	baseDir := os.Args[1] // 1. Grab dirname

	// 2. Iterate through all files recursively, storing filenames in a list
	var searchQueue []string
	searchQueue = append(searchQueue, baseDir)

	var readableFiles []string

	// Call a function on every subfile to add it to readableFiles if it's not a Dir
	filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		Check(err)

		// Skip directories starting with a dot, unless it's the current directory ( . )
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") && len(d.Name()) > 1 {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			// Disregard dotfiles and binaries
			if d.Name()[0] != '.' && strings.Contains(d.Name(), ".") {
				readableFiles = append(readableFiles, path)
			}
		}

		return nil
	})

	// 3. Loop through list, opening each filivke and counting its total lines, adding their counts to an ongoing sum

	totalLinesOfCode := 0
	for _, f := range readableFiles {
		lines, _ := CountLines(f)
		totalLinesOfCode += lines
	}

	// 4. Print TotalLinesOfCode
	fmt.Println("Total Lines: ", totalLinesOfCode)
}

// wow low comments!
