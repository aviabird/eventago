package eventago

// SequentialCommandBus SequentialCommandBus
type SequentialCommandBus struct {
	locator      string
	commandStack []string
	executing    bool
}

func (s *SequentialCommandBus) init(c CommandHandlerLocator) {
	s.locator = c.locator
}

func (s *SequentialCommandBus) handle(cmd Command) bool {
	s.commandStack = append(s.commandStack, cmd.command)

	if s.executing == true {
		return false
	}
	return true
}
