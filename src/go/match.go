package match

import (
  "fmt"
  "container/vector"
  "./cell"
  "./board"
  "./player"
  "./history"
)

type Status int
const (
  PLAYING = 0
  FINISH = 1
)

type Match struct {
  turn int
  board *board.Board
  history *history.History
  players [2]player.Player
  agehama [2]int
  Winner *player.Player
}

func New(b *board.Board, players [2]player.Player) *Match {
  match := &Match{0, b, history.New(), players, [2]int{0, 0}, nil}
  return match
}

func (m *Match)NextTurn() {
  m.turn++
  m.turn %= 2
}

func (m *Match)Next() (Status, player.Status) {
  p := m.players[m.turn]
  var color cell.Cell
  if p.Teban() == player.SENTE { color = cell.BLACK } else { color = cell.WHITE }
  resp := p.Next(m.board, m.history, m.agehama)
  var status Status
  switch resp.Status {
    case player.PUT:
      fmt.Printf("%v\n", resp.Data)
      taken := resp.Data["taken"].(vector.Vector)
      m.history.Add(color, resp.Data["x"].(int), resp.Data["y"].(int), taken)
      m.agehama[m.turn] += len(taken)
      status = PLAYING
      m.NextTurn()
    case player.PASS:
      if m.history.IsLastPass() {
        m.history.Pass(color)
        // m.Winner = ?
        status = FINISH
      } else {
        m.history.Pass(color)
        status = PLAYING
        m.NextTurn()
      }
    case player.GIVEUP:
      m.Winner = &m.players[(m.turn+1)%2]
      status = FINISH
  }
  return status, resp.Status
}
