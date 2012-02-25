include $(GOROOT)/src/Make.inc

TARG    = gomud
GOFILES = \
	gomud.go\
	vt100.go\
	server.go\

include $(GOROOT)/src/Make.cmd
