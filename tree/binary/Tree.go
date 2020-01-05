package binary_tree

func New(root *Node) Tree {
    return Tree{root: root}
}

type Tree struct {
    root *Node
}

func (t Tree) Empty() bool {
    return t.root == nil
}

func (t Tree) Root() *Node {
    return t.root
}
