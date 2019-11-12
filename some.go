package main

import "fmt"

// SomeImpl implements the Some interface
type SomeImpl struct {
	value int
}

// Method the implementation of Some.Method()
func (u SomeImpl) Method(argument string, arg2 int) string {
	return fmt.Sprintf("Method called, %v %v %v\n", argument, u.value, arg2)
}
