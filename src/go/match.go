package match

import (
  "fmt"
  "container/vector"
  "./cell"
  "./board"
  "./history"
)

type ResponseType int
const (
  PUT ResponseType = 1
  PASS ResponseType = 2
  GIVEUP ResponseType = 3
  KO ResponseType = 4
  FORBIDDEN ResponseType = 5
)

type Response struct {
  Type ResponseType
  Data map[string]interface{}
}
func NewPutResponse(x int, y int, v vector.Vector) *Response {
  return &Response{PUT, map[string]interface{}{"x":x, "y":y, "taken":v}}
}
func NewPassResponse() *Response { return &Response{PASS, nil} }
func NewGiveupResponse() *Response { return &Response{GIVEUP, nil} }
func NewKoResponse() *Response { return &Response{KO, nil} }
func NewForbiddenResponse() *Response { return &Response{FORBIDDEN, nil} }

type Teban int
const (
  SENTE Teban = 0
  GOTE Teban = 1
)
func (t Teban)Color() cell.Cell {
  if t == SENTE {
    return cell.BLACK
  }
  return cell.WHITE
}

type Player interface {
  Name() string
  Teban() Teban
  Next(m *Match) *Response
}

type Status int
const (
  PLAYING = 0
  FINISH = 1
)

type Match struct {
  Turn int
  Board *board.Board
  History *history.History
  Players [2]Player
  Agehama [2]int
  Winner *Player
}

func New(b *board.Board, players [2]Player) *Match {
  match := &Match{0, b, history.New(), players, [2]int{0, 0}, nil}
  return match
}

func (m *Match)NextTurn() {
  m.Turn++
  m.Turn %= 2
}

func (m *Match)CurrentPlayer() Player {
  return m.Players[m.Turn]
}

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
      m.History.Pass(color)
      if m.History.IsLastPass() {
        // m.Winner = ?
        status = FINISH
      } else {
        status = PLAYING
        m.NextTurn()
      }
    case GIVEUP:
      m.Winner = &m.Players[(m.Turn+1)%2]
      status = FINISH
  }
  return status, resp
}

func (m *Match)Json() string {
  return fmt.Sprintf("({'board':%s, 'turn':%d, 'agehama':[%d, %d], 'version':%d})",
    m.Board.Json(), m.Turn, m.Agehama[0], m.Agehama[1], m.History.Size())
}
