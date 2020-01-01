package trie

func New() Tree {
    return Tree{root: newNode("")}
}

type Tree struct {
    length int
    root   *node
}

func (t Tree) Len() int {
    return t.length
}

func (t Tree) Find(word string) bool {
    if _, curr := t.find(word); curr != nil {
        return curr.isWord
    }
    return false
}

func (t Tree) Words(prefix string) []string {
    if prefix == "" {
        return t.AllWords()
    }
    if _, curr := t.find(prefix); curr != nil {
        return curr.words(prefix)
    }
    return nil
}

func (t Tree) RangeWords(prefix string, callback func(word string) bool) {
    if prefix == "" {
        t.RangeAllWords(callback)
        return
    }
    if _, curr := t.find(prefix); curr != nil {
        curr.rangeWords(prefix, func(word string, curr *node) bool {
            return callback(word)
        })
    }
}

func (t Tree) AllWords() []string {
    return t.root.words("")
}

func (t Tree) RangeAllWords(callback func(word string) bool) {
    t.root.rangeWords("", func(word string, curr *node) bool {
        return callback(word)
    })
}

func (t *Tree) Save(word string) {
    if word == "" {
        return
    }
    curr := t.root
    for _, r := range []rune(word) {
        curr = curr.save(string(r))
    }
    curr.isWord = true
    t.length++
}

func (t *Tree) Remove(word string, deleteNode bool) {
    prev, curr := t.find(word)
    if curr == nil {
        return
    }
    if !curr.isWord {
        return
    }
    curr.isWord = false
    t.length--
    if deleteNode && curr.canRemove() {
        prev.remove(curr.value)
    }
}

func (t *Tree) RemovePrefix(prefix string, deleteNode bool) (words []string) {
    if prefix == "" {
        return t.RemoveAll(deleteNode)
    }
    prec, curr := t.find(prefix)
    if curr == nil {
        return nil
    }
    curr.rangeWords(prefix, func(word string, curr *node) bool {
        words = append(words, word)
        curr.isWord = false
        t.length--
        return true
    })
    if deleteNode {
        prec.remove(curr.value)
    }
    return
}

func (t *Tree) RemoveAll(deleteNode bool) (words []string) {
    if deleteNode {
        words = t.AllWords()
        *t = New()
        return
    }
    t.root.rangeWords("", func(word string, curr *node) bool {
        curr.isWord = false
        t.length--
        return true
    })
    return
}

func (t Tree) find(word string) (prev, curr *node) {
    if word == "" {
        return nil, nil
    }
    curr = t.root
    for _, r := range []rune(word) {
        prev, curr = curr, curr.child(string(r))
        if curr == nil {
            return
        }
    }
    return
}
