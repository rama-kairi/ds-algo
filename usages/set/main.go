package main

import (
	"fmt"

	"github.com/rama-kairi/ds-algo/ds/set"
)

func main() {
	s := set.Set[int64]{}.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Add(6)

	fmt.Println(s.Contains(2))
	fmt.Println(s.Filter(func(x int64) bool { return x%2 == 0 }))

	fmt.Println(s.Size())
}
