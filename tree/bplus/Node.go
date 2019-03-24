package bplustree

import (
    common2 "DataStructure/useful/common"

    "DataStructure/useful/compare"

    "DataStructure/useful/search"

    "DataStructure/internal"
    "DataStructure/list/doubly"
)

func newNode(num int) *Node {
    return &Node{
        common: common{
            num:  num,
            keys: make(keys, 0, num)}}
}

func newDataNode(node *LeafNode) *doublylist.Node {
    return doublylist.NewNode(node)
}

func newLeafNode(num int) *LeafNode {
    return &LeafNode{
        common: common{
            num:  num,
            keys: make(keys, 0, num+1)},
        data: make([]DataValue, 0, num)}
}

func leafNodeOf(node *doublylist.Node) *LeafNode {
    return node.Value.(*LeafNode)
}

type Node struct {
    common // 子节点个数
    children  []*Node
    dataNodes []*doublylist.Node
}

func (n Node) Len() int {
    return len(n.children) + len(n.dataNodes)
}

func (n Node) Empty() bool {
    return n.Len() == 0
}

func (n Node) Full() bool {
    return n.Len() == n.MaxNum()
}

func (n Node) Overflow() bool {
    return n.Len() > n.MaxNum()
}

func (n Node) TooLittle() bool {
    return n.Len() < n.MinNum()
}

func (n *Node) split() (k key, node *Node) {
    node = newNode(n.num)
    pivot := n.Len() / 2
    k = n.keys[pivot]
    node.keys = append(node.keys, n.keys[pivot+1:]...)
    n.keys = n.keys[:pivot]
    if n.children != nil {
        node.children = make([]*Node, 0, node.num+1)
        node.children = append(node.children, n.children[pivot+1:]...)
        n.children = n.children[:pivot+1]
    }
    if n.dataNodes != nil {
        node.dataNodes = make([]*doublylist.Node, 0, node.num+1)
        node.dataNodes = append(node.dataNodes, n.dataNodes[pivot+1:]...)
        n.dataNodes = n.dataNodes[:pivot+1]
    }
    return
}

func (n *Node) insertNode(index int, child *Node, dataNode *doublylist.Node) {
    if index < 0 || index > n.Len() {
        return
    }
    if child != nil {
        n.children = append(n.children, nil)
        for i := n.Len() - 1; i > index; i-- {
            n.children[i] = n.children[i-1]
        }
        n.children[index] = child
    }
    if dataNode != nil {
        n.dataNodes = append(n.dataNodes, nil)
        for i := n.Len() - 1; i > index; i-- {
            n.dataNodes[i] = n.dataNodes[i-1]
        }
        n.dataNodes[index] = dataNode
    }
}

func (n *Node) deleteNode(index int) {
    if index < 0 || index >= n.Len() {
        return
    }
    if n.children != nil {
        n.children = append(n.children[:index], n.children[index+1:]...)
    }
    if n.dataNodes != nil {
        n.dataNodes = append(n.dataNodes[:index], n.dataNodes[index+1:]...)
    }
}

type LeafNode struct {
    common // 当数据存储在磁盘上时，尽可能使得每次读取的数据小于一页
    data []DataValue
}

func (n *LeafNode) split() (k key, node *LeafNode) {
    node = newLeafNode(n.num)
    pivot := n.Len() / 2
    node.keys = append(node.keys, n.keys[pivot:]...)
    node.data = append(node.data, n.data[pivot:]...)
    n.keys, n.data = n.keys[:pivot], n.data[:pivot]
    return n.keys[pivot-1], node
}

func (n *LeafNode) fuse(node *LeafNode) (ok bool, k key) {
    if n.Len()+node.Len() <= n.MaxNum() {
        n.keys = append(n.keys, node.keys...)
        n.data = append(n.data, node.data...)
        return true, nil
    }
    fuseNum := n.MinNum() - n.Len()
    k = node.keys[fuseNum-1]
    n.keys = append(n.keys, node.keys[:fuseNum]...)
    n.data = append(n.data, node.data[:fuseNum]...)
    node.keys, node.data = node.keys[fuseNum:], node.data[fuseNum:]
    return
}

func (n *LeafNode) insert(index int, k key, value DataValue) {
    if index < 0 || index > n.Len() {
        return
    }
    n.keys = append(n.keys, nil)
    n.data = append(n.data, "")
    for i := n.Len() - 1; i > index; i-- {
        n.keys[i], n.data[i] = n.keys[i-1], n.data[i-1]
    }
    n.keys[index], n.data[index] = k, value
}

func (n *LeafNode) delete(index int) (origin DataValue) {
    n.deleteKey(index)
    return n.deleteData(index)
}

func (n *LeafNode) deleteData(index int) (origin DataValue) {
    if index < 0 || index >= n.Len() {
        return common2.Nil
    }
    origin = n.data[index]
    n.data = append(n.data[:index], n.data[index+1:]...)
    return
}

func (LeafNode) Compare(b internal.Comparable) int8 {
    return compare.EqualTo
}

type common struct {
    num  int
    keys keys
}

func (c common) Num() int {
    return c.num
}

func (c common) MinNum() int {
    return c.num / 2
}

func (c common) MaxNum() int {
    return c.num
}

func (c common) Len() int {
    return c.keys.Len()
}

func (c common) Empty() bool {
    return c.Len() == 0
}

func (c common) Full() bool {
    return c.Len() == c.MaxNum()
}

func (c common) Overflow() bool {
    return c.Len() > c.MaxNum()
}

func (c common) TooLittle() bool {
    return c.Len() < c.MinNum()
}

func (c common) Find(k key) (exist bool, index int) {
    return search.BinarySearch(c.keys, k)
}

func (c *common) insertKey(index int, k key) {
    if index < 0 || index > c.Len() {
        return
    }
    c.keys = append(c.keys, nil)
    for i := c.Len() - 1; i > index; i-- {
        c.keys[i] = c.keys[i-1]
    }
    c.keys[index] = k
}

func (c *common) deleteKey(index int) {
    if index < 0 || index >= c.Len() {
        return
    }
    c.keys = append(c.keys[:index], c.keys[index+1:]...)
}
