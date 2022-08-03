package main

import (
	"fmt"

	"github.com/rama-kairi/ds-algo/ds/queue"
)

func main() {
	q := queue.NewQueue[int64]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.String()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.String()
	q.Clear()
	q.String()
}
