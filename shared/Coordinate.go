package shared

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Neighbours(gridWidth int, gridHeight int, diagonal bool) (out []Coordinate) {
	spaceLeft := c.X > 0
	spaceRight := c.X < gridWidth-1
	spaceUp := c.Y > 0
	spaceDown := c.Y < gridHeight-1

	if spaceLeft {
		out = append(out, Coordinate{c.X - 1, c.Y})
	}
	if spaceRight {
		out = append(out, Coordinate{c.X + 1, c.Y})
	}
	if spaceUp {
		out = append(out, Coordinate{c.X, c.Y - 1})
	}
	if spaceDown {
		out = append(out, Coordinate{c.X, c.Y + 1})
	}

	if diagonal {
		if spaceUp && spaceLeft {
			out = append(out, Coordinate{c.X - 1, c.Y - 1})
		}
		if spaceUp && spaceRight {
			out = append(out, Coordinate{c.X + 1, c.Y - 1})
		}
		if spaceDown && spaceLeft {
			out = append(out, Coordinate{c.X - 1, c.Y + 1})
		}
		if spaceDown && spaceRight {
			out = append(out, Coordinate{c.X + 1, c.Y + 1})
		}
	}

	return
}
