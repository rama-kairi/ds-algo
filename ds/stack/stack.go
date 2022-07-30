package stack

import "fmt"

// A stack is an abstract data type that follows FILO (First-In-Last-Out) principle while inserting and deleting elements. Which means the element which is inserted first will be deleted last and the element inserted last will be deleted first. Stack can be implemented using arrays or linked list. Here we are creating stack using arrays in go. Knowledge of go slices will be an added advantage.

// Basic operation of Stack
// These are the operations we are going to implement

// Push: Adds element to the top of the stack and return “Stack Overflow” if the stack is full.

// Pop: Removes the top most element from the stack and return “Stack Underflow” if the stack is empty.

// isEmpty: Check if the stack is empty.

// isFull: Check if the stack is full

// Peek: Return the top most element from the stack.

// String: Display the items of stack

type stack[T any] struct {
	arr []any
	top int
}

// Push: Adds element to the top of the stack and return “Stack Overflow” if the stack is full.
func (s *stack[T]) Push(item any) {
	if s.top == len(s.arr) {
		fmt.Println("Stack Overflow")
		return
	}
	s.arr[s.top] = item
	s.top++
}

// Pop: Removes the top most element from the stack and return “Stack Underflow” if the stack is empty.
func New[T any]() *stack[T] {
	var s stack[T]
	s.arr = make([]any, 10)
	s.top = 0
	return &s
}

func (s *stack[T]) Pop() {
	if s.top == 0 {
		fmt.Println("Stack Underflow")
		return
	}
	s.top--
	s.arr[s.top] = nil
}

// isEmpty: Check if the stack is empty.
func (s *stack[T]) isEmpty() bool {
	return s.top == 0
}

// isFull: Check if the stack is full
func (s *stack[T]) isFull() bool {
	return s.top == len(s.arr)
}

// Peek: Return the top most element from the stack.
func (s *stack[T]) Peek() any {
	if s.top == 0 {
		fmt.Println("Stack Underflow")
		return nil
	}
	return s.arr[s.top-1]
}

// String: Display the items of stack
func (s *stack[T]) String() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	var str string
	for i := 0; i < s.top; i++ {
		str += fmt.Sprintf("%v ", s.arr[i])
	}
	fmt.Println(str)
}
