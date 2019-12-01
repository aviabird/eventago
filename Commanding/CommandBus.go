package commanding

import eventago "github.com/aviabird/eventago"

// CommandBus interface for handle related
type CommandBus interface {
	handle(command eventago.Command)
}
