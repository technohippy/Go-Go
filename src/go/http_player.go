// This package implements a user interface for web.
package http_player

import (
  "fmt"
  "./board"
  "./match"
)

// Web UI.
type HttpPlayer struct {
  name string
  teban match.Teban
  x int
  y int
}

// New returns a user interface for web.
func New(n string, t match.Teban) *HttpPlayer {
  return &HttpPlayer{n, t, 0, 0}
}

// Name returns a name of a player.
func (hp *HttpPlayer)Name() string {
  return hp.name
}

// Teban returns a teban of a player.
func (hp *HttpPlayer)Teban() match.Teban {
  return hp.teban
}

// Next will decide the next play.
func (hp *HttpPlayer)Next(m *match.Match) *match.Response {
  color := hp.teban.Color()
  takenOffs, ok := m.Board.PutAt(color, hp.x, hp.y, m.History)
  if ok == board.OK {
    fmt.Printf("[%s] Put (%d, %d) and Take %d\n", hp.name, hp.x, hp.y, len(takenOffs))
    return match.NewPutResponse(hp.x, hp.y, takenOffs)
  } else if ok == board.KO {
    fmt.Printf("[%s] Ko\n", hp.name)
    return match.NewKoResponse()
  } else if ok == board.FORBIDDEN {
    fmt.Printf("[%s] Forbidden\n", hp.name)
    return match.NewForbiddenResponse()
  }
  return nil // never reached
}

// SetNext will set a next play.
func (hp *HttpPlayer)SetNext(x int, y int) {
  hp.x = x
  hp.y = y
}
