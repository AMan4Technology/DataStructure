package trie

func newAcNode(node *node) *acNode {
    return &acNode{
        node:     node,
        children: make(map[rune]*acNode)}
}

type acNode struct {
    length   int
    node     *node
    next     *acNode
    children map[rune]*acNode
}
