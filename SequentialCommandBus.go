package eventago

type SequentialCommandBus struct {
	locator      string
	commandStack []string
	executing    bool
}

func (s *SequentialCommandBus) init(c CommandHandlerLocator) {
	s.locator = c.locator
}

func (s *SequentialCommandBus) handle(c Command) bool {
	s.commandStack = append(c.command)

	if s.executing == true {
		return false
	}
	return true
}
