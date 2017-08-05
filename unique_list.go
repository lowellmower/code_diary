package main

func (list *List) Insert(data interface{}) {
	n := &Node{Data: data}
	if list.isEmpty() {
		list.Head = n
		return
	}
	node := list.Head
	for {
		if node.isLast() {
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

func (list *List) Unique() bool {
	tracker := make(map[interface{}]bool)
	node := list.Head
	if node == nil || node.isLast() {
		return true
	}
	for {
		tracker[node.Data] = true
		if node.isLast() {
			break
		}
		if tracker[node.Next.Data] {
			return false
		}
		node = node.Next
	}
	return true
}

func (list *List) isEmpty() bool {
	return list.Head == nil
}

func (n *Node) isLast() bool {
	return n.Next == nil
}

func (n *Node) isFirst() bool {
	return n.Prev == nil
}

func initNode(data interface{}, prev, next *Node) (n *Node) {
	n.Data = data
	n.Prev = prev
	n.Next = next
	return
}
