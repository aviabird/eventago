package eventago

// CommandHandlerLocator is for locator
type CommandHandlerLocator struct {
	locator string
}

// CommandHandlerLocatorroot CommandHandlerLocatorroot
type CommandHandlerLocatorroot interface {
	getCommandHandler(command CommandRoot)
}
