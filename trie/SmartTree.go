package trie

func NewSmartTree() SmartTree {
    return SmartTree{root: newSmartNode([]rune(""), false)}
}

type SmartTree struct {
    length int
    root   *smartNode
}

func (t SmartTree) Len() int {
    return t.length
}

func (t SmartTree) Find(word string) bool {
    if exist, _, curr, index := t.find(word); exist {
        return index == curr.len() && curr.isWord
    }
    return false
}

func (t SmartTree) Words(prefix string) []string {
    if prefix == "" {
        return t.AllWords()
    }
    if exist, _, curr, index := t.find(prefix); exist {
        return curr.words(prefix + string(curr.value[index:]))
    }
    return nil
}

func (t SmartTree) RangeWords(prefix string, callback func(word string) bool) {
    if prefix == "" {
        t.RangeAllWords(callback)
        return
    }
    if exist, _, curr, index := t.find(prefix); exist {
        curr.rangeWords(prefix+string(curr.value[index:]), func(word string, curr *smartNode) bool {
            return callback(word)
        })
    }
}

func (t SmartTree) AllWords() []string {
    return t.root.words("")
}

func (t SmartTree) RangeAllWords(callback func(word string) bool) {
    t.root.rangeWords("", func(word string, curr *smartNode) bool {
        return callback(word)
    })
}

func (t *SmartTree) Save(word string) {
    if word == "" {
        return
    }
    curr := t.root.save([]rune(word))
    curr.isWord = true
    t.length++
}

func (t *SmartTree) Remove(word string, deleteNode bool) {
    exist, prev, curr, index := t.find(word)
    if !exist {
        return
    }
    if index != curr.len() || !curr.isWord {
        return
    }
    curr.isWord = false
    t.length--
    if deleteNode && curr.canRemove() {
        prev.remove(curr.value[0])
    }
}

func (t *SmartTree) RemovePrefix(prefix string, deleteNode bool) (words []string) {
    if prefix == "" {
        return t.RemoveAll(deleteNode)
    }
    exist, prev, curr, index := t.find(prefix)
    if !exist {
        return nil
    }
    curr.rangeWords(prefix+string(curr.value[index:]), func(word string, curr *smartNode) bool {
        words = append(words, word)
        curr.isWord = false
        t.length--
        return true
    })
    if deleteNode {
        prev.remove(curr.value[0])
    }
    return
}

func (t *SmartTree) RemoveAll(deleteNode bool) (words []string) {
    if deleteNode {
        words = t.AllWords()
        *t = NewSmartTree()
        return
    }
    t.root.rangeWords("", func(word string, curr *smartNode) bool {
        curr.isWord = false
        t.length--
        return true
    })
    return
}

func (t SmartTree) find(word string) (exist bool, prev, curr *smartNode, index int) {
    if word == "" {
        return
    }
    curr = t.root
    runes := []rune(word)
    for i := 0; i < len(runes); {
        prev, curr = curr, curr.child(runes[i])
        if curr == nil {
            return
        }
        exist, aIncludeB, index := curr.include(runes[i:])
        if !exist {
            return exist, prev, curr, index
        }
        if aIncludeB {
            return true, prev, curr, index
        }
        i += index
    }
    return
}
