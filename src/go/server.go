// http://d.hatena.ne.jp/tokuhirom/20091117/1258418742
// http://d.hatena.ne.jp/sfujiwara/20091116/1258388969
// http://d.hatena.ne.jp/sfujiwara/20091117/1258438816
package server

import (
  "os"
  "fmt"
  "http"
  "bufio"
  "log"
  "strconv"
  "template"
  "./board"
  "./player"
  "./cell"
)

const (
  template_dir = "templates"
)

type Server struct {
  board *board.Board
  players [2]player.Player
  turn int
  pass bool
}

func New(b *board.Board, ps [2]player.Player) *Server {
  return &Server{b, ps, 0, false}
}

func getTemplate(filepath string) *template.Template {
  file, err := os.Open(fmt.Sprintf("%s/%s", template_dir, filepath), os.O_RDONLY, 0666)
  if err != nil { log.Exit("not exist: %s", filepath) }
  reader := bufio.NewReader(file)
  body, _ := reader.ReadString('~')
  tmpl := template.MustParse(body, nil)
  return tmpl
}

func (s *Server)Start(port int) {
  http.HandleFunc("/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    params := new(struct { board *board.Board });
    params.board = s.board
    tmpl := getTemplate("index.html")
    err := tmpl.Execute(params, rw)
    if err != nil { log.Exit("%v", err) }
  }))

  http.HandleFunc("/put", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    x, _ := strconv.Atoi(req.Form["x"][0])
    y, _ := strconv.Atoi(req.Form["y"][0])

    if s.players[s.turn].Teban() == player.SENTE {
      s.board.PutAt(cell.BLACK, x, y)
    } else {
      s.board.PutAt(cell.WHITE, x, y)
    }
    log.Printf("\n%v\n", s.board)
    s.turn++
    s.turn %= 2

    fmt.Fprint(rw, s.board.Json())
  }))

  http.Handle("/css/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    http.ServeFile(rw, req, "./htdocs/" + req.URL.Path);
  }));

  http.Handle("/js/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    http.ServeFile(rw, req, "./htdocs/" + req.URL.Path);
  }));

  err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
  if err != nil {
    log.Exit("ListenAndServe: ", err.String())
  }
}
