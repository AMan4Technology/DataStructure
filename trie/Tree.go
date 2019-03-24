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
    if curr := t.find(word); curr != nil {
        return curr.isWord
    }
    return false
}

func (t Tree) Words(prefix string) []string {
    if prefix == "" {
        return t.AllWords()
    }
    if curr := t.find(prefix); curr != nil {
        return curr.words(prefix)
    }
    return nil
}

func (t Tree) RangeWords(prefix string, callback func(word string) bool) {
    if prefix == "" {
        t.RangeAllWords(callback)
        return
    }
    if curr := t.find(prefix); curr != nil {
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
    curr := t.find(word)
    if curr == nil {
        return
    }
    if !curr.isWord {
        return
    }
    curr.isWord = false
    t.length--
    if !deleteNode || !curr.canRemove() {
        return
    }
    t.removeBranch(word)
}

func (t *Tree) RemovePrefix(prefix string, deleteNode bool) (words []string) {
    if prefix == "" {
        return t.RemoveAll(deleteNode)
    }
    curr := t.find(prefix)
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
        t.removeBranch(prefix)
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

func (t Tree) find(word string) (curr *node) {
    if word == "" {
        return nil
    }
    curr = t.root
    for _, r := range []rune(word) {
        curr = curr.child(string(r))
        if curr == nil {
            return nil
        }
    }
    return
}

func (t *Tree) removeBranch(prefix string) {
    var (
        parent = t.root
        curr   *node
    )
    for _, r := range []rune(prefix) {
        curr = parent.child(string(r))
        if curr.canRemove() {
            parent.remove(string(r))
            return
        }
        parent = curr
    }
}
