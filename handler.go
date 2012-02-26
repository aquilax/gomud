package main

type Handler interface {
	Enter()
	Leave()
	Hungup()
	Flooded()
	Handle(command string)
}
