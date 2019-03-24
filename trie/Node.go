package trie

func newNode(value string) *node {
    return &node{
        value:    value,
        children: make(map[string]*node)}
}

type node struct {
    value    string
    isWord   bool
    children map[string]*node
}

func (n *node) child(value string) *node {
    return n.children[value]
}

func (n *node) countOfWords() (count int) {
    n.rangeWords("", func(word string, curr *node) bool {
        count++
        return true
    })
    return
}

func (n *node) words(prefix string) (words []string) {
    n.rangeWords(prefix, func(word string, curr *node) bool {
        words = append(words, word)
        return true
    })
    return
}

func (n *node) save(value string) (child *node) {
    child = n.children[value]
    if child == nil {
        child = newNode(string(value))
        n.children[value] = child
    }
    return
}

func (n *node) canRemove() (can bool) {
    can = true
    n.rangeWords("", func(word string, curr *node) bool {
        can = false
        return false
    })
    return
}

func (n *node) remove(child string) {
    delete(n.children, child)
}

func (n *node) rangeWords(prefix string, callback func(word string, curr *node) bool) bool {
    if n.isWord && !callback(prefix, n) {
        return false
    }
    for _, child := range n.children {
        if !child.rangeWords(prefix+child.value, callback) {
            return false
        }
    }
    return true
}
