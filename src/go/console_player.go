// This package implements a user interface for console.
package console_player

import (
  "os"
  "fmt"
  "bufio"
  "regexp"
  "strconv"
  "./board"
  "./match"
)

// Console UI.
type ConsolePlayer struct {
  name string
  teban match.Teban
}

// New returns a user interface for console.
func New(n string, t match.Teban) *ConsolePlayer {
  return &ConsolePlayer{n, t}
}

// Name returns a name of a player.
func (cp *ConsolePlayer)Name() string {
  return cp.name
}

// Teban returns a teban of a player.
func (cp *ConsolePlayer)Teban() match.Teban {
  return cp.teban
}

// Next will decide the next play.
func (cp *ConsolePlayer)Next(m *match.Match) *match.Response {
  r := bufio.NewReader(os.Stdin)
  numberRe := regexp.MustCompile("[0-9]+")
  color := cp.teban.Color()

  for {
    fmt.Printf("[%s] x, y > ", cp.name)
    in, _ := r.ReadString('\n')
    if in == "pass\n" {
      fmt.Printf("[%s] Pass!\n", cp.name)
      return match.NewPassResponse()
    } else if in == "giveup\n" {
      return match.NewGiveupResponse()
    } else if in == "help\n" {
      fmt.Println("Commands: (x,y), pass, giveup, help")
      continue
    }

    xy := numberRe.FindAllString(in, 2)
    if len(xy) < 2 {
      fmt.Printf("[%s] Input again!\n", cp.name)
    } else {
      x, _ := strconv.Atoi(xy[0])
      y, _ := strconv.Atoi(xy[1])
      takenOffs, ok := m.Board.PutAt(color, x, y, m.History)
      if ok == board.OK {
        fmt.Printf("[%s] Put (%d, %d) and Take %d\n", cp.name, x, y, len(takenOffs))
        return match.NewPutResponse(x, y, takenOffs)
      } else if ok == board.KO {
        fmt.Printf("[%s] Ko\n", cp.name)
        return match.NewKoResponse()
      } else if ok == board.FORBIDDEN {
        fmt.Printf("[%s] Forbidden\n", cp.name)
        return match.NewForbiddenResponse()
      }
    }
  }
  return nil // never reached
}
