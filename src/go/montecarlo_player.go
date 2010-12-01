package montecarlo_player

import (
  "container/vector"
  "./martch"
)

type MonteCarloPlayer struct {
  teban match.Teban
}

func New(t match.Teban) *MonteCarloPlayer {
  return &MonteCarloPlayer{t}
}

func (*MonteCarloPlayer)Name() string {
  return "montecarlo"
}

// Teban returns a teban of a player.
func (mp *MonteCarloPlayer)Teban() match.Teban {
  return mp.teban
}

// Next will chose a random point to put a piece.
func (mp *MonteCarloPlayer)Next(m *match.Match) *match.Response {
  color := ap.teban.Color()
  candidates := new(vector.Vector)
  size := m.Board.Size()
  for y := 0; y < size; y++ {
    for x := 0; x < size; x++ {
      if m.Board.CanPutAt(color, x, y, m.History) {
        candidates.Push(&point.Point{x, y})
      }
    }
  }
  //m.Board.CreateCheckTable()

  //virtualMatch := m.Copy()
}
