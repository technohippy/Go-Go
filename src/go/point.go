// This package implemtns a point of a game board.
package point

// Point
type Point [2]int

// X returns a x coordination.
func (p *Point)X() int {
  return p[0]
}

// X returns a y coordination.
func (p *Point)Y() int {
  return p[1]
}
