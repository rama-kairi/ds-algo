package main

import (
	"fmt"

	"github.com/rama-kairi/ds-algo/ds/stack"
)

func main() {
	s := stack.New[int64]()
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Push(3)
	fmt.Println(s.Peek())
	s.Pop()
	s.Pop()

	s.String()
}
