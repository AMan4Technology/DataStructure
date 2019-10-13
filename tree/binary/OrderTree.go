package binarytree

import (
    "github.com/AMan4Technology/DataStructure/useful/compare"
)

func NewOrderTree(root *Node) OrderTree {
    return OrderTree{Tree: New(root)}
}

type OrderTree struct {
    Tree
}

func (t OrderTree) Find(value Value) (parent, node *Node) {
    if t.root == nil {
        return
    }
    for node = t.root; node != nil; {
        switch value.Compare(node.Value) {
        case compare.LessThan:
            parent, node = node, node.Left
        case compare.MoreThan:
            parent, node = node, node.Right
        default:
            return
        }
    }
    return
}

func (t *OrderTree) Save(value Value) (node *Node) {
    node = &Node{Value: value}
    if t.root == nil {
        t.root = node
        return
    }
    for curr := t.root; ; {
        if value.Compare(curr.Value) == compare.LessThan {
            if curr.Left == nil {
                curr.Left = node
                return
            }
            curr = curr.Left
            continue
        }
        if curr.Right == nil {
            curr.Right = node
            return
        }
        curr = curr.Right
    }
}

func (t *OrderTree) Remove(value Value) (parent, node *Node) {
    parent, node = t.Find(value)
    if node == nil {
        return
    }
    if parent == nil {
        t.Dequeue()
        return
    }
    if node.Compare(parent) == compare.LessThan {
        parent.DeleteLeft()
    } else {
        parent.DeleteRight()
    }
    return
}

func (t *OrderTree) Dequeue() (root *Node) {
    root = t.root
    rightRoot := root.Right
    if rightRoot == nil {
        t.root = root.Left
        return
    }
    parentOfLeft, left := rightRoot.LeftDFS()
    if left == nil {
        t.root, rightRoot.Left = rightRoot, root.Left
        return
    }
    parentOfLeft.Left, t.root = nil, left
    left.Left, left.Right, root.Left, root.Right = root.Left, root.Right, nil, nil
    return
}
