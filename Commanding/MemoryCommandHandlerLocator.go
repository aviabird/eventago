package commanding

import (
	"fmt"
	"reflect"

	eventago "github.com/aviabird/eventago"
)

// MemoryCommandHandlerLocator MemoryCommandHandlerLocator
type MemoryCommandHandlerLocator struct {
	handlers []string
}

// MemoryCommandHandlerLocatorRoot MemoryCommandHandlerLocatorRoot
type MemoryCommandHandlerLocatorRoot interface {
	getCommandHandler(command eventago.Command)
	register(commandType string, service []string)
}

func getCommandHandler(command eventago.Command) eventago.Command {
	var checkEmpty eventago.Command
	commandType := command
	if commandType == checkEmpty {
		fmt.Printf("No service registered for command type %s", commandType)
	}
	return commandType
}

func register(commandType string, service []string) {

	if reflect.TypeOf(service) != nil {
		fmt.Printf("No service registered for command type %s", commandType)
	}
}

// func (mel MemoryEventHandlerLocator) getHandlersFor(eventname eventago.EventNameRoot) []string {
// 	for _, v := range mel.handlers {
// 		if v == eventname.aggregateID {
// 			return nil
// 		}
// 	}
// 	mel.handlers = append(mel.handlers, eventname.aggregateID)
// 	return mel.handlers
// }
