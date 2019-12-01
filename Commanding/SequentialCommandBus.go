package commanding

import eventago "github.com/aviabird/eventago"

// SequentialCommandBus SequentialCommandBus
type SequentialCommandBus struct {
	locator      string
	commandStack []string
	executing    bool
}

func (s *SequentialCommandBus) init(c CommandHandlerLocator) {
	s.locator = c.locator
}

func (s *SequentialCommandBus) handle(cmd eventago.Command) bool {
	// s.commandStack = append(s.commandStack, cmd.Co)

	// if s.executing == true {
	// 	return false
	// }
	return true
}
