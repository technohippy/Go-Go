package history

import (
  "container/vector"
  "./cell"
  "./point"
)

type Record struct {
  color cell.Cell
  x int
  y int
  takenPoints vector.Vector
}

type History struct {
  history vector.Vector
}

func New() *History {
  return &History{vector.Vector{}}
}

func (h *History)IsKou(color cell.Cell, x int, y int) bool {
  if len(h.history) == 0 { return false }

  last := h.Last()
  if len(last.takenPoints) == 1 {
    takenPoint := last.takenPoints.Last().(point.Point)
    if takenPoint.X() == x && takenPoint.Y() == y {
      return true
    }
  }
  return false
}
func (h *History)Add(color cell.Cell, x int, y int, takenPoints vector.Vector) {
  record := Record{color, x, y, takenPoints}
  h.history.Push(record)
}

func (h *History)Pass(color cell.Cell) {
  h.Add(color, -1, -1, nil)
}

func (h *History)Last() Record {
  return h.history.Last().(Record)
}

func (h *History)IsLastPass() bool {
  return h.history.Last().(Record).x == -1
}
