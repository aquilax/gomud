package main

type Game struct {
	server         *Server
	itemDatabase   *ItemDatabase
	playerDatabase *PlayerDatabase
	running        bool
}

func NewGame(server *Server) *Game {
	return &Game{
		server:         server,
		running:        false,
		itemDatabase:   NewItemDatabase(),
		playerDatabase: NewPlayerDatabase(),
	}
}

func (g *Game) Start() error {
	err := g.server.Serve()
	if err != nil {
		return err
	}
	g.running = true
}
