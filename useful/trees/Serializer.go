package trees

import (
    "reflect"

    "github.com/AMan4Technology/Serialize"
    "github.com/AMan4Technology/Serialize/codec"

    "github.com/AMan4Technology/DataStructure/tree/binary"
    "github.com/AMan4Technology/DataStructure/useful/common"
)

func init() {
    _ = serialize.Register(reflect.TypeOf(binarytree.Tree{}), binaryTreeSerializer{}, false)
}

type binaryTreeSerializer struct{}

func (binaryTreeSerializer) Serialize(value interface{}, tag string) (string, error) {
    var (
        tree  = value.(binarytree.Tree)
        nodes serialize.StringSlice
    )
    PreOrder(tree.Root(), func(node *binarytree.Node) bool {
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
        return binarytree.New(nil), err
    }
    var (
        nodes = nodeSlice.(serialize.StringSlice)
        root  = new(binarytree.Node)
    )
    tree := binarytree.New(root)
    deserialize(root, nodes, new(int), tag)
    return tree, nil
}

func deserialize(node *binarytree.Node, nodes serialize.StringSlice, i *int, tag string) {
    value, _, err := serialize.Deserialize(nodes[*i], codec.String, tag)
    if err != nil {
        return
    }
    node.Value = value.(binarytree.Value)
    if *i++; nodes[*i] != common.Nil {
        node.Left = new(binarytree.Node)
        deserialize(node.Left, nodes, i, tag)
    }
    if *i++; nodes[*i] != common.Nil {
        node.Right = new(binarytree.Node)
        deserialize(node.Right, nodes, i, tag)
    }
}
