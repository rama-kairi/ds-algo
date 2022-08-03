package stack

import "fmt"

// Stack is a linear data structure which follows a particular order in which the operations are performed. The order may be LIFO(Last In First Out) or FILO(First In Last Out). The order is determined by the implementation of the data structure.

// This strategy states that the element that is inserted last will come out first. You can take a pile of plates kept on top of each other as a real-life example. The plate which we put last is on the top and since we remove the plate that is at the top, we can say that the plate that was put last comes out first.

// There are many real-life examples of a stack. Consider an example of plates stacked over one another in the canteen. The plate which is at the top is the first one to be removed, i.e. the plate which has been placed at the bottommost position remains in the stack for the longest period of time. So, it can be simply seen to follow LIFO(Last In First Out)/FILO(First In Last Out) order.

// Basic operation of Stack

// These are the operations we are going to implement
// Push: Adds element to the top of the stack and return “Stack Overflow” if the stack is full.
// Pop: Removes the top most element from the stack and return “Stack Underflow” if the stack is empty.
// isEmpty: Check if the stack is empty.
// isFull: Check if the stack is full
// Peek: Return the top most element from the stack.
// String: Display the items of stack

// Basic usage of stack:
// Converting infix to postfix expressions.
// Undo/Redo button/operation in word processors.
// Syntaxes in languages are parsed using stacks.
// It is used in many virtual machines like JVM.
// Forward-backward surfing in the browser.
// History of visited websites.
// Message logs and all messages you get are arranged in a stack.
// Call logs, E-mails, Google photos’ any gallery, YouTube downloads, Notifications ( latest appears first ).
// Scratch card’s earned after Google pay transaction.
// Wearing/Removing Bangles, Pile of Dinner Plates, Stacked chairs.
// Changing wearables on a cold evening, first in, comes out at last.
// Last Hired, First Fired - which is typically utilized when a company reduces its workforce in an economic recession.
// Loading bullets into the magazine of a gun. The last one to go in is fired first. Bam!
// Java Virtual Machine.
// Recursion.
// Used in IDEs to check for proper parentheses matching
// Media playlist. T o play previous and next song

// Implementing a stack using Array/Slice

// Pros: Easy to implement. Memory is saved as pointers are not involved.
// Cons: It is not dynamic. It doesn’t grow and shrink depending on needs at runtime.

// Pros and Cons of Stack Data Structure
// Advantages of Stack:

// - The linked list implementation of a stack can grow and shrink according to the needs at runtime.
// - It is used in many virtual machines like JVM.
// - Stacks are more secure and reliable as they do not get corrupted easily.
// - Stack cleans up the objects automatically.

// Disadvantages of Stack:

// - Requires extra memory due to involvement of pointers.
// - Random accessing is not possible in stack.
// - The total of size of the stack must be defined before.
// - If the stack falls outside the memory it can lead to abnormal termination.
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
