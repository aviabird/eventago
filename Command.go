package eventago

// Command used for locator and commands
type Command struct {
	locator string
	command string
}

// CommandRoot Command
type CommandRoot interface {
}
