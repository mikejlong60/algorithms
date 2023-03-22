package chapter5

import "fmt"

type Stack[T any] struct {
	head T
	tail *Stack[T]
}

func (w *Stack[T]) String() string {
	return fmt.Sprintf("Stack {head: %v, tail: %v}", w.head, w.tail)
}

// Pushes the T to the top of the stack and returns a new stack with that T on the top.
// The l stack is not modified in any way
// A nil stack l is OK
func Push[T any](h T, l *Stack[T]) *Stack[T] {
	if l == nil {
		return &Stack[T]{
			head: h,
			tail: nil,
		}
	} else {
		return &Stack[T]{
			head: h,
			tail: l,
		}
	}
}

// Returns the T at the top of the stack and the new stack without that T on the top.
// The l stack is not modified in any way
// If the stack is empty it returns an error and the zero value of T and the stack
func Pop[T any](l *Stack[T]) (error, T, *Stack[T]) {
	if l == nil {
		var a T
		return fmt.Errorf("Stack is empty"), a, Zero[T]()
	} else {
		return nil, l.head, l.tail
	}
}

// Returns the top value of type T without modifying the stack
// If the stack is empty it returns an error and the zero value of T
func Peek[T any](l *Stack[T]) (error, T) {
	if l == nil {
		var a T
		return fmt.Errorf("Stack is empty"), a
	} else {
		return nil, l.head
	}
}

func Zero[T any]() *Stack[T] {
	return nil
}

// Converts an array into a stack efficiently and preserves the order, first array element of source array is on top.
func FromArray[T any](xs []T) *Stack[T] {
	var r = Zero[T]()
	l := len(xs)
	for i := l - 1; i >= 0; i-- {
		r = Push(xs[i], r)
	}
	return r
}
