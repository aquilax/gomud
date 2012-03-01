package main

const (
	NEWCONNECTION = iota
	NEWUSER
	ENTERNEWPASS
	ENTERPASS
)

type LogonState int

type Logon struct {
	Handler
	m_state LogonState
	m_errors int
	m_name string
	m_pass string
	p_connection *Connection
}

func (l *Logon) Hungup() {
	//log
}

func (l *Logon) Flooded() {
	//log
}

func (l *Logon) Enter() {
	//log
	l.p_connection.SendString(red+bold+"Welcome"+newline+"Please enter your name or \"new\" if you are new: "+reset)
}
