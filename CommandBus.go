package eventago

type CommandBus interface {
	handle(command Command)
}
