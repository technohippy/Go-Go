package console_player

import (
  "os"
  "fmt"
  "bufio"
  "regexp"
  "strconv"
  "./cell"
  "./board"
  "./player"
  "./history"
)

type ConsolePlayer struct {
  name string
  teban player.Teban
}

func New(n string, t player.Teban) *ConsolePlayer {
  return &ConsolePlayer{n, t}
}

func (cp *ConsolePlayer)Name() string {
  return cp.name
}

func (cp *ConsolePlayer)Teban() player.Teban {
  return cp.teban
}

func (cp *ConsolePlayer)Next(b *board.Board, h *history.History, agehama [2]int) *player.Response {
  res := new(player.Response)

  r := bufio.NewReader(os.Stdin)
  numberRe := regexp.MustCompile("[0-9]+")
  var color cell.Cell
  if cp.teban == player.SENTE {
    color = cell.BLACK
  } else {
    color = cell.WHITE
  }

  for {
    fmt.Printf("[%s] x, y > ", cp.name)
    in, _ := r.ReadString('\n')
    if in == "pass\n" {
      b.Pass(color)
      fmt.Printf("[%s] Pass!\n", cp.name)
      res.Status = player.PASS
      return res
    } else if in == "giveup\n" {
      res.Status = player.GIVEUP
      return res
    }

    xy := numberRe.FindAllString(in, 2)
    if len(xy) < 2 {
      b.Pass(color)
      fmt.Printf("[%s] Input again!\n", cp.name)
    } else {
      x, _ := strconv.Atoi(xy[0])
      y, _ := strconv.Atoi(xy[1])
      takenOffs, ok := b.PutAt(color, x, y)
      if ok == board.OK {
        fmt.Printf("[%s] Put (%d, %d) and Take %d\n", cp.name, x, y, len(takenOffs))
        res.Status = player.PUT
        res.Data = map[string]interface{}{"x":x, "y":y, "taken":takenOffs}
        break
      } else if ok == board.KO {
        fmt.Printf("[%s] Ko\n", cp.name)
        res.Status = player.KO
      } else if ok == board.FORBIDDEN {
        fmt.Printf("[%s] Forbidden\n", cp.name)
        res.Status = player.FORBIDDEN
      }
    }
  }
  return res
}
