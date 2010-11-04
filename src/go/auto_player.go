package auto_player

import (
  "./match"
)

type AutoPlayer struct {
  teban match.Teban
}

func New(t match.Teban) *AutoPlayer {
  return &AutoPlayer{}
}

func (cp *AutoPlayer)Name() string {
  return "auto"
}

func (cp *AutoPlayer)Teban() match.Teban {
  return cp.teban
}

func (cp *AutoPlayer)SetTeban(t match.Teban) {
  cp.teban = t
}

func (cp *AutoPlayer)Next(m *match.Match) *match.Response {
  return match.NewGiveupResponse()
}
