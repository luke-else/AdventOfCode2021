package main

type Board struct {
	Values [5][5]boardValue
	Hash   map[int]location
}

type boardValue struct {
	Value   int
	Visited bool
}

type location struct {
	x int
	y int
}
