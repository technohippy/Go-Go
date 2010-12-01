// This package implements an AI for a Go game.
package random_player

import (
  "fmt"
  "crypto/rand"
  "container/vector"
  "./board"
  "./point"
  "./match"
)

// AI for a Go game.
type RandomPlayer struct {
  teban match.Teban
}

// New returns an AI for a Go game.
func New(t match.Teban) *RandomPlayer {
  return &RandomPlayer{t}
}

// Name returns a name of a player. (Now "random" is returned.)
func (*RandomPlayer)Name() string {
  return "random"
}

// Teban returns a teban of a player.
func (ap *RandomPlayer)Teban() match.Teban {
  return ap.teban
}

// Next will chose a random point to put a piece.
func (ap *RandomPlayer)Next(m *match.Match) *match.Response {
  color := ap.teban.Color()
  candidates := new(vector.Vector)
  size := m.Board.Size()
  for y := 1; y <= size; y++ {
    for x := 1; x <= size; x++ {
      if m.Board.CanPutAt(color, x, y, m.History) {
        if !m.Board.IsEye(color, x, y) {
          candidates.Push(&point.Point{x, y})
        }
      }
    }
  }

  if candidates.Len() != 0 {
    bs := make([]byte, 1)
    _, err := rand.Read(bs)
    if err == nil {
      p := candidates.At(int(bs[0]) % candidates.Len()).(*point.Point)
      ts, resp := m.Board.PutAt(color, p.X(), p.Y(), m.History)
      if resp == board.OK {
        fmt.Printf("[random] put %d,%d\n", p.X(), p.Y())
        return match.NewPutResponse(p.X(), p.Y(), ts)
      }
    }
  }
  return match.NewPassResponse()
}
