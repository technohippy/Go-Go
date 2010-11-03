package main

import (
  "os"
  "fmt"
  "flag"
  "./match"
  "./board"
  "./player"
  "./server"
  "./console_player"
  //"./auto_player"
)

func startConsole(b *board.Board) {
  players := [2]player.Player{
    console_player.New("ando", player.SENTE),
    console_player.New("yasushi", player.GOTE)}
  m := match.New(b, players)

  fmt.Printf("%v", b)
  for {
    matchStatus, playerStatus := m.Next()
    switch playerStatus {
      case player.PUT:
        fmt.Printf("%v", b)
      case player.KO:
        fmt.Printf("Ko", b)
      case player.FORBIDDEN:
        fmt.Printf("Forbidden", b)
      case player.PASS:
        if matchStatus == match.FINISH {
          fmt.Println("Finish!")
          os.Exit(0)
        }
      case player.GIVEUP:
        fmt.Printf("%s Win!\n", (*m.Winner).Name())
        os.Exit(0)
    }
  }
}

func startServer(port int, b *board.Board) {
  players := [2]player.Player{
    //http_player.New("ando", player.SENTE),
    //http_player.New("yasushi", player.GOTE)}
    console_player.New("ando", player.SENTE),
    console_player.New("yasushi", player.GOTE)}
  s := server.New(b, players)
  s.Start(port)
}

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
