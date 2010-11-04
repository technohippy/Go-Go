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
  "container/vector"
  "./board"
  //"./cell"
  "./match"
  "./http_player"
)

const (
  template_dir = "templates"
)

type Server struct {
  matches *vector.Vector
  board *board.Board
  players [2]match.Player
  turn int
  pass bool
}

func New(b *board.Board, ps [2]match.Player) *Server {
  ms := new(vector.Vector)
  ms.Push(match.New(b, ps))
  return &Server{ms, b, ps, 0, false}
}

func (s *Server)currentMatch() *match.Match {
  return s.matches.Last().(*match.Match)
}

var templates = make(map[string]*template.Template)
func getTemplate(filepath string) *template.Template {
  if templates[filepath] == nil {
    file, err := os.Open(fmt.Sprintf("%s/%s", template_dir, filepath), os.O_RDONLY, 0666)
    if err != nil { log.Exit("not exist: %s", filepath) }
    reader := bufio.NewReader(file)
    body, _ := reader.ReadString('~')
    tmpl := template.MustParse(body, nil)
    templates[filepath] = tmpl
  }
  return templates[filepath]
}

func (s *Server)Start(port int) {
  http.HandleFunc("/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    params := new(struct { });
    tmpl := getTemplate("index.html")
    err := tmpl.Execute(params, rw)
    if err != nil { log.Exit("%v", err) }
  }))

  http.HandleFunc("/put", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    x, _ := strconv.Atoi(req.Form["x"][0])
    y, _ := strconv.Atoi(req.Form["y"][0])

    m := s.currentMatch()
    if hp, ok := m.CurrentPlayer().(*http_player.HttpPlayer); ok {
      hp.SetNext(x, y)
    }
    matchStatus, playerResponse := m.Next()
    switch playerResponse.Type {
      case match.PUT:
        fmt.Fprint(rw, m.Json())
      case match.KO:
        fmt.Fprint(rw, m.Json())
      case match.FORBIDDEN:
        fmt.Fprint(rw, m.Json())
      case match.PASS:
        if matchStatus == match.FINISH {
          fmt.Fprint(rw, m.Json())
        } else {
          fmt.Fprint(rw, m.Json())
        }
      case match.GIVEUP:
        fmt.Fprint(rw, m.Json())
    }
  }))

  http.HandleFunc("/put", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    m := s.currentMatch()
    fmt.Fprint(rw, m.Json())
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
