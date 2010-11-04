// Ref. http://en.wikipedia.org/wiki/Go_(game)
package main

import (
  "os"
  "fmt"
  "flag"
  "./match"
  "./board"
  "./server"
  "./console_player"
  "./http_player"
  "./auto_player"
)

// startConsole will start a console for a Go game.
func startConsole(b *board.Board) {
  players := [2]match.Player{
    console_player.New("ando", match.SENTE),
    //console_player.New("yasushi", match.GOTE)}
    auto_player.New(match.GOTE)}
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
          os.Exit(0)
        }
      case match.GIVEUP:
        fmt.Printf("%s Win!\n", (*m.Winner).Name())
        os.Exit(0)
    }
  }
}

// startConsole will start a server for a Go game.
func startServer(port int, b *board.Board) {
  players := [2]match.Player{
    http_player.New("ando", match.SENTE),
    http_player.New("yasushi", match.GOTE)}
  s := server.New(b, players)
  s.Start(port)
}

// main
func main() {
  server := flag.Bool("server", false, "run server or not")
  port := flag.Int("port", 55555, "port number")
  size := flag.Int("size", 19, "9, 13 or 19")
  filepath := flag.String("load", "", "load figure")
  flag.Parse()

  b := board.New(*size)
  if len(*filepath) != 0 {
    fmt.Printf("%s loading...\n", *filepath)
    b.Load(*filepath)
  }

  if *server {
    startServer(*port, b)
  } else {
    startConsole(b)
  }
}
