package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")
	fmt.Println(content)

	startingString := (*content)[0]
	pairs := make(map[string]string)
	count := make(map[string]int)

	//Create map of pair values
	for i := 2; i < len(*content); i++ {
		split := strings.Split((*content)[i], " -> ")
		pairs[split[0]] = split[1]
	}

	//Fill initial list
	list := LinkedList{Head: &Node{Value: string(startingString[0]), Next: nil}}
	current := list.Head
	for i := 1; i < len(startingString); i++ {
		node := &Node{Value: string(startingString[i]), Next: nil}
		count[string(startingString[i])]++
		current.Next = node
		current = node
	}

	//Run iterations on list
	iterations := 10
	for i := 1; i <= iterations; i++ {
		current = list.Head

		for current.Next != nil {
			value := pairs[current.Value+current.Next.Value]
			list.InsertItem(current, current.Next, value)
			count[value]++
			current = current.Next.Next
		}
		list.PrintList()
	}

	//determine min and max
	min := math.MaxInt
	max := 0
	for _, value := range count {
		if value > max {
			max = value
		} else if value < min {
			min = value
		}
	}

	fmt.Println(max - min)

}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertItem(first *Node, second *Node, value string) {
	first.Next = &Node{Value: value, Next: second}
}

func (l *LinkedList) PrintList() {
	list := []string{}
	current := l.Head
	for current != nil {
		list = append(list, current.Value)
		current = current.Next
	}
	fmt.Println(list)
}

type Node struct {
	Value string
	Next  *Node
}

func returnContent(path string) *[]string {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content []string

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return &content
}
