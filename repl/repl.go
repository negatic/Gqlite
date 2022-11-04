package repl

import (
	"net"
)

func ExecuteCommand(c string) (string, error) {

	if string(c[0]) == "." {
		if c == ".help" {
			help()
		} else if c == ".exit\n" {
			exit()
		}
	} else {
		parser(c)
	}

	return "Command Not Found", nil
}

func parser(c string) {

}

func help() (string, error) {
	return "help", nil
}

func exit() (c net.Conn) {
	c.Close()
	return
}
