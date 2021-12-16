package shared

type Board struct {
	Values [5][5]BoardValue
	Hash   map[int]Location
}

type BoardValue struct {
	Value   int
	Visited bool
}

type Location struct {
	X int
	Y int
}
