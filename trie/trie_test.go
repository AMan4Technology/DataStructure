package trie

import (
    "fmt"
    "testing"
)

func TestTree(t *testing.T) {
    trie := New()
    trie.Save("王大")
    trie.Save("王二")
    trie.Save("常三")
    fmt.Println(trie.Find("王大"))
    fmt.Println(trie.Words("王"))
    trie.Remove("王大", true)
    fmt.Println(trie.root.child("王").child("大"))
    fmt.Println(trie.Len())
    fmt.Println(trie.Find("王大"))
    trie.Save("王五四")
    fmt.Println(trie.RemovePrefix("王", false))
    fmt.Println(trie.root.child("王"))
}

func TestAcTree(t *testing.T) {
    acTree := NewAc()
    acTree.Save("王大牛逼")
    acTree.Save("大牛冲")
    acTree.Save("王二")
    acTree.Save("常三")
    acTree.Match([]rune("王大牛逼大牛冲，今天是棒棒哒，王二还行，常三牛逼"), func(start, end int) bool {
        fmt.Println(start, end)
        return true
    })
    fmt.Println(acTree.root.children["王"].children["大"].children["牛"].next.children)
}
