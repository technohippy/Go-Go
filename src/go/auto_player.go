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

func (cp *ConsolePlayer)SetTeban(t player.Teban) {
  cp.teban = t
}

func (cp *AutoPlayer)Next(b *board.Board) player.Status {
  return player.GIVEUP
}

func (cp *AutoPlayer)Next2(b *board.Board, h *history.History, agehama [2]int) Status {
  return cp.Next(b)
}
