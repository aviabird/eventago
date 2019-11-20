package eventago

// CommandBus interface for handle related
type CommandBus interface {
	handle(command Command)
}
