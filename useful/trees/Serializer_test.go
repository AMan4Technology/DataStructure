package trees

import (
    "fmt"
    "strconv"
    "testing"

    "github.com/AMan4Technology/DataStructure/internal"
    "github.com/AMan4Technology/DataStructure/tree/binary"
    "github.com/AMan4Technology/DataStructure/useful/compare"
    "github.com/AMan4Technology/DataStructure/useful/serialize"
)

func TestNodeSerializer(t *testing.T) {
    root := &binarytree.Node{Value: val(1)}
    left := &binarytree.Node{Value: val(2)}
    right := &binarytree.Node{Value: val(3)}
    rightLeft := &binarytree.Node{Value: val(4)}
    rightRight := &binarytree.Node{Value: val(5)}
    root.Left = left
    root.Right = right
    right.Left = rightLeft
    right.Right = rightRight
    data, _ := serialize.Serialize(root)
    fmt.Println(data)
    root2, _ := serialize.Deserialize(data)
    fmt.Println(root == root2.(*binarytree.Node))
    fmt.Println(serialize.Serialize(root2))
    PreOrder(root2.(*binarytree.Node), func(node *binarytree.Node) bool {
        fmt.Println(node.Value)
        return true
    })
    fmt.Println(serialize.NumOfSerializers())
    serialize.RangeSerializers(func(name string) bool {
        fmt.Println(name)
        return true
    })
}

const SerializerOfVal = "test.integer"

type val int

func (v val) Compare(b internal.Comparable) int8 {
    if v < b.(val) {
        return compare.LessThan
    } else if v > b.(val) {
        return compare.MoreThan
    }
    return compare.EqualTo
}

func (val) SerializerName() string {
    return SerializerOfVal
}

func init() {
    _ = serialize.Register(SerializerOfVal, intValSerializer{}, false)
}

type intValSerializer struct{}

func (intValSerializer) Serialize(value serialize.Serializable) (data string, err error) {
    return strconv.Itoa(int(value.(val))), nil
}

func (intValSerializer) Deserialize(data string) (value serialize.Serializable, err error) {
    i, err := strconv.Atoi(data)
    return val(i), err
}
