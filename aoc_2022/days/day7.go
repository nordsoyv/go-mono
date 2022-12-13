package days

import (
	"aoc_2022/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day07A() {
	println("Day 07 A")
	size := task07a("./res/day7.txt")
	fmt.Println("Size of dir under 100000 is", size, "Expected", 1454188)
}

func Day07B() {
	println("Day 07 B")
	sizeInfo := task07b("./res/day7.txt")
	fmt.Println("Found dir", sizeInfo.name, sizeInfo.size, "Correct is wvq")
}

func Day07TestA() {
	println("Day 07 test A")
	size := task07a("./res/day7test.txt")
	fmt.Println("Size of dir under 100000 is", size, "Expected", 95437)
}

func Day07TestB() {
	println("Day 07 test B")
	sizeInfo := task07b("./res/day7test.txt")
	fmt.Println("Found dir", sizeInfo.name, sizeInfo.size, "Correct is d")
}

func task07a(path string) int {
	lines := common.ReadFileToLines(path)
	topNode := parseInput(lines)
	fmt.Print(topNode.Print(0))
	total := 0
	topNode.Size(&total)
	return total
}

func task07b(path string) SizeInfo {
	lines := common.ReadFileToLines(path)
	topNode := parseInput(lines)
	fmt.Print(topNode.Print(0))
	total := TotalSizeInfo{infos: make([]SizeInfo, 0)}
	topNode.GetSizeInfo(&total)
	for _, info := range total.infos {
		fmt.Println(info.name, info.size)
	}
	usedSpace := total.infos[len(total.infos)-1].size
	fmt.Println("Used space is\t", usedSpace)
	freeSpace := 70000000 - usedSpace
	fmt.Println("Free space is\t", freeSpace)
	neededSpace := 30000000 - freeSpace
	fmt.Println("Needed space is\t", neededSpace)
	sort.Slice(total.infos, func(i, j int) bool {
		return total.infos[i].size < total.infos[j].size
	})
	for i := 0; i < len(total.infos); i++ {
		if total.infos[i].size > neededSpace {
			return total.infos[i]
		}
	}
	return SizeInfo{
		name: "",
		size: 0,
	}
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

type SizeInfo struct {
	name string
	size int
}

type TotalSizeInfo struct {
	infos []SizeInfo
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

func (d Directory) GetSizeInfo(total *TotalSizeInfo) SizeInfo {
	sum := 0
	for _, entry := range d.dirs {
		info := entry.GetSizeInfo(total)
		sum += info.size
	}
	for _, entry := range d.files {
		sum += entry.size
	}
	fmt.Printf("Size of dir %v is: %v\n", d.name, sum)
	sizeInfo := SizeInfo{
		name: d.name,
		size: sum,
	}
	total.infos = append(total.infos, sizeInfo)
	return sizeInfo
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

func parseInput(lines []string) Directory {
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
	return topNode
}
