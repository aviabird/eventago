package eventago

type CommandHandlerLocator interface {
	getCommandHandler(command Command)
}
