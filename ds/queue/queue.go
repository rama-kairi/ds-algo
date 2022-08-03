package queue

import "fmt"

// # Queue - Data Structure

// Similar to Stack, Queue is a linear data structure that follows a particular order in which the operations are performed for storing data. The order is First In First Out (FIFO).  One can imagine a queue as a line of people waiting to receive something in sequential order which starts from the beginning of the line. It is an ordered list in which insertions are done at one end which is known as the rear and deletions are done from the other end known as the front. A good example of a queue is any queue of consumers for a resource where the consumer that came first is served first.

// The difference between stacks and queues is in removing. In a stack we remove the item the most recently added; in a queue, we remove the item the least recently added.

// ## The usages of the Queue are:
// - As a stack (LIFO)
// - As a queue (FIFO)
// - As a priority queue
// - As a queue of jobs to be executed
// - As a queue of jobs to be executed in a particular order
// - As a queue of jobs to be executed in a particular order, with a time limit
// - As a queue of jobs to be executed in a particular order, with a time limit, and with a time-out mechanism

// ## Some Real World Usages:
// - Operating System uses queues for job scheduling.
// - To handle congestion in the networking queue can be used.
// - Data packets in communication are arranged in queue format.
// - Sending an e-mail, it will be queued.
// - Server while responding to request
// - Uploading and downloading photos, first kept for uploading/downloading will be completed first (Not if there is threading)
// - Most internet requests and processes use queue.
// - While switching multiple applications, windows use circular queue.
// - In Escalators, Printer spooler, Car washes queue.
// - A circular queue is used to maintain the playing sequence of multiple players in a game.
// - A queue can be implemented in - Linked List-based Queue, Array-based Queue, Stack-based Queue.
// - Uploading and downloading photos, first kept for uploading/downloading will be completed first (Not if there is threading).
// - Handle website traffic
// - CPU scheduling

// ## Operations on Queue:

// - enqueue : Inserts an element at the end of the queue i.e. at the rear end.
// - dequeue: This operation removes and returns an element that is at the front end of the queue.
// - front: This operation returns the element at the front end without removing it.
// - rear: This operation returns the element at the rear end without removing it.
// - isEmpty: This operation indicates whether the queue is empty or not.
// - int size: This operation returns the size of the queue i.e. the total number of elements it contains.

// ## Types of Queues:

// - Simple Queue: Simple queue also known as a linear queue is the most basic version of a queue. Here, insertion of an element i.e. the Enqueue operation takes place at the rear end and removal of an element i.e. the Dequeue operation takes place at the front end.

// - Circular Queue:  In a circular queue, the element of the queue act as a circular ring. The working of a circular queue is similar to the linear queue except for the fact that the last element is connected to the first element. Its advantage is that the memory is utilized in a better way. This is because if there is an empty space i.e. if no element is present at a certain position in the queue, then an element can be easily added at that position.

// - Priority Queue: This queue is a special type of queue. Its specialty is that it arranges the elements in a queue based on some priority. The priority can be something where the element with the highest value has the priority so it creates a queue with decreasing order of values. The priority can also be such that the element with the lowest value gets the highest priority so in turn it creates a queue with increasing order of values.

// - Dequeue: Dequeue is also known as Double Ended Queue. As the name suggests double ended, it means that an element can be inserted or removed from both the ends of the queue unlike the other queues in which it can be done only from one end. Because of this property it may not obey the First In First Out property.

// ## Applications of Queue:
// Queue is used when things don’t have to be processed immediately, but have to be processed in First In First Out order like Breadth First Search. This property of Queue makes it also useful in following kind of scenarios.

// 1) When a resource is shared among multiple consumers. Examples include CPU scheduling, Disk Scheduling.
// 2) When data is transferred asynchronously (data not necessarily received at same rate as sent) between two processes. Examples include IO Buffers, pipes, file IO, etc.
// 3) Queue can be used as an essential component in various other data structures.

// ## Array implementation Of Queue

// For implementing queue, we need to keep track of two indices, front and rear. We enqueue an item at the rear and dequeue an item from the front. If we simply increment front and rear indices, then there may be problems, the front may reach the end of the array. The solution to this problem is to increase front and rear in circular manner.

// ## Steps for ENQUEUE

// 1) Check the queue is full or not
// 2) If full, print overflow and exit
// 3) If queue is not full, increment tail and add the element

// ## Steps for DEQUEUE

// 1) Check queue is empty or not
// 2) if empty, print underflow and exit
// 3) if not empty, print element at the head and increment head
type Queue[T comparable] struct {
	items []T
	size  int
}

// NewQueue returns a new Queue.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue - adds an item to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
	q.size++
}

// Dequeue - removes and returns the item at the front of the queue.
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var empty T
		return empty
	}

	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

// Peek - returns the item at the front of the queue without removing it.
func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		var empty T
		return empty
	}
	return q.items[0]
}

// IsEmpty - returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size - returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// Clear - removes all items from the queue.
func (q *Queue[T]) Clear() {
	q.items = nil
	q.size = 0
}

// String - returns a string representation of the queue.
func (q *Queue[T]) String() {
	fmt.Println(q.items)
}
