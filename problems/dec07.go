package problems

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const TERMINATOR int = 100000
const TOTAL_SPACE int = 70000000
const FREE_SPACE_NEEDED int = 30000000

type Read interface {
}

type file struct {
	owner *directory
	name  string
	size  int
}

type directory struct {
	owner   *directory
	name    string
	subdirs []*directory
	files   []file
}

func (d directory) addDir(nd *directory) directory {
	(*nd).owner = &d
	d.subdirs = append(d.subdirs, nd)
	return d
}

func (d directory) addFile(f file) directory {
	f.owner = &d
	d.files = append(d.files, f)
	return d
}

func (d directory) String(indent string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(" - %s (dir)\n", d.name))
	new_indent := indent + "  "
	for _, dir := range d.subdirs {
		builder.Write([]byte(new_indent))
		myRepr := (*dir).String(new_indent)
		builder.WriteString(myRepr)
	}

	for _, f := range d.files {
		myRepr := f.String(new_indent)
		builder.WriteString(fmt.Sprintf("%s\n", myRepr))
	}
	return builder.String()
}

func (f file) String(indent string) string {
	return fmt.Sprintf("%s - %s (file, size=%d)", indent, f.name, f.size)
}

func Problem7() {
	//raw := testInput7()
	raw, _ := getInput("input/problem7.txt")

	rootDir := directory{name: "/"}
	var workingDir *directory
	workingDir = &rootDir
	for _, cmd := range strings.Split(string(raw), "$") {
		cmd := strings.Trim(cmd, " ")
		workingDir = parseCmd(cmd, workingDir)
	}
	fmt.Print(rootDir.String(""))
	var leafDirs []*directory
	leafDirs = getLeafDirs(&rootDir, leafDirs)

	sizeMap := make(map[*directory]int)
	for _, dir := range leafDirs {
		sizeMap = (*getSizes(dir, &sizeMap))
	}
	totalSize := 0
	spaceUsed := sizeMap[&rootDir]
	space_available := TOTAL_SPACE - spaceUsed
	space_needed := FREE_SPACE_NEEDED - space_available

	fmt.Printf("UsedSpace: %d SpaceAvail: %d SpaceNeeded:  %d\n", spaceUsed, space_available, space_needed)

	var dirSizes []int
	for _, val := range sizeMap {
		dirSizes = append(dirSizes, val)
	}
	sort.Ints(dirSizes)
	for i, n := range dirSizes {
		if n > space_needed {
			idx := i - 1
			fmt.Println(dirSizes[idx])
		}
	}

	fmt.Println(totalSize)

}

func getRoot(d *directory) (root *directory) {
	for {
		if (*d).owner == nil {
			return d
		}
		return getRoot((*d).owner)
	}
}

func getSizes(d *directory, sizeMap *map[*directory]int) (outSM *map[*directory]int) {
	working := 0
	for _, f := range (*d).files {
		working += f.size
	}

	for _, sd := range (*d).subdirs {
		working += (*sizeMap)[sd]
	}
	(*sizeMap)[d] = working

	if d.owner == nil {
		return sizeMap
	}
	return getSizes(d.owner, sizeMap)
}

func getLeafDirs(d *directory, leafDirs []*directory) []*directory {
	for _, dir := range (*d).subdirs {
		if len((*dir).subdirs) < 1 {
			leafDirs = append(leafDirs, dir)
		} else {
			leafDirs = getLeafDirs(dir, leafDirs)
		}
	}
	return leafDirs
}

func buildContents(d *directory, contents []string) *directory {
	for _, line := range contents {
		elems := strings.Split(line, " ")
		if len(elems) < 2 {
			continue
		}
		//fmt.Printf("Elem: %s  Len: %d\n", elems, len(elems))
		if elems[0] == "dir" {
			(*d).subdirs = append((*d).subdirs, &directory{name: elems[1], owner: d})
		} else {
			size, err := strconv.Atoi(elems[0])
			if err != nil {
				log.Printf("ERROR CONVERTING SIZE! %s", err)
				continue
			}
			(*d).files = append((*d).files, file{name: elems[1], size: size, owner: d})
		}
	}
	return d
}

func getSubdir(d *directory, name string) (tgt *directory) {
	for _, sd := range (*d).subdirs {
		if (*sd).name == name {
			return sd
		}
	}
	log.Printf("WARNING NO DIR NAMED %s found in %s", name, d.name)
	return d
}

func parseCmd(cmd string, curDir *directory) (tgtDir *directory) {
	parsed := strings.Split(cmd, "\n")
	cmd = parsed[0]

	if strings.HasPrefix(cmd, "cd") {
		tgt := strings.Split(cmd, " ")[1]
		switch tgt {

		case "/":
			fmt.Println("going to root")
			return getRoot(curDir)
		case "..":
			fmt.Println("cd ..")
			return ((*curDir).owner)
		default:
			fmt.Printf("CDing into %s\n", tgt)
			return getSubdir(curDir, tgt)
		}
	}
	return buildContents(curDir, parsed[1:])
}

func testInput7() []byte {
	return []byte(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
dir ef
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ef
$ ls
589 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)
}
