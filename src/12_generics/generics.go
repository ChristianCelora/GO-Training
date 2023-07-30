package main

import "fmt"

// a map key requires comparable type
func getKeys[K comparable, T any](m map[K]T) []K {
	keys := make([]K, len(m))
	i := 0 

	// map elements are iterated in random order
	for k := range m {
		fmt.Println(k)
		keys[i] = k
		i++
	}

	return keys
}

// Linked list example
// any is an alias for interface{}
type List[T any] struct {
	head *Node[T]
}

func (l List[T]) print() {
	fmt.Println("Print list")
	node := l.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

type Node[T any] struct {
	value T
	next *Node[T]
}

func (node *Node[T]) push(value T) *Node[T] {
	if (node.next == nil) {
		node.next = &Node[T]{}
	}
	node.next.value = value
	fmt.Println(node.next, value)

	return node.next
}

func main() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Println(getKeys(m))

	// Linked list 
	head := Node[string]{
		value: "abc",
	}
	head.push("def").push("ghi").push("jkl")

	linked_list := List[string]{
		head: &head,
	}
	linked_list.print()
}