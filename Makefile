include $(GOROOT)/src/Make.inc

TARG    = gomud
GOFILES = \
	gomud.go\
	server.go\
	vt100.go\
	entity.go\

include $(GOROOT)/src/Make.cmd
