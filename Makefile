include $(GOROOT)/src/Make.inc

TARG=go
GOFILES=src/go/go.go
O_FILES=point.$O cell.$O history.$O board.$O match.$O console_player.$O http_player.$O auto_player.$O server.$O go.$O

all: $(O_FILES)
install: $(O_FILES)

include $(GOROOT)/src/Make.cmd

point.$O:
	$(QUOTED_GOBIN)/$(GC) -o point.$O src/go/point.go

cell.$O:
	$(QUOTED_GOBIN)/$(GC) -o cell.$O src/go/cell.go

history.$O:
	$(QUOTED_GOBIN)/$(GC) -o history.$O src/go/history.go

board.$O:
	$(QUOTED_GOBIN)/$(GC) -o board.$O src/go/board.go

player.$O:
	$(QUOTED_GOBIN)/$(GC) -o player.$O src/go/player.go

console_player.$O:
	$(QUOTED_GOBIN)/$(GC) -o console_player.$O src/go/console_player.go

http_player.$O:
	$(QUOTED_GOBIN)/$(GC) -o http_player.$O src/go/http_player.go

auto_player.$O:
	$(QUOTED_GOBIN)/$(GC) -o auto_player.$O src/go/auto_player.go

match.$O:
	$(QUOTED_GOBIN)/$(GC) -o match.$O src/go/match.go

server.$O:
	$(QUOTED_GOBIN)/$(GC) -o server.$O src/go/server.go

go.$O:
	$(QUOTED_GOBIN)/$(GC) -o go.$O src/go/go.go
