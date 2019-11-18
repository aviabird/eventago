package eventago

import (
	"fmt"
	"reflect"
)

type MemoryCommandHandlerLocator interface {
	getCommandHandler(command Command)
	register(commandType string, service []string)
}

func getCommandHandler(command Command) Command {
	commandType := command
	if commandType == nil {
		fmt.Sprintf("No service registered for command type %s", commandType)
	}
	return commandType
}

func register(commandType string, service []string) {
	if (reflect.TypeOf(service) != []string) {
		fmt.Sprintf("No service registered for command type %s", commandType)
	}
}
