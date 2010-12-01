// Ref. http://en.wikipedia.org/wiki/Go_(game)
package main

import (
  "os"
  "fmt"
  "flag"
  "./match"
  "./board"
  "./http_server"
  "./rpc_server"
  "./console_player"
  "./http_player"
  "./rpc_player"
  "./random_player"
)

// startConsole will start a console for a Go game.
func startConsole(b *board.Board, useAI bool) {
  var players [2]match.Player
  if useAI {
    players = [2]match.Player{
      //console_player.New("ando", match.SENTE),
      random_player.New(match.SENTE),
      random_player.New(match.GOTE)}
  } else {
    players = [2]match.Player{
      console_player.New("ando", match.SENTE),
      console_player.New("yasushi", match.GOTE)}
  }
  m := match.New(b, players)

  fmt.Printf("%v", b)
  for {
    matchStatus, playerResponse := m.Next()
    switch playerResponse.Type {
      case match.PUT:
        fmt.Printf("%v", b)
      case match.KO:
        fmt.Printf("Ko", b)
      case match.FORBIDDEN:
        fmt.Printf("Forbidden", b)
      case match.PASS:
        if matchStatus == match.FINISH {
          fmt.Println("Finish!")
          winner, blackArea, whiteArea := m.Judge()
          fmt.Printf("%v won. (%d:%d)\n", winner, blackArea, whiteArea)
          os.Exit(0)
        } else {
          fmt.Println("Pass")
        }
      case match.GIVEUP:
        fmt.Printf("%s Win!\n", (*m.Winner).Name())
        os.Exit(0)
    }
  }
}

// startConsole will start a server for a Go game.
func startHTTPServer(port int, b *board.Board, useAI bool) {
  var players [2]match.Player
  if useAI {
    players = [2]match.Player{
      http_player.New("ando", match.SENTE),
      random_player.New(match.GOTE)}
  } else {
    players = [2]match.Player{
      http_player.New("ando", match.SENTE),
      http_player.New("yasushi", match.GOTE)}
  }
  s := http_server.New(b, players)
  s.Start(port)
}

func startRPCServer(port int, b *board.Board, useAI bool) {
  var players [2]match.Player
  if useAI {
    players = [2]match.Player{
      rpc_player.New("ando", match.SENTE),
      random_player.New(match.GOTE)}
  } else {
    players = [2]match.Player{
      rpc_player.New("ando", match.SENTE),
      rpc_player.New("yasushi", match.GOTE)}
  }
  s := rpc_server.New(b, players)
  s.Start(port)
}

// main
func main() {
  server := flag.Bool("server", false, "run server or not")
  port := flag.Int("port", 55555, "port number")
  size := flag.Int("size", 19, "9, 13 or 19")
  filepath := flag.String("load", "", "load figure")
  ai := flag.Bool("ai", false, "use AI")
  rpc := flag.Bool("rpc", false, "enable RPC server")
  flag.Parse()

  b := board.New(*size)
  if len(*filepath) != 0 {
    fmt.Printf("%s loading...\n", *filepath)
    b.Load(*filepath)
  }

  if *server {
    startHTTPServer(*port, b, *ai)
  } else if *rpc {
    startRPCServer(*port, b, *ai)
  } else {
    startConsole(b, *ai)
  }
}
