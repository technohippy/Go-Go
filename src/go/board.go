package board

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
)

type PutStatus int
const (
  OK PutStatus = 0
  OCCUPIED PutStatus = 1
  KOU PutStatus = 2
)

type Cell int
const (
  SPACE Cell = 0
  BLACK Cell = 1
  WHITE Cell = 2
)

type Board struct {
  board [][]Cell
  size int
}

func New(size int) *Board {
  b := new(Board)
  b.size = size
  b.board = make([][]Cell, size)
  for y := 0; y < size; y++ {
    b.board[y] = make([]Cell, size)
    for x := 0; x < size; x++ {
      b.board[y][x] = SPACE
    }
  }
  return b
}
func New19() *Board { return New(19) }
func New13() *Board { return New(13) }
func New9() *Board { return New(9) }

func (b *Board)At(x int, y int) Cell {
  return b.board[y-1][x-1]
}

func (b *Board)charAt(x int, y int) byte {
  return map[Cell]byte{SPACE:'+', BLACK:'@', WHITE:'O'}[b.At(x, y)]
}

func (b *Board)PutAt(c Cell, x int, y int) (int, PutStatus) {
  if b.board[y-1][x-1] != SPACE {
    return -1, OCCUPIED
  }
  count := 0
  if b.isSurrounded(x-1, y, c) { count += b.takeOff(x-1, y) }
  if b.isSurrounded(x+1, y, c) { count += b.takeOff(x+1, y) }
  if b.isSurrounded(x, y-1, c) { count += b.takeOff(x, y-1) }
  if b.isSurrounded(x, y+1, c) { count += b.takeOff(x, y+1) }
  b.board[y-1][x-1] = c
  return count, OK
}

func (b *Board)isSurrounded(x int, y int, c Cell) bool {
  return false
}

func (b *Board)takeOff(x int, y int) int {
  return 0
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
      case '+': b.PutAt(SPACE, x, y)
      case '@': b.PutAt(BLACK, x, y)
      case 'O': b.PutAt(WHITE, x, y)
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
