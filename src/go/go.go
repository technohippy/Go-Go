package main

import (
  "os"
  "fmt"
  "flag"
  "./board"
  "./player"
  "./server"
  "./console_player"
  //"./auto_player"
)

func startConsole(b *board.Board, ps [2]player.Player) {
  fmt.Printf("%v", b)
  turn := 0
  pass := false
  for {
    switch ps[turn].Next(b) {
      case player.PUT:
        fmt.Printf("%v", b)
        pass = false
      case player.PASS:
        if pass {
          fmt.Println("Finish!")
          os.Exit(0)
        } else {
          pass = true
        }
      case player.GIVEUP:
        fmt.Printf("%s Win!\n", ps[turn].Name())
        os.Exit(0)
    }
    turn++
    turn %= 2
  }
}

func startServer(port int, b *board.Board, ps [2]player.Player) {
  s := server.New(b, ps)
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
    players := [2]player.Player{
      //http_player.New("ando", player.SENTE),
      //http_player.New("yasushi", player.GOTE)}
      console_player.New("ando", player.SENTE),
      console_player.New("yasushi", player.GOTE)}
    startServer(*port, b, players)
  } else {
    players := [2]player.Player{
      console_player.New("ando", player.SENTE),
      console_player.New("yasushi", player.GOTE)}
    startConsole(b, players)
  }
}
