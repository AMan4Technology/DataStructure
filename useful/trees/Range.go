package trees

import "github.com/AMan4Technology/DataStructure/tree/binary"

// PreOrder 前序遍历树中所有节点
func PreOrder(root *binary_tree.Node, callback func(*binary_tree.Node) bool) {
    preOrder(root, callback)
}

// InOrder 中序遍历树中所有节点
func InOrder(root *binary_tree.Node, callback func(*binary_tree.Node) bool) {
    inOrder(root, callback)
}

// PostOrder 后序遍历树中所有节点
func PostOrder(root *binary_tree.Node, callback func(*binary_tree.Node) bool) {
    postOrder(root, callback)
}

func preOrder(node *binary_tree.Node, callback func(*binary_tree.Node) bool) bool {
    if node != nil && (!callback(node) ||
      !preOrder(node.Left, callback) ||
      !preOrder(node.Right, callback)) {
        return false
    }
    return true
}

func inOrder(node *binary_tree.Node, callback func(*binary_tree.Node) bool) bool {
    if node != nil && (!inOrder(node.Left, callback) ||
      !callback(node) ||
      !inOrder(node.Right, callback)) {
        return false
    }
    return true
}

func postOrder(node *binary_tree.Node, callback func(*binary_tree.Node) bool) bool {
    if node != nil && (!postOrder(node.Left, callback) ||
      !postOrder(node.Right, callback)) ||
      !callback(node) {
        return false
    }
    return true
}
