package auto_player

import (
  "./board"
  "./player"
)

type AutoPlayer struct {
  teban player.Teban
}

func New(t player.Teban) *AutoPlayer {
  return &AutoPlayer{}
}

func (cp *AutoPlayer)Name() string {
  return "auto"
}

func (cp *AutoPlayer)Teban() player.Teban {
  return cp.teban
}

func (cp *AutoPlayer)Next(b *board.Board) player.Status {
  return player.GIVEUP
}
