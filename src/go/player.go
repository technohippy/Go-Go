package player

import (
  "./board"
  "./history"
  "./match"
)

type Status int
const (
  PUT Status = 0
  PASS Status = 1
  GIVEUP Status = 2
  KO Status = 3
  FORBIDDEN Status = 4
)

type Response struct {
  Status Status
  Data map[string] interface{}
}

type Teban int
const (
  SENTE Teban = 0
  GOTE Teban = 1
)

type Player interface {
  Name() string
  Teban() Teban
  Next(b *board.Board, h *history.History, agehama [2]int) *Response
  Next2(m *match.Match) *Response
}
