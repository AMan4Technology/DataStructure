package bplus_tree

import (
    "github.com/AMan4Technology/DataStructure/useful/compare"

    common2 "github.com/AMan4Technology/DataStructure/useful/common"

    "github.com/AMan4Technology/DataStructure/list/doubly"
)

func New(numOfSubNode, numOfLeafData int) Tree {
    if numOfSubNode < 2 {
        numOfSubNode = 2
    }
    if numOfLeafData < 1 {
        numOfLeafData = 1
    }
    var (
        root     = newNode(numOfSubNode)
        list     = doubly_list.New()
        dataNode = newDataNode(newLeafNode(numOfLeafData))
    )
    list.Enqueue(dataNode)
    root.insertNode(0, nil, dataNode)
    return Tree{root: root, dataList: list}
}

type Tree struct {
    length   int
    root     *Node
    dataList doubly_list.List
}

func (t Tree) Len() int {
    return t.length
}

func (t Tree) Empty() bool {
    return t.length == 0
}

func (t Tree) NumOfSubNode() int {
    return t.root.num
}

func (t Tree) NumOfLeafData() int {
    return leafNodeOf(t.dataList.Head().Next()).num
}

func (t Tree) Find(k key) (exist bool, value DataValue) {
    var dataNode *doubly_list.Node
    for curr := t.root; curr.Len() > 0; {
        _, index := curr.Find(k)
        if curr.dataNodes != nil {
            dataNode = curr.dataNodes[index]
            break
        }
        curr = curr.children[index]
    }
    leaf := leafNodeOf(dataNode)
    if exist, index := leaf.Find(k); exist {
        return true, leaf.data[index]
    }
    return
}

func (t *Tree) Save(k key, value DataValue, update bool) (origin DataValue) {
    var (
        indexes, nodes, dataNode = t.find(k)
        leaf                     = leafNodeOf(dataNode)
    )
    exist, index := leaf.Find(k)
    if exist {
        origin = leaf.data[index]
        if update {
            leaf.data[index] = value
        }
        return
    }
    origin = common2.Nil
    leaf.insert(index, k, value)
    t.length++
    if !leaf.Overflow() {
        return
    }
    var (
        k2, leaf2  = leaf.split()
        dataNode2  = newDataNode(leaf2)
        lenOfNodes = len(nodes)
    )
    t.dataList.InsertAfter(dataNode, dataNode2)
    nodes[lenOfNodes-1].insertKey(indexes[lenOfNodes-1], k2)
    nodes[lenOfNodes-1].insertNode(indexes[lenOfNodes-1]+1, nil, dataNode2)
    for i := lenOfNodes - 1; i > 0; i-- {
        if !nodes[i].Overflow() {
            return
        }
        k2, node2 := nodes[i].split()
        nodes[i-1].insertKey(indexes[i-1], k2)
        nodes[i-1].insertNode(indexes[i-1]+1, node2, nil)
    }
    if t.root.Overflow() {
        k2, node2 := t.root.split()
        node := newNode(t.root.num)
        node.insertKey(0, k2)
        node.insertNode(0, t.root, nil)
        node.insertNode(1, node2, nil)
        t.root = node
    }
    return
}

func (t *Tree) Update(k key, value DataValue) (origin DataValue) {
    var (
        _, _, dataNode = t.find(k)
        leaf           = leafNodeOf(dataNode)
    )
    if exist, index := leaf.Find(k); exist {
        origin = leaf.data[index]
        leaf.data[index] = value
        return
    }
    return common2.Nil
}

func (t *Tree) Remove(k key) (origin DataValue) {
    var (
        indexes, nodes, dataNode = t.find(k)
        leaf                     = leafNodeOf(dataNode)
    )
    exist, index := leaf.Find(k)
    if !exist {
        return common2.Nil
    }
    origin = leaf.delete(index)
    t.length--
    if !leaf.TooLittle() {
        return
    }
    lenOfNodes := len(nodes)
    if indexes[lenOfNodes-1] == nodes[lenOfNodes-1].Len()-1 {
        if !leaf.Empty() || t.dataList.Len() == 1 {
            return
        }
        t.dataList.DeleteNode(dataNode)
        nodes[lenOfNodes-1].deleteKey(indexes[lenOfNodes-1] - 1)
        nodes[lenOfNodes-1].deleteNode(indexes[lenOfNodes-1])
    } else {
        dataNode2 := nodes[lenOfNodes-1].dataNodes[indexes[lenOfNodes-1]+1]
        if ok, k2 := leaf.fuse(leafNodeOf(dataNode2)); !ok {
            nodes[lenOfNodes-1].keys[indexes[lenOfNodes-1]] = k2
            return
        }
        t.dataList.DeleteNode(dataNode2)
        nodes[lenOfNodes-1].deleteKey(indexes[lenOfNodes-1])
        nodes[lenOfNodes-1].deleteNode(indexes[lenOfNodes-1] + 1)
    }
    for i := lenOfNodes - 1; i > 0; i-- {
        if !nodes[i].TooLittle() {
            return
        }
        if indexes[i-1] == nodes[i-1].Len()-1 {
            if !nodes[i].Empty() || nodes[i] == t.root {
                return
            }
            nodes[i-1].deleteKey(indexes[i-1] - 1)
            nodes[i-1].deleteNode(indexes[i-1])
        } else {
            node2 := nodes[i-1].children[indexes[i-1]+1]
            if nodes[i].Len()+node2.Len() <= node2.MaxNum() {
                nodes[i].insertKey(nodes[i].keys.Len(), nodes[i-1].keys[indexes[i-1]])
                nodes[i-1].deleteKey(indexes[i-1])
                nodes[i-1].deleteNode(indexes[i-1] + 1)
                for j, k := range node2.keys {
                    nodes[i].insertKey(nodes[i].keys.Len(), k)
                    if node2.dataNodes != nil {
                        nodes[i].insertNode(nodes[i].Len(), nil, node2.dataNodes[j])
                        continue
                    }
                    nodes[i].insertNode(nodes[i].Len(), node2.children[j], nil)
                }
                if node2.dataNodes != nil {
                    nodes[i].insertNode(nodes[i].Len(), nil, node2.dataNodes[node2.Len()-1])
                    continue
                }
                nodes[i].insertNode(nodes[i].Len(), node2.children[node2.Len()-1], nil)
            } else {
                var (
                    splitK  = nodes[i-1].keys[indexes[i-1]]
                    fuseNum = nodes[i].MinNum() - nodes[i].Len()
                )
                for j := 0; j < fuseNum; j++ {
                    nodes[i].insertKey(nodes[i].keys.Len(), splitK)
                    splitK = node2.keys[j]
                    if node2.dataNodes != nil {
                        nodes[i].insertNode(nodes[i].Len(), nil, node2.dataNodes[j])
                        continue
                    }
                    nodes[i].insertNode(nodes[i].Len(), node2.children[j], nil)
                }
                nodes[i-1].keys[indexes[i-1]] = splitK
                node2.keys = node2.keys[fuseNum:]
                if node2.dataNodes != nil {
                    node2.dataNodes = node2.dataNodes[fuseNum:]
                    continue
                }
                node2.children = node2.children[fuseNum:]
            }
        }
    }
    for ; t.root.Len() == 1 && t.root.children != nil; t.root = t.root.children[0] {
    }
    return
}

func (t Tree) Range(start, end key, callback func(k key, value DataValue) bool) {
    if t.Empty() {
        return
    }
    if start != nil && end != nil && start.Compare(end) == compare.MoreThan {
        start, end = end, start
    }
    var (
        startI, endI       int
        startNode, endNode *doubly_list.Node
    )
    if start == nil {
        startNode = t.dataList.Head().Next()
    } else {
        _, _, startNode = t.find(start)
        _, startI = leafNodeOf(startNode).Find(start)
    }
    if end == nil {
        endNode = t.dataList.Tail().Prev()
        endI = leafNodeOf(endNode).Len()
    } else {
        _, _, endNode = t.find(end)
        _, endI = leafNodeOf(endNode).Find(end)
    }
    if startNode == endNode {
        leaf := leafNodeOf(startNode)
        for i := startI; i < endI; i++ {
            if !callback(leaf.keys[i], leaf.data[i]) {
                return
            }
        }
        return
    }
    for leaf, i := leafNodeOf(startNode), startI; i < leaf.Len(); i++ {
        if !callback(leaf.keys[i], leaf.data[i]) {
            return
        }
    }
    for curr := startNode.Next(); curr != endNode; curr = curr.Next() {
        leaf := leafNodeOf(curr)
        for i, k := range leaf.keys {
            if !callback(k, leaf.data[i]) {
                return
            }
        }
    }
    for leaf, i := leafNodeOf(endNode), 0; i < endI; i++ {
        if !callback(leaf.keys[i], leaf.data[i]) {
            return
        }
    }
}

func (t Tree) find(k key) (indexes []int, nodes []*Node, dataNode *doubly_list.Node) {
    indexes = make([]int, 0, 1<<2)
    nodes = make([]*Node, 0, 1<<2)
    for curr := t.root; curr.Len() > 0; {
        nodes = append(nodes, curr)
        _, index := curr.Find(k)
        indexes = append(indexes, index)
        if curr.dataNodes != nil {
            dataNode = curr.dataNodes[index]
            break
        }
        curr = curr.children[index]
    }
    return
}
