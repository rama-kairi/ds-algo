package linkedlist

import "fmt"

// # Linked List

// Like arrays, Linked List is a linear data structure. Unlike arrays, linked list elements are not stored at a contiguous location; the elements are linked using pointers. They includes a series of connected nodes. Here, each node stores the data and the address of the next node.
// ![linkedlist_image](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2013/03/Linkedlist.png)

// ## Why Linked List?
// Arrays can be used to store linear data of similar types, but arrays have the following limitations.

// - The size of the arrays is fixed: So we must know the upper limit on the number of elements in advance. Also, generally, the allocated memory is equal to the upper limit irrespective of the usage.

// - Insertion of a new element / Deletion of a existing element in an array of elements is expensive: The room has to be created for the new elements and to create room existing elements have to be shifted but in Linked list if we have the head node then we can traverse to any node through it and insert new node at the required position.

// For example, in a system, if we maintain a sorted list of IDs in an array id[].
// id[] = [1000, 1010, 1050, 2000, 2040].
// And if we want to insert a new ID 1005, then to maintain the sorted order, we have to move all the elements after 1000 (excluding 1000).
// Deletion is also expensive with arrays until unless some special techniques are used. For example, to delete 1010 in id[], everything after 1010 has to be moved due to this so much work is being done which affects the efficiency of the code.

// ## Advantages over arrays:
// - Dynamic Array.
// - Ease of Insertion/Deletion.

// ## Drawbacks:
// - Random access is not allowed. We have to access elements sequentially starting from the first node(head node). So we cannot do binary search with linked lists efficiently with its default implementation. Read about it here.
// - Extra memory space for a pointer is required with each element of the list.
// - Not cache friendly. Since array elements are contiguous locations, there is locality of reference which is not there in case of linked lists.

// ## Representation:
// A linked list is represented by a pointer to the first node of the linked list. The first node is called the head. If the linked list is empty, then the value of the head points to NULL.

// **Each node in a list consists of at least two parts:**

// - A Data Item (we can store integer, strings or any type of data).
// - Pointer (Or Reference) to the next node (connects one node to another) or An address of another node

// In C, we can represent a node using structures. Below is an example of a linked list node with integer data.
// In Java or C#, LinkedList can be represented as a class and a Node as a separate class. The LinkedList class contains a reference of Node class type.

// In simple words, a linked list consists of nodes where each node contains a data field and a reference(link) to the next node in the list.
// Node is a node in a linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList is a linked list.
type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Some Basic Operations on Linked List
// 1. Insertion
// 2. Deletion
// 3. Search
// 4. Traversal

// New - Create a new linked list.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Insert - Insert a new node at the beginning of the linked list.
func (l *LinkedList[T]) Insert(value T) {
	node := &Node[T]{Value: value}
	node.Next = l.Head
	l.Head = node
}

// Delete - Delete the first node from the linked list.
func (l *LinkedList[T]) Delete() {
	l.Head = l.Head.Next
}

// Search - Search for a node with the given value.
// func (l *LinkedList[T]) Search(value T) bool {
// 	for node := l.Head; node != nil; node = node.Next {
// 		if node.Value == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// Traverse - Traverse the linked list.
func (l *LinkedList[T]) Traverse() {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

// Reverse - Reverse the linked list.
func (l *LinkedList[T]) Reverse() {
	var prev *Node[T]
	for node := l.Head; node != nil; node = node.Next {
		node.Next, prev, l.Head = prev, node, node.Next
	}
	l.Head = prev
}

// ReverseRecursive - Reverse the linked list recursively.
func (l *LinkedList[T]) ReverseRecursive() {
	l.Head = l.reverseRecursive(l.Head)
}

func (l *LinkedList[T]) reverseRecursive(node *Node[T]) *Node[T] {
	if node == nil || node.Next == nil {
		return node
	}
	node.Next = l.reverseRecursive(node.Next)
	return node.Next
}

// ReverseKGroup - Reverse the linked list in groups of given size.
func (l *LinkedList[T]) ReverseKGroup(k int) {
	l.Head = l.reverseKGroup(l.Head, k)
}

func (l *LinkedList[T]) reverseKGroup(node *Node[T], k int) *Node[T] {
	if node == nil || node.Next == nil {
		return node
	}
	var prev *Node[T]
	for i := 0; i < k && node != nil; i++ {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	if node != nil {
		node.Next = l.reverseKGroup(node.Next, k)
	}
	return prev
}

// ReverseKGroupRecursive - Reverse the linked list in groups of given size recursively.
func (l *LinkedList[T]) ReverseKGroupRecursive(k int) {
	l.Head = l.reverseKGroupRecursive(l.Head, k)
}

func (l *LinkedList[T]) reverseKGroupRecursive(node *Node[T], k int) *Node[T] {
	if node == nil || node.Next == nil {
		return node
	}
	var prev *Node[T]
	for i := 0; i < k && node != nil; i++ {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	if node != nil {
		node.Next = l.reverseKGroupRecursive(node.Next, k)
	}
	return prev
}
