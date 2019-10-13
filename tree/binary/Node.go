package binarytree

import (
    "errors"

    "github.com/AMan4Technology/DataStructure/internal"
)

const SerializerOfNode = "binaryTree.Node"

type Node struct {
    Value       Value
    Left, Right *Node
}

func (n *Node) Compare(b internal.Comparable) int8 {
    return n.Value.Compare(b.(*Node).Value)
}

func (n *Node) IsLeaf() bool {
    return n.Left == nil && n.Right == nil
}

func (n *Node) LeftDFS() (parent, left *Node) {
    for parent, left = n, n.Left; left != nil && left.Left != nil; {
        parent, left = left, left.Left
    }
    return
}

func (n *Node) RightDFS() (parent, right *Node) {
    for parent, right = n, n.Right; right != nil && right.Right != nil; {
        parent, right = right, right.Right
    }
    return
}

func (n *Node) LinkLeft(left *Node) error {
    if n.Left != nil {
        return errors.New("leftNode is exist")
    }
    n.Left = left
    return nil
}

func (n *Node) LinkRight(right *Node) error {
    if n.Right != nil {
        return errors.New("rightNode is exist")
    }
    n.Right = right
    return nil
}

func (n *Node) DeleteLeft() (node *Node) {
    if node = n.Left; node == nil {
        return
    }
    rightRoot := node.Right
    if rightRoot == nil {
        n.Left = node.Left
        return
    }
    parentOfLeft, left := rightRoot.LeftDFS()
    if left == nil {
        n.Left, rightRoot.Left = rightRoot, node.Left
        return
    }
    parentOfLeft.Left, n.Left = nil, left
    left.Left, left.Right, node.Left, node.Right = node.Left, node.Right, nil, nil
    return
}

func (n *Node) DeleteRight() (node *Node) {
    if node = n.Right; node == nil {
        return
    }
    rightRoot := node.Right
    if rightRoot == nil {
        n.Right = node.Left
        return
    }
    parentOfLeft, left := rightRoot.LeftDFS()
    if left == nil {
        n.Right, rightRoot.Left = rightRoot, node.Left
        return
    }
    parentOfLeft.Left, n.Right = nil, left
    left.Left, left.Right, node.Left, node.Right = node.Left, node.Right, nil, nil
    return
}

func (n *Node) SerializerName() string {
    return SerializerOfNode
}

type Value = internal.Comparable
