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

func AddToFrontOfSlice(slice []string, toAdd []string) []string {
	for _, add := range toAdd {
		slice = append(slice, "")
		copy(slice[1:], slice)
		slice[0] = add
	}
	//slice = append(toAdd, slice...)
	//copy(slice[1:], slice)
	//slice[0] = toAdd
	return slice
}

func AddStackToFrontOfSlice(slice []string, toAdd []string) []string {
	slice = append(toAdd, slice...)
	return slice
}

func RemoveFromFrontOfSlice(slice []string, amount int) (newSlice []string, removed []string) {
	//removed := make([]string, 0)
	for i := 0; i < amount; i++ {
		removed = append(removed, slice[i])
	}
	newSlice = slice[amount:]
	return
}
