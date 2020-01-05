package trees

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/AMan4Technology/Serialize"
    "github.com/AMan4Technology/Serialize/codec"

    "github.com/AMan4Technology/DataStructure/internal"
    "github.com/AMan4Technology/DataStructure/tree/binary"
    "github.com/AMan4Technology/DataStructure/useful/compare"
)

func init() {
    _ = serialize.Register(reflect.TypeOf(val(0)), nil, true)
}

func TestBinaryTreeSerializer_Serialize(t *testing.T) {
    var (
        root       = &binary_tree.Node{Value: val(1)}
        tree       = binary_tree.New(root)
        left       = &binary_tree.Node{Value: val(2)}
        right      = &binary_tree.Node{Value: val(3)}
        rightLeft  = &binary_tree.Node{Value: val(4)}
        rightRight = &binary_tree.Node{Value: val(5)}
    )
    root.Left = left
    root.Right = right
    right.Left = rightLeft
    right.Right = rightRight
    PreOrder(tree.Root(), func(node *binary_tree.Node) bool {
        fmt.Println(node.Value)
        return true
    })
    data, _ := serialize.Serialize(tree, codec.String, "tree", "")
    fmt.Println(data)
}

func TestBinaryTreeSerializer_Deserialize(t *testing.T) {
    treeValue, _, err := serialize.Deserialize(`github.com/AMan4Technology/DataStructure/tree/binary.Tree|tree|github.com/AMan4Technology/Serialize/internal.StringSlice|nodes|11|60|github.com/AMan4Technology/DataStructure/useful/trees.val||160|github.com/AMan4Technology/DataStructure/useful/trees.val||25|[nil]5|[nil]60|github.com/AMan4Technology/DataStructure/useful/trees.val||360|github.com/AMan4Technology/DataStructure/useful/trees.val||45|[nil]5|[nil]60|github.com/AMan4Technology/DataStructure/useful/trees.val||55|[nil]5|[nil]`, codec.String, "")
    tree := treeValue.(binary_tree.Tree)
    if err != nil {
        fmt.Println(err)
    }
    PreOrder(tree.Root(), func(node *binary_tree.Node) bool {
        fmt.Println(node.Value)
        return true
    })
}

type val int

func (v val) Compare(b internal.Comparable) int8 {
    if v < b.(val) {
        return compare.LessThan
    } else if v > b.(val) {
        return compare.MoreThan
    }
    return compare.EqualTo
}
