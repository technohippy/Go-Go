// This package implements a point of a game board.
package cell

// A point of a game board.
type Cell int

// Contents of a cell.
const (
  SPACE Cell = 0
  BLACK Cell = 1
  WHITE Cell = 2
  OB Cell = 3
)

// Reverse returns a cell with different color.
func (c Cell)Reverse() Cell {
  if c == BLACK { return WHITE }
  if c == WHITE { return BLACK }
  return SPACE
}

func (c Cell)Is(cell Cell) bool {
  return c == cell || c == OB
}

func (c Cell)String() string {
  if c == BLACK { return "Black" }
  return "White"
}
