package trees

import "DataStructure/tree/binary"

// PreOrder 前序遍历树中所有节点
func PreOrder(tree binarytree.Tree, callback func(*binarytree.Node) bool) {
    preOrder(tree.Root(), callback)
}

// InOrder 中序遍历树中所有节点
func InOrder(tree binarytree.Tree, callback func(*binarytree.Node) bool) {
    inOrder(tree.Root(), callback)
}

// PostOrder 后序遍历树中所有节点
func PostOrder(tree binarytree.Tree, callback func(*binarytree.Node) bool) {
    postOrder(tree.Root(), callback)
}

func preOrder(node *binarytree.Node, callback func(*binarytree.Node) bool) bool {
    if node != nil && (!callback(node) ||
      !preOrder(node.Left, callback) ||
      !preOrder(node.Right, callback)) {
        return false
    }
    return true
}

func inOrder(node *binarytree.Node, callback func(*binarytree.Node) bool) bool {
    if node != nil && (!inOrder(node.Left, callback) ||
      !callback(node) ||
      !inOrder(node.Right, callback)) {
        return false
    }
    return true
}

func postOrder(node *binarytree.Node, callback func(*binarytree.Node) bool) bool {
    if node != nil && (!postOrder(node.Left, callback) ||
      !postOrder(node.Right, callback)) ||
      !callback(node) {
        return false
    }
    return true
}
