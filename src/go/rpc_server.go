package rpc_server

import (
  "./board"
  "./match"
)

type RPCServer int

func New(b *board.Board, ps [2]match.Player) *RPCServer {
  return new(RPCServer)
}

func (rs *RPCServer)Start(port int) {
}
