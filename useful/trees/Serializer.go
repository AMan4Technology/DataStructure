package trees

import (
    "DataStructure/tree/binary"
    "DataStructure/useful/common"
    "DataStructure/useful/serialize"
)

func init() {
    _ = serialize.Register(binarytree.SerializerOfNode, nodeSerializer{}, false)
}

type nodeSerializer struct{}

func (nodeSerializer) Serialize(value serialize.Serializable) (data string, err error) {
    var (
        root  = value.(*binarytree.Node)
        nodes serialize.Strings
    )
    PreOrder(root, func(node *binarytree.Node) bool {
        data, _ := serialize.Serialize(node.Value.(serialize.Serializable))
        nodes = append(nodes, data)
        if node.IsLeaf() {
            nodes = append(nodes, common.Nil, common.Nil)
        }
        return true
    })
    return serialize.Serialize(nodes)
}

func (nodeSerializer) Deserialize(data string) (value serialize.Serializable, err error) {
    values, err := serialize.Deserialize(data)
    if err != nil {
        return
    }
    var (
        nodes = values.(serialize.Strings)
        root  = new(binarytree.Node)
    )
    deserialize(root, nodes, new(int))
    return root, nil
}

func deserialize(node *binarytree.Node, nodes serialize.Strings, i *int) {
    value, _ := serialize.Deserialize(nodes[*i])
    node.Value = value.(binarytree.Value)
    if *i++; nodes[*i] != common.Nil {
        node.Left = new(binarytree.Node)
        deserialize(node.Left, nodes, i)
    }
    if *i++; nodes[*i] != common.Nil {
        node.Right = new(binarytree.Node)
        deserialize(node.Right, nodes, i)
    }
}
