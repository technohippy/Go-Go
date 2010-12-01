// This package implements 
package rpc_player

import (
//  "fmt"
//  "./board"
//  "./point"
  "./match"
)

type RPCPlayer struct {
  name string
  teban match.Teban
}

func New(name string, t match.Teban) *RPCPlayer {
  return &RPCPlayer{name, t}
}

// Name returns a name of a player. (Now "rpc" is returned.)
func (rp *RPCPlayer)Name() string {
  return rp.name
}

// Teban returns a teban of a player.
func (rp *RPCPlayer)Teban() match.Teban {
  return rp.teban
}

// Next will chose a random point to put a piece.
func (ap *RPCPlayer)Next(m *match.Match) *match.Response {
  return nil
}
