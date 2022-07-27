package main

import (
	"fmt"

	"github.com/rama-kairi/ds-algo/ds/slice"
)

func main() {
	s := slice.Slice[int64]{}

	s = s.New(1)
	s = s.Append(2)
	s = s.Append(3)
	s = s.Append(6)
	s = s.Append(4)
	s = s.Append(5)
	s = s.Prepend(0)
	println(s.Len())
	fmt.Println("3rd Index Value", s.Get(3))
	fmt.Println("4th Index Value", s.Get(4))
	s.Print()
	s = s.Reverse()
	s.Print()
}
