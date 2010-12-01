// This package implements a Go match.
package match

import (
  "fmt"
  "container/vector"
  "./cell"
  "./board"
  "./history"
)

// Response type
type ResponseType int

// Response types
const (
  PUT ResponseType = 1
  PASS ResponseType = 2
  GIVEUP ResponseType = 3
  KO ResponseType = 4
  FORBIDDEN ResponseType = 5
)

// Response of a player's move.
type Response struct {
  Type ResponseType
  Data map[string]interface{}
}

// NewPutResponse returns a response with the PUT reponse type.
func NewPutResponse(x int, y int, v vector.Vector) *Response {
  return &Response{PUT, map[string]interface{}{"x":x, "y":y, "taken":v}}
}

// NewPassResponse returns a response with the PASS reponse type.
func NewPassResponse() *Response { return &Response{PASS, nil} }

// NewGiveupResponse returns a response with the GIVEUP reponse type.
func NewGiveupResponse() *Response { return &Response{GIVEUP, nil} }

// NewKoResponse returns a response with the KO reponse type.
func NewKoResponse() *Response { return &Response{KO, nil} }

// NewForbiddenResponse returns a response with the FORBIDDEN reponse type.
func NewForbiddenResponse() *Response { return &Response{FORBIDDEN, nil} }

// Teban
type Teban int
const (
  SENTE Teban = 0
  GOTE Teban = 1
)

// Color returns the piece's color for the taben.
func (t Teban)Color() cell.Cell {
  if t == SENTE {
    return cell.BLACK
  }
  return cell.WHITE
}

// Player
type Player interface {
  Name() string
  Teban() Teban
  Next(m *Match) *Response
}

// Match status
type Status int

// Match statuses
const (
  PLAYING = 0
  FINISH = 1
)

// Go match
type Match struct {
  Turn int
  Board *board.Board
  History *history.History
  Players [2]Player
  Agehama [2]int
  Winner *Player
}

// New returns a Go match.
func New(b *board.Board, players [2]Player) *Match {
  match := &Match{0, b, history.New(), players, [2]int{0, 0}, nil}
  return match
}

// NextTurn will set turn next.
func (m *Match)NextTurn() {
  m.Turn++
  m.Turn %= 2
}

// CurrentPlayer returns the current playing player.
func (m *Match)CurrentPlayer() Player {
  return m.Players[m.Turn]
}

// Next will ask a nest hand to a player.
func (m *Match)Next() (Status, *Response) {
  p := m.CurrentPlayer()
  var color cell.Cell
  if p.Teban() == SENTE { color = cell.BLACK } else { color = cell.WHITE }

  resp := p.Next(m)
  var status Status
  switch resp.Type {
    case PUT:
      takenPieces := resp.Data["taken"].(vector.Vector)
      m.History.Add(color, resp.Data["x"].(int), resp.Data["y"].(int), takenPieces)
      m.Agehama[m.Turn] += len(takenPieces)
      status = PLAYING
      m.NextTurn()
    case PASS:
      if m.History.IsLastPass() {
        m.History.Pass(color)
        // m.Winner = ?
        status = FINISH
      } else {
        m.History.Pass(color)
        status = PLAYING
        m.NextTurn()
      }
    case GIVEUP:
      m.Winner = &m.Players[(m.Turn+1)%2]
      status = FINISH
  }
  return status, resp
}

func (m *Match)Judge() (cell.Cell, int, int) {
  blackAreas, whiteAreas := m.Board.CountAreas()
  if blackAreas < whiteAreas {
    return cell.WHITE, blackAreas, whiteAreas
  } else if whiteAreas < blackAreas {
    return cell.BLACK, blackAreas, whiteAreas
  }
  return cell.OB, blackAreas, whiteAreas
}

// Json returns a string form of a match.
func (m *Match)Json() string {
  return fmt.Sprintf("({'board':%s, 'turn':%d, 'agehama':[%d, %d], 'version':%d})",
    m.Board.Json(), m.Turn, m.Agehama[0], m.Agehama[1], m.History.Size())
}
