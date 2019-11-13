// Copyright 2016 Jet Basrawi. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cqrs provides a CQRS reference implementation.
//
package cqrs

import (
	uuid "dummy_bank/dummycqrs/Library/internal/uuid"
	"reflect"
)

// typeOf is a convenience function that returns the name of a type
//
// This is used so commonly throughout the code that it is better to
// have this convenience function and also allows for changing the scheme
// used for the type name more easily if desired.
func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

// NewUUID returns a new v4 uuid as a string
func NewUUID() string {
	return uuid.NewUUID()
}

// Int returns a pointer to int.
//
// There are a number of places where a pointer to int
// is required such as expectedVersion argument on the repository
// and this helper function makes keeps the code cleaner in these
// cases.
func Int(i int) *int {
	return &i
}
