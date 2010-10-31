package main

import (
  "os"
  "fmt"
  "flag"
  "./board"
  "./player"
  "./console_player"
  //"./auto_player"
)

func start(b *board.Board, ps [2]player.Player) {
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

func main() {
  size := flag.Int("size", 19, "9, 13 or 19")
  filepath := flag.String("load", "", "load figure")
  flag.Parse()

  b := board.New(*size)
  if len(*filepath) != 0 {
    fmt.Printf("%s loading...\n", *filepath)
    b.Load(*filepath)
  }
  fmt.Printf("%v", b)

  players := [2]player.Player{
    console_player.NewConsolePlayer("ando", player.SENTE),
    console_player.NewConsolePlayer("yasushi", player.GOTE)}

  start(b, players)
}
