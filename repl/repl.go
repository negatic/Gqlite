package repl

func ParseCommands(c string) (string, error) {
	if c == "salam\n" {
		return help()
	}
	return "Command Not Found", nil
}

func help() (string, error) {
	return "help", nil
}
