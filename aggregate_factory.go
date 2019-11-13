package cqrs

import (
	"fmt"
)

// AggregateFactory returns aggregate instances of a specified type with the
// AggregateID set to the uuid provided.
//
// An aggregate factory is typically a dependency of the repository that will
// delegate instantiation of aggregate instances to the Aggregate factory.
type AggregateFactory interface {
	GetAggregate(string, string) AggregateMain
}

// DelegateAggregateFactory is an implementation of the AggregateFactory interface
// that supports registration of delegate functions to perform aggregate instantiation.
type DelegateAggregateFactory struct {
	delegates map[string]func(string) AggregateMain
}

// NewDelegateAggregateFactory contructs a new DelegateAggregateFactory
func NewDelegateAggregateFactory() *DelegateAggregateFactory {
	return &DelegateAggregateFactory{
		delegates: make(map[string]func(string) AggregateMain),
	}
}

// RegisterDelegate is used to register a new funtion for instantiation of an
// aggregate instance.
//
// 	func(id string) AggregateMain {return NewMyAggregateType(id)}
// 	func(id string) AggregateMain { return &MyAggregateType{AggregateBase:NewAggregateBase(id)} }
func (t *DelegateAggregateFactory) RegisterDelegate(aggregate AggregateMain, delegate func(string) AggregateMain) error {
	typeName := typeOf(aggregate)
	if _, ok := t.delegates[typeName]; ok {
		return fmt.Errorf("Factory delegate already registered for type: \"%s\"", typeName)
	}
	t.delegates[typeName] = delegate
	return nil
}

// GetAggregate calls the delegate for the type specified and returns the result.
func (t *DelegateAggregateFactory) GetAggregate(typeName string, id string) AggregateMain {
	if f, ok := t.delegates[typeName]; ok {
		return f(id)
	}
	return nil
}
