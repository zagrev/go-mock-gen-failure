package main

import "fmt"

// SomeImpl implements the Some interface
type SomeImpl struct {
	value int
}

// Method the implementation of Some.Method()
func (u SomeImpl) Method(argument string) string {
	return fmt.Sprintf("Method called, %v %v\n", argument, u.value)
}
