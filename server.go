package main

import (
	"os"
	"net"
	"log"
	"fmt"
)

type Request struct {
	conn net.Conn;
}

type Server struct {
	port int
	log *log.Logger
	connections int
	totalConnections int
}

func (server *Server) Handle (req *Request) {
	req.conn.Write([]byte(red+"great now bye\n"))
	//req.conn.Close()
	//server.connections--
}

func (server *Server) Logger() *log.Logger {
	return server.log
}

func (server *Server) Serve() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))
	if err != nil {
		server.log.Printf("cannot start server: %s\n", err)
		os.Exit(1)
	}
	server.log.Printf("waiting for connections on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			server.log.Printf("could not accept: %s\n", err)
		} else {
			server.log.Printf("connected: %s\n", conn.RemoteAddr())
			server.connections++
			server.totalConnections++
			go server.Handle(&Request{conn})
		} 
	}
}
