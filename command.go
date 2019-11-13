package cqrs

// CommandMessage is the interface that a command message must implement.
type CommandMessage interface {

	// AggregateID returns the ID of the Aggregate that the command relates to
	AggregateID() string

	// Command returns the actual command which is the payload of the command message.
	Command() interface{}

	// CommandType returns a string descriptor of the command name
	CommandType() string
}

// CommandInformation is an implementation of the command message interface.
type CommandInformation struct {
	id      string
	command interface{}
	headers map[string]interface{}
}

// NewCommandMessage returns a new command descriptor
func NewCommandMessage(aggregateID string, command interface{}) *CommandInformation {
	return &CommandInformation{
		id:      aggregateID,
		command: command,
		headers: make(map[string]interface{}),
	}
}

// CommandType returns the command type name as a string
func (c *CommandInformation) CommandType() string {
	return typeOf(c.command)
}

// AggregateID returns the ID of the aggregate that the command relates to.
func (c *CommandInformation) AggregateID() string {
	return c.id
}

// Command returns the actual command payload of the message.
func (c *CommandInformation) Command() interface{} {
	return c.command
}
