package serialize

import (
    "fmt"
    "strings"

    "DataStructure/useful/common"
)

const Split = "|"

func Serialize(value Serializable) (data string, err error) {
    name := value.SerializerName()
    if serializerWithName[name] == nil {
        return name + Split + common.Nil, fmt.Errorf("serializable %s not exist", name)
    }
    if data, err = serializerWithName[name].Serialize(value); err != nil {
        return name + Split + common.Nil, err
    }
    return name + Split + data, nil
}

func Deserialize(data string) (value Serializable, err error) {
    var (
        position = strings.Index(data, Split)
        name     = data[:position]
    )
    if serializerWithName[name] == nil {
        return nil, fmt.Errorf("serializable %s not exist", name)
    }
    return serializerWithName[name].Deserialize(data[position+1:])
}

type Serializer interface {
    Serialize(value Serializable) (data string, err error)
    Deserialize(data string) (value Serializable, err error)
}
