// This package implements a web server for Go games.
// Note that games are volatile.
package http_server

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
  "./match"
  "./http_player"
)

// Directory path for template files.
const (
  template_dir = "templates/"
)

// Web server for a Go game.
type HTTPServer struct {
  matches *vector.Vector
}

// New returns a server for Go games.
func New(b *board.Board, ps [2]match.Player) *HTTPServer {
  ms := new(vector.Vector)
  ms.Push(match.New(b, ps))
  return &HTTPServer{ms}
}

// currentMatch returns the recent match.
func (s *HTTPServer)currentMatch() *match.Match {
  return s.matches.Last().(*match.Match)
}

// Cache for templates.
var templates = make(map[string]*template.Template)

// getTemplate returns a template at a given filepath.
func getTemplate(filepath string) *template.Template {
  if templates[filepath] == nil {
    file, err := os.Open(template_dir+filepath, os.O_RDONLY, 0666)
    if err != nil { log.Exit("not exist: %s", filepath) }
    reader := bufio.NewReader(file)
    body, _ := reader.ReadString('~')
    tmpl := template.MustParse(body, nil)
    templates[filepath] = tmpl
  }
  return templates[filepath]
}

// Start will start service.
func (s *HTTPServer)Start(port int) {
  http.HandleFunc("/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    log.Println("access /")
    params := new(struct { });
    tmpl := getTemplate("index.html")
    err := tmpl.Execute(params, rw)
    if err != nil { log.Exit("%v", err) }
  }))

  http.HandleFunc("/put", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    log.Println("access /put")
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
        fmt.Fprint(rw, "({'message':'コウです'})")
      case match.FORBIDDEN:
        fmt.Fprint(rw, "({'message':'着手禁止です'})")
      case match.PASS:
        if matchStatus == match.FINISH {
          fmt.Fprint(rw, "({'message':'終了しました'})")
        } else {
          fmt.Fprint(rw, "({'message':'パスしました'})")
        }
      case match.GIVEUP:
        fmt.Fprint(rw, "({'message':'終了しました'})")
    }
  }))

  http.HandleFunc("/get", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    //log.Println("access /get")
    m := s.currentMatch()
    if _, ok := m.CurrentPlayer().(*http_player.HttpPlayer); !ok {
      m.Next()
    }
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
