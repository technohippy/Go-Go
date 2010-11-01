package point

type Point [2]int

func (p *Point)X() int {
  return p[0]
}

func (p *Point)Y() int {
  return p[1]
}
