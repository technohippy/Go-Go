package player

import (
  "./board"
)

type Status int
const (
  PUT Status = 0
  PASS Status = 1
  GIVEUP Status = 2
)

type Teban int
const (
  SENTE Teban = 0
  GOTE Teban = 1
)


type Player interface {
  Name() string
  Teban() Teban
  Next(b *board.Board) Status
}
