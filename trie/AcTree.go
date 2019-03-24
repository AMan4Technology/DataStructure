package trie

import "DataStructure/queue/array"

func NewAc() AcTree {
    return AcTree{Tree: &Tree{root: newNode("")}}
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
        for value := string(text[i]); prev != nil; prev = prev.next {
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
    queue := arrayqueue.New(t.length)
    queue.Enqueue(t.root, true)
    for count := 1; !queue.Empty(); count++ {
        for lenOfLayer, i := queue.Len(), 0; i < lenOfLayer; i++ {
            acNode := queue.Dequeue().(*acNode)
        next:
            for value, child := range acNode.node.children {
                childAcNode := newAcNode(child)
                if childAcNode.node.isWord {
                    childAcNode.length = count
                }
                acNode.children[value] = childAcNode
                queue.Enqueue(childAcNode, true)
                for curr := acNode.next; curr != nil; curr = curr.next {
                    if nextNode := curr.children[value]; nextNode != nil {
                        childAcNode.next = nextNode
                        continue next
                    }
                }
                childAcNode.next = t.root
            }
        }
    }
}
