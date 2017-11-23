package structures

// Node is the struct which encompasses data being inserted
// into a list.
type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

// TreeNode represents the node and its associations as is
// necessitated by a binary tree
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}
