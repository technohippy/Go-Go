package cell

type Cell int
const (
  SPACE Cell = 0
  BLACK Cell = 1
  WHITE Cell = 2
  OB Cell = 3
)
func (c Cell)Reverse() Cell {
  if c == BLACK { return WHITE }
  if c == WHITE { return BLACK }
  return SPACE
}
