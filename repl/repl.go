package repl

import (
	"net"
)

func ParseCommands(c string) (string, error) {
	if c == "salam\n" {
		return help()
	} else if c == ".exit\n" {
		exit()
	}
	return "Command Not Found", nil
}

func help() (string, error) {
	return "help", nil
}

func exit() (c net.Conn) {
	c.Close()
	return
}
