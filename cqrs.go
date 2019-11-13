// Package cqrs provides a CQRS reference implementation.
//
package cqrs

import (
	uuid "dummy_bank/dummycqrs/Library/internal/uuid"
	"reflect"
)

// used for the type name more easily if desired.
func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

// NewUUID returns a new v4 uuid as a string
func NewUUID() string {
	return uuid.NewUUID()
}

// Int returns a pointer to int.
func Int(i int) *int {
	return &i
}
