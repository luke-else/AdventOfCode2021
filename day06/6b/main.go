package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	content := returnContent("../input")
	//content := returnContent("testInput")
	fmt.Println(content)

	days := 256

	var waitTime [9]int

	for _, v := range *content {
		waitTime[v]++
	}

	fmt.Println(waitTime)

	for gen := 0; gen < days; gen++ {
		justBred := waitTime[0]

		for daysToWait := range waitTime[:len(waitTime)-1] {
			//Move array down a position
			waitTime[daysToWait] = waitTime[daysToWait+1]
		}

		//Add new fish and fish that have just bred (fallen off end of array)
		waitTime[6] += justBred //start waiting other 6 days
		waitTime[8] = justBred  //new fishes
	}

	var numFishes int

	for _, fishes := range waitTime {
		numFishes += fishes
	}

	fmt.Println(numFishes)
}

func returnContent(path string) *[]int {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content []int

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		for _, v := range text {
			num, _ := strconv.Atoi(v)
			content = append(content, num)
		}
	}

	return &content
}
