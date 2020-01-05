package trees

import (
    "reflect"

    "github.com/AMan4Technology/Serialize"
    "github.com/AMan4Technology/Serialize/codec"

    "github.com/AMan4Technology/DataStructure/tree/binary"
    "github.com/AMan4Technology/DataStructure/useful/common"
)

func init() {
    _ = serialize.Register(reflect.TypeOf(binary_tree.Tree{}), binaryTreeSerializer{}, false)
}

type binaryTreeSerializer struct{}

func (binaryTreeSerializer) Serialize(value interface{}, tag string) (string, error) {
    var (
        tree  = value.(binary_tree.Tree)
        nodes serialize.StringSlice
    )
    PreOrder(tree.Root(), func(node *binary_tree.Node) bool {
        data, _ := serialize.Serialize(node.Value, codec.String, "", tag)
        nodes = append(nodes, data)
        if node.IsLeaf() {
            nodes = append(nodes, common.Nil, common.Nil)
        }
        return true
    })
    return serialize.Serialize(nodes, codec.String, "nodes", tag)
}

func (binaryTreeSerializer) Deserialize(data string, tag string) (interface{}, error) {
    nodeSlice, _, err := serialize.Deserialize(data, codec.String, tag)
    if err != nil {
        return binary_tree.New(nil), err
    }
    var (
        nodes = nodeSlice.(serialize.StringSlice)
        root  = new(binary_tree.Node)
    )
    tree := binary_tree.New(root)
    deserialize(root, nodes, new(int), tag)
    return tree, nil
}

func deserialize(node *binary_tree.Node, nodes serialize.StringSlice, i *int, tag string) {
    value, _, err := serialize.Deserialize(nodes[*i], codec.String, tag)
    if err != nil {
        return
    }
    node.Value = value.(binary_tree.Value)
    if *i++; nodes[*i] != common.Nil {
        node.Left = new(binary_tree.Node)
        deserialize(node.Left, nodes, i, tag)
    }
    if *i++; nodes[*i] != common.Nil {
        node.Right = new(binary_tree.Node)
        deserialize(node.Right, nodes, i, tag)
    }
}
