package main

type nodeFunc func(*Node)

// List is the struct for building a collection of nodes
// which could then, by their fields and attributes determine
// the list type, e.g. linked list, doubly linked list,
// circularly linked list, etc.
type List struct {
	Head *Node
}

// Node is the struct which encompasses data being inserted
// into a list.
type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}


