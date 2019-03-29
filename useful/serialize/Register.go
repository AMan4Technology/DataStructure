package serialize

import "fmt"

var serializerWithName = make(map[string]Serializer)

func Register(name string, serializer Serializer, update bool) (err error) {
    if serializerWithName[name] != nil && !update {
        return fmt.Errorf("serializable %s is exist", name)
    }
    serializerWithName[name] = serializer
    return nil
}

func NumOfSerializers() int {
    return len(serializerWithName)
}

func RangeSerializers(callback func(name string) bool) {
    for name := range serializerWithName {
        if !callback(name) {
            break
        }
    }
}
