package bplustree

import (
    "fmt"
    "testing"

    "github.com/AMan4Technology/DataStructure/internal/example"
)

func TestTree(t *testing.T) {
    tree := New(5, 3)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(5), "5", true)
    tree.Save(example.ComparableOf(2), "5", true)
    tree.Save(example.ComparableOf(8), "5", true)
    tree.Save(example.ComparableOf(6), "5", true)
    tree.Save(example.ComparableOf(4), "5", true)
    tree.Save(example.ComparableOf(9), "5", true)
    tree.Save(example.ComparableOf(11), "5", true)
    tree.Save(example.ComparableOf(19), "5", true)
    tree.Save(example.ComparableOf(100), "5", true)
    tree.Save(example.ComparableOf(0), "5", true)
    tree.Save(example.ComparableOf(1), "5", true)

    tree.Remove(example.ComparableOf(100))
    tree.Remove(example.ComparableOf(19))
    tree.Remove(example.ComparableOf(11))
    tree.Remove(example.ComparableOf(9))
    tree.Remove(example.ComparableOf(8))
    tree.Remove(example.ComparableOf(0))
    fmt.Println(tree.Find(example.ComparableOf(5)))
    tree.Remove(example.ComparableOf(5))
    fmt.Println(tree.Find(example.ComparableOf(5)))

    tree.Remove(example.ComparableOf(5))
    fmt.Println(tree.Find(example.ComparableOf(5)))

    tree.Remove(example.ComparableOf(5))
    fmt.Println(tree.Find(example.ComparableOf(5)))

    tree.Remove(example.ComparableOf(5))
    fmt.Println(tree.Find(example.ComparableOf(5)))

    tree.Remove(example.ComparableOf(1))
    tree.Remove(example.ComparableOf(2))
    tree.Remove(example.ComparableOf(4))
    tree.Remove(example.ComparableOf(6))

    tree.Save(example.ComparableOf(2), "5", true)
    tree.Save(example.ComparableOf(8), "5", true)
    tree.Save(example.ComparableOf(6), "5", true)
    tree.Save(example.ComparableOf(4), "5", true)
    tree.Save(example.ComparableOf(9), "5", true)

    tree.Range(nil, nil, func(k key, value DataValue) bool {
        fmt.Println(k, value)
        return true
    })
    fmt.Println("---------------------------")
    tree.Range(example.ComparableOf(0), example.ComparableOf(3), func(k key, value DataValue) bool {
        fmt.Println(k, value)
        return true
    })

}
