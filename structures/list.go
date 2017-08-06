package structures

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

func (list *List) Insert(data interface{}) {
	n := &Node{Data: data}
	if list.IsEmpty() {
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
	if list.IsEmpty() || node.isLast() {
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

func (list *List) Length() int {
	var counter int
	if list.IsEmpty() {
		return counter
	}
	n := list.Head
	for {
		counter++
		if n.isLast() {
			return counter
		}
		n = n.Next
	}
}

func (list *List) Remove(v interface{}) {
	node := list.Head
	if list.IsEmpty() {
		return
	}

	if node.isLast() && node.Data == v {
		list.Head = node.Next
		return
	}

	for {
		// [ 0, 0 ]
		next := node.Next
		if next == nil {
			return
		}
		if next.isLast() && next.Data == v {
			node.Next = next.Next
		}
		if node.isLast() && node.Data == v && node == list.Head {
			list.Head = nil
			return
		}
		if node.isLast() {
			return
		}
		if node.Data == v {
			node.Data = next.Data
			node.Next = next.Next
		}
		node = next
	}
}

func (list *List) IsEmpty() bool {
	return list.Head == nil
}

func (n *Node) isLast() bool {
	return n.Next == nil
}

func (n *Node) isFirst() bool {
	return n.Prev == nil
}

func initNode() *Node {
	return new(Node)
}
