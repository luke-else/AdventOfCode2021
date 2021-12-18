package main

import (
	"AdventOfCode2021/shared"
)

func main() {
	//content := returnContent("../input")
	//content := returnContent("../testInput")

	list := []int{5, 1, 3, 5, 4, 6, 2, 7, 9, 8, 0}

	tree := shared.BinaryTree{}

	for _, v := range list {
		tree.Insert(v)
	}

	tree.InOrder(tree.Head)

	//Had no Idea where to even start with this challenge

	//Ideally wanted to use a binary tree
	//will be looking to make use of generics when they release in Go verison 1.18

}

// func returnContent(path string) *[]int {
// 	//read file and return it as an array of integers

// 	file, err := os.Open(path)
// 	var content []int

// 	if err != nil {
// 		fmt.Println("Unlucky, the file didn't open")
// 		return &content
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	regex, _ := regexp.Compile(`[-+]?[\d]+`)

// 	strings := []string{}

// 	for scanner.Scan() {
// 		strings = regex.FindAllString(scanner.Text(), 4)
// 	}

// 	for _, val := range strings {
// 		num, _ := strconv.Atoi(val)
// 		content = append(content, num)
// 	}

// 	return &content
// }
