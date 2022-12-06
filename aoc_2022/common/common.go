package common

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileToLines(path string) []string {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}

func StringToInt(stringNumber string) int {
	parsed, err := strconv.ParseInt(stringNumber, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(parsed)
}
