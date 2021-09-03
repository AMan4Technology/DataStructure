package trie

func newSmartNode(value []rune, isWord bool) *smartNode {
    return &smartNode{
        value:    value,
        isWord:   isWord,
        children: make(map[rune]*smartNode)}
}

type smartNode struct {
    value    []rune
    isWord   bool
    children map[rune]*smartNode
}

func (n *smartNode) len() int {
    return len(n.value)
}

func (n *smartNode) toString() string {
    return string(n.value)
}

func (n *smartNode) child(value rune) *smartNode {
    return n.children[value]
}

func (n *smartNode) countOfWords() (count int) {
    n.rangeWords("", func(word string, curr *smartNode) bool {
        count++
        return true
    })
    return
}

func (n *smartNode) words(prefix string) (words []string) {
    n.rangeWords(prefix, func(word string, curr *smartNode) bool {
        words = append(words, word)
        return true
    })
    return
}

func (n *smartNode) save(value []rune) (child *smartNode) {
    include, aIncludeB, index := n.include(value)
    if include {
        if aIncludeB {
            child = newSmartNode(n.value[index:], n.isWord)
            n.value, n.isWord = n.value[:index], true
            n.children[child.value[0]] = child
            return
        }
        if child = n.children[value[index]]; child != nil {
            return child.save(value[index:])
        }
        child = newSmartNode(value[index:], true)
        n.children[value[index]] = child
        return
    }
    child = newSmartNode(n.value[index:], n.isWord)
    n.value, n.isWord = n.value[:index], false
    for r, node := range n.children {
	child.children[r] = node
	delete(n.children, r)
    }
    n.children[child.value[0]] = child
    child = newSmartNode(value[index:], true)
    n.children[value[index]] = child
    return
}

func (n *smartNode) canRemove() (can bool) {
    can = true
    n.rangeWords("", func(word string, curr *smartNode) bool {
        can = false
        return false
    })
    return
}

func (n *smartNode) remove(child rune) {
    delete(n.children, child)
    if n.isWord || len(n.children) != 1 {
        return
    }
    for _, child := range n.children {
        n.value, n.isWord, n.children = append(n.value, child.value...), child.isWord, child.children
    }
}

func (n *smartNode) rangeWords(prefix string, callback func(word string, curr *smartNode) bool) bool {
    if n.isWord && !callback(prefix, n) {
        return false
    }
    for _, child := range n.children {
        if !child.rangeWords(prefix+child.toString(), callback) {
            return false
        }
    }
    return true
}

func (n *smartNode) include(value []rune) (can, aIncludeB bool, index int) {
    length := n.len()
    if length >= len(value) {
        length = len(value)
        aIncludeB = true
    }
    for ; index < length; index++ {
        if value[index] != n.value[index] {
            return
        }
    }
    return true, aIncludeB, length
}
