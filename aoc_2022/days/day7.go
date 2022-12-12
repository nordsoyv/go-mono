package days

import (
	"aoc_2022/common"
	"fmt"
	"strconv"
	"strings"
)

func Day07A() {
	println("Day 07 A")
	task07a("./res/day7.txt")
}

func Day07B() {
	println("Day 07 B")

}

func Day07TestA() {
	println("Day 07 test A")
	task07a("./res/day7test.txt")

}

func Day07TestB() {
	println("Day 07 test B")
}

func task07a(path string) {
	lines := common.ReadFileToLines(path)
	topNode := Directory{
		name:   "/",
		dirs:   make([]*Directory, 0),
		files:  make([]*File, 0),
		parent: nil,
	}
	currentNode := &topNode
	for _, line := range lines[1:] {
		if isCommand(line) {
			if isLs(line) {
				continue
			}
			_, command := isCd(line)
			if command == ".." {
				fmt.Println("Move up")
				currentNode = currentNode.parent
				fmt.Println("new current dir is", currentNode.name)
				continue
			} else {
				fmt.Println("Move to", command)
				for _, entry := range currentNode.dirs {
					if entry.name == command {
						currentNode = entry
						break
					}
				}
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				dir := parseDirLsLine(line)
				dir.parent = currentNode
				currentNode.dirs = append(currentNode.dirs, &dir)
			} else {
				file := parseFilesLine(line)
				file.parent = currentNode
				currentNode.files = append(currentNode.files, &file)
			}
		}
	}
	fmt.Print(topNode.Print(0))
	total := 0
	topNode.Size(&total)
	println(total)
}

func isCommand(line string) bool {
	return strings.HasPrefix(line, "$")
}

func parseDirLsLine(line string) Directory {
	return Directory{
		name:   line[4:],
		dirs:   make([]*Directory, 0),
		files:  make([]*File, 0),
		parent: nil,
	}
}

func parseFilesLine(line string) File {
	parts := strings.Split(line, " ")
	return File{
		name:   parts[1],
		size:   common.StringToInt(parts[0]),
		parent: nil,
	}

}

func isLs(line string) bool {
	return line == "$ ls"
}

func isCd(line string) (bool, string) {
	if strings.HasPrefix(line, "$ cd") {
		return true, line[5:]
	}
	return false, ""
}

type Directory struct {
	name   string
	files  []*File
	dirs   []*Directory
	parent *Directory
}

type File struct {
	name   string
	size   int
	parent *Directory
}

func (d Directory) Size(total *int) int {
	sum := 0
	for _, entry := range d.dirs {
		sum += entry.Size(total)
	}
	for _, entry := range d.files {
		sum += entry.size
	}
	fmt.Printf("Size of dir %v is: %v\n", d.name, sum)
	if sum < 100000 {
		*total += sum
	}
	return sum
}

func (d Directory) Print(indent int) string {
	fmt.Println("Printing dir", d.name)
	fmt.Println(len(d.dirs))
	res := ""
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}
	res += indentStr + d.name + "\n"
	for _, entry := range d.dirs {
		res += entry.Print(indent + 1)
	}
	for _, entry := range d.files {
		res += entry.Print(indent + 1)
	}
	return res
}

func (f File) Print(indent int) string {
	fmt.Println("Printing file", f.name)
	res := ""
	for i := 0; i < indent; i++ {
		res += "  "
	}
	res += f.name + " " + strconv.FormatInt(int64(f.size), 10) + "\n"
	return res
}
