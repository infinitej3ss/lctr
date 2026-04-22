package main

import (
	"bufio"
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lctr <directory>")
		os.Exit(1)
	}
}

func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}
