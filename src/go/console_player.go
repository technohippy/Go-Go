package console_player

import (
  "os"
  "fmt"
  "bufio"
  "regexp"
  "strconv"
  "./board"
  "./player"
)

type ConsolePlayer struct {
  name string
  teban player.Teban
}

func NewConsolePlayer(n string, t player.Teban) *ConsolePlayer {
  return &ConsolePlayer{n, t}
}

func (cp *ConsolePlayer)Name() string {
  return cp.name
}

func (cp *ConsolePlayer)Teban() player.Teban {
  return cp.teban
}

func (cp *ConsolePlayer)Next(b *board.Board) player.Status {
  r := bufio.NewReader(os.Stdin)
  re := regexp.MustCompile("[0-9]+")
  var color board.Cell
  if cp.teban == player.SENTE {
    color = board.BLACK
  } else {
    color = board.WHITE
  }

  for {
    fmt.Print("x, y > ")
    in, _ := r.ReadString('\n')
    xy := re.FindAllString(in, 2)
    if len(xy) < 2 {
      fmt.Printf("[%s] Pass!\n", cp.name)
      return player.PASS
    }
    x, _ := strconv.Atoi(xy[0])
    y, _ := strconv.Atoi(xy[1])
    takenOffs, ok := b.PutAt(color, x, y)
    if ok == board.OK {
      fmt.Printf("[%s] Put (%d, %d) and Take %d\n", cp.name, x, y, len(takenOffs))
      break
    }
  }
  return player.PUT
}
