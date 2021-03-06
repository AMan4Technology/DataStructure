package trie

import array_queue "github.com/AMan4Technology/DataStructure/queue/array"

func NewAc() AcTree {
    return AcTree{Tree: &Tree{root: newNode('/')}}
}

type AcTree struct {
    root *acNode
    *Tree
}

func (t *AcTree) Match(text []rune, callback func(start, end int) bool) {
    if t.root == nil {
        t.BuildFailPointer()
    }
    for prev, length, i := t.root, len(text), 0; i < length; i++ {
        for value := text[i]; prev != nil; prev = prev.next {
            curr := prev.children[value]
            if curr != nil {
                prev = curr
                goto word
            }
        }
        prev = t.root
        continue
    word:
        for next := prev; next != t.root; next = next.next {
            if next.node.isWord && !callback(i+1-next.length, i+1) {
                return
            }
        }
    }
}

func (t *AcTree) BuildFailPointer() {
    t.root = newAcNode(t.Tree.root)
    queue := array_queue.New(t.length)
    queue.Enqueue(t.root, true)
    for count := 1; !queue.Empty(); count++ {
        for lenOfLayer, i := queue.Len(), 0; i < lenOfLayer; i++ {
            acNode := queue.Dequeue().(*acNode)
            for value, child := range acNode.node.children {
                childAcNode := newAcNode(child)
                childAcNode.length = count
                childAcNode.next = t.root
                for curr := acNode.next; curr != nil; curr = curr.next {
                    if nextNode := curr.children[value]; nextNode != nil {
                        childAcNode.next = nextNode
                        break
                    }
                }
                acNode.children[value] = childAcNode
                queue.Enqueue(childAcNode, true)
            }
        }
    }
}
