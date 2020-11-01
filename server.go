package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Connection struct {
	conn    net.Conn
	buffer  *bufio.ReadWriter
	handler Handler
}

type Server struct {
	port             int
	log              *log.Logger
	connections      int
	totalConnections int
}

func (server *Server) Handle(c *Connection) {
	c.SendString(red + "great now bye\n")
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
			go server.Handle(NewConnection(conn))
		}
	}
}

func NewConnection(connection net.Conn) *Connection {
	return &Connection{
		conn:   connection,
		buffer: bufio.NewReadWriter(bufio.NewReader(connection), bufio.NewWriter(connection)),
	}
}

func (c *Connection) SendString(text string) {
	c.conn.Write([]byte(text))
}

func (c *Connection) BufferData(text string) {
	c.buffer.Write([]byte(text))
}

func (c *Connection) SendBuffer() {
	c.buffer.Flush()
}
