package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	findEnd("", "start", *content, []string{"start"}, true)
	fmt.Println(counter)

}

//Cave Class
type Cave struct {
	name   string
	next   []string
	isBig  bool
	isEdge bool
}

var counter = 0

func findEnd(path string, cave string, caves map[string]*Cave, visited []string, oneSmall bool) {
	if cave == "end" {
		// fmt.Println(path + "," + "end")
		counter++
		return
	}
	for _, c := range (*caves[cave]).next {
		if !caves[c].isBig {
			if isIn, _ := isInSlice(c, visited); !isIn {
				findEnd(path+","+cave, c, caves, append(visited, c), oneSmall)
			} else if oneSmall && !caves[c].isEdge {
				findEnd(path+","+cave, c, caves, visited, false)
			}
		} else {
			findEnd(path+","+cave, c, caves, visited, oneSmall)
		}
	}
}

func isInSlice(target string, slice []string) (bool, int) {
	for i, s := range slice {
		if s == target {
			return true, i
		}
	}
	return false, 0
}

func returnContent(path string) *map[string]*Cave {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	content := make(map[string]*Cave)

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		from, to := line[0], line[1]
		parseCave(from, to, content)
		parseCave(to, from, content)
	}

	return &content
}

func parseCave(cave string, to string, caves map[string]*Cave) {
	if target, ok := caves[cave]; !ok {
		caves[cave] = &Cave{
			name:   cave,
			next:   []string{to},
			isBig:  strings.ToUpper(cave) == cave,
			isEdge: cave == "start",
		}
	} else {
		target.next = append(target.next, to)
	}

}
