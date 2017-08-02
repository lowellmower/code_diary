package main

import (
	"fmt"
)

func (list *List) Each(fn nodeFunc) {
	node := list.Head
	for {
		fn(node)
		if node == nil || node.IsLast() {
			break
		}
		node = node.Next
	}
}

func (list *List) Insert(data interface{}) {
	n := &Node{Data: data}
	if list.IsEmpty() {
		list.Head = n
		return
	}
	node := list.Head
	for {
		if node.IsLast() {
			// linked list
			node.Next = n
			// doubly linked list
			n.Prev = node
			break
		}
		node = node.Next
	}
	return
}

func (list *List) IsEmpty() bool {
	return list.Head == nil
}

func (list *List) Unique() bool {
	tracker := make(map[interface{}]bool)
	node := list.Head
	if node == nil || node.IsLast() {
		return true
	}
	for {
		tracker[node.Data] = true
		if tracker[node.Next.Data] {
			return false
		}
		if node.IsLast() {
			break
		}
		node = node.Next
	}
	return true
}

func (list *List) remove(n *Node) {
}

func (list *List) RemoveDuplicates() {
	tracker := make(map[interface{}]bool)
	node := list.Head
	if node == nil || node.IsLast() {
		return
	}
	for {
		tracker[node.Data] = true
		if tracker[node.Next.Data] {
			list.remove(node)
		}
		if node.IsLast() {
			break
		}
		node = node.Next
	}
}

func (n *Node) IsLast() bool {
	return n.Next == nil
}

func (n *Node) IsFirst() bool {
	return n.Prev == nil
}

func InitNode(data interface{}, prev, next *Node) (n *Node) {
	n.Data = data
	n.Prev = prev
	n.Next = next
	return
}

func Printer(n *Node) {
	fmt.Printf("NODE: %#v", n)
}

func DataPrinter(n *Node) {
	fmt.Printf("%v", n.Data)
}

func main() {
	list := new(List)
	list.Insert(1)
	list.Insert(2)
	list.Insert(2)
	list.Insert(3)
	// list.Each(Printer)
	list.Each(DataPrinter)
	fmt.Println(list.Unique())
	list.RemoveDuplicates()
	// list.Each(DataPrinter)
	fmt.Println(list.Unique())

	fmt.Printf("\nLIST: %#v\n", list)
}
