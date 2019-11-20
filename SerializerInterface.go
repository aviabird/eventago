package eventago

// SerializerInterface SerializerInterface
type SerializerInterface interface {
	serialize(event DomainEvent, format string)
	deserialize(eventClass string, data []string, format string)
}
