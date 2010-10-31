include $(GOROOT)/src/Make.inc

TARG=go-go
GOFILES=src/go/go.go
O_FILES=board.$O player.$O console_player.$O auto_player.$O go.$O

all: $(O_FILES)
install: $(O_FILES)

include $(GOROOT)/src/Make.cmd

board.$O:
	$(QUOTED_GOBIN)/$(GC) -o board.$O src/go/board.go

player.$O:
	$(QUOTED_GOBIN)/$(GC) -o player.$O src/go/player.go

console_player.$O:
	$(QUOTED_GOBIN)/$(GC) -o console_player.$O src/go/console_player.go

auto_player.$O:
	$(QUOTED_GOBIN)/$(GC) -o auto_player.$O src/go/auto_player.go

go.$O:
	$(QUOTED_GOBIN)/$(GC) -o go.$O src/go/go.go
