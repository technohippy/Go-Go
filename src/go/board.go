package board

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
  "container/vector"
  "./point"
  "./cell"
  "./history"
)

type PutStatus int
const (
  OK PutStatus = 0
  OCCUPIED PutStatus = 1
  KOU PutStatus = 2
  TAKEN PutStatus = 3
)

type Board struct {
  board [][]cell.Cell
  size int
  history *history.History
}

func New(size int) *Board {
  b := new(Board)
  b.history = history.New()
  b.size = size
  b.board = make([][]cell.Cell, size)
  for y := 0; y < size; y++ {
    b.board[y] = make([]cell.Cell, size)
    for x := 0; x < size; x++ {
      b.board[y][x] = cell.SPACE
    }
  }
  return b
}
func New19() *Board { return New(19) }
func New13() *Board { return New(13) }
func New9() *Board { return New(9) }

func (b *Board)At(x int, y int) cell.Cell {
  if x < 1 || b.size < x || y < 1 || b.size < y { return cell.OB }
  return b.board[y-1][x-1]
}

func (b *Board)charAt(x int, y int) byte {
  return map[cell.Cell]byte{cell.SPACE:'+', cell.BLACK:'@', cell.WHITE:'O'}[b.At(x, y)]
}

func (b *Board)PutAt(c cell.Cell, x int, y int) (vector.Vector, PutStatus) {
  if b.At(x, y) != cell.SPACE {
    return nil, OCCUPIED
  }

  b.putAt(c, x, y)
  takenOffs := vector.Vector{}
  rc := c.Reverse()
  if b.shouldBeTakenOff(x-1, y, rc) { b.takeOff(x-1, y, rc, &takenOffs) }
  if b.shouldBeTakenOff(x+1, y, rc) { b.takeOff(x+1, y, rc, &takenOffs) }
  if b.shouldBeTakenOff(x, y-1, rc) { b.takeOff(x, y-1, rc, &takenOffs) }
  if b.shouldBeTakenOff(x, y+1, rc) { b.takeOff(x, y+1, rc, &takenOffs) }

  if b.shouldBeTakenOff(x, y, c) {
    b.putAt(cell.SPACE, x, y)
    return nil, TAKEN
  }

  if len(takenOffs) == 1 && b.history.IsKou(c, x, y) {
    b.putAt(cell.SPACE, x, y)
    taken := takenOffs.Last().(point.Point)
    b.putAt(c.Reverse(), taken.X(), taken.Y())
    return nil, KOU
  }

  b.history.Add(c, x, y, takenOffs)

  return takenOffs, OK
}

func (b *Board)TakeAt(x int, y int) cell.Cell {
  c := b.At(x, y)
  b.putAt(cell.SPACE, x, y)
  return c
}

func (b *Board)putAt(c cell.Cell, x int, y int) {
  b.board[y-1][x-1] = c
}

func (b *Board)createCheckTable() [][]int {
  checked := make([][]int, b.size+2)
  for i := 0; i <= b.size+1; i++ {
    checked[i] = make([]int, b.size+2)
    for j := 0; j <= b.size+1; j++ {
      checked[i][j] = 0
    }
  }
  return checked
}

func (b *Board)shouldBeTakenOff(x int, y int, c cell.Cell) bool {
  if b.At(x, y) != c { return false }
  return b.isTangentToSpace(x, y, c, b.createCheckTable())
}

func (b *Board)isTangentToSpace(x int, y int, c cell.Cell, checked [][]int) bool {
  if checked[y][x-1] == 0 {
    checked[y][x-1] = 1
    switch b.At(x-1, y) {
      case cell.SPACE:
        return false
      case c:
        if !b.isTangentToSpace(x-1, y, c, checked) { return false }
    }
  }

  if checked[y][x+1] == 0 {
    checked[y][x+1] = 1
    switch b.At(x+1, y) {
      case cell.SPACE:
        return false
      case c:
        if !b.isTangentToSpace(x+1, y, c, checked) { return false }
    }
  }

  if checked[y-1][x] == 0 {
    checked[y-1][x] = 1
    switch b.At(x, y-1) {
      case cell.SPACE:
        return false
      case c:
        if !b.isTangentToSpace(x, y-1, c, checked) { return false }
    }
  }

  if checked[y+1][x] == 0 {
    checked[y+1][x] = 1
    switch b.At(x, y+1) {
      case cell.SPACE:
        return false
      case c:
        if !b.isTangentToSpace(x, y+1, c, checked) { return false }
    }
  }

  return true
}

func (b *Board)takeOff(x int, y int, c cell.Cell, takenOffs *vector.Vector) {
  if b.At(x, y) == c {
    b.TakeAt(x, y)
    takenOffs.Push(point.Point{x, y})
    b.takeOff(x-1, y, c, takenOffs)
    b.takeOff(x+1, y, c, takenOffs)
    b.takeOff(x, y-1, c, takenOffs)
    b.takeOff(x, y+1, c, takenOffs)
  }
}

func (b *Board)Load(filepath string) {
  file, err := os.Open(filepath, os.O_RDONLY, 0666)
  if err != nil {
    panic(err)
  }

  y := 1
  re := regexp.MustCompile("[^O@+\n]")
  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      if err == os.EOF { break } else { panic(err) }
    }

    line = re.ReplaceAllString(line, "")
    if len(line) == 0 || len(line) == 1 {
      continue
    }
    x := 1
    for i := 0; i < len(line); i++ {
      switch line[i] {
      case '+': b.putAt(cell.SPACE, x, y)
      case '@': b.putAt(cell.BLACK, x, y)
      case 'O': b.putAt(cell.WHITE, x, y)
      }
      x++
      if b.size < x { break }
    }
    y++
    if b.size < y { break }
  }
}

func (b *Board)String() string {
  ret := "     ["
  for i := 1; i <= b.size; i++ {
    ret = fmt.Sprintf("%s%2d", ret, i)
  }
  ret = fmt.Sprintf("%s]\n", ret)
  for y := 0; y < b.size; y++ {
    ret = fmt.Sprintf("%s[%2d]  ", ret, y+1)
    for x := 0; x < b.size; x++ {
      ret = fmt.Sprintf("%s %c", ret, b.charAt(x+1, y+1))
    }
    ret = fmt.Sprintf("%s\n", ret)
  }
  return ret
}
