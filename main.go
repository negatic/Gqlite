package main

import (
	"bufio"
	"net"

	"github.com/negatic/gqlite/repl"
)

func handleConn(c net.Conn) {

	c.Write([]byte("gqlite>"))
	data, _ := bufio.NewReader(c).ReadString('\n')
	r, _ := repl.ExecuteCommand(data)
	c.Write([]byte(r))
	c.Write([]byte("\n"))
}

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:8080")
	conn, _ := l.Accept()
	l.Close()

	for {
		handleConn(conn)
	}
}
