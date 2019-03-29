package serialize

import (
    "strconv"
    "strings"
)

const SerializerOfStrings = "serialize.Strings"

func init() {
    _ = Register(SerializerOfStrings, stringsSerializer{}, false)
}

type Strings []string

func (Strings) SerializerName() string {
    return SerializerOfStrings
}

type stringsSerializer struct{}

func (stringsSerializer) Serialize(value Serializable) (data string, err error) {
    var (
        values = value.(Strings)
        length = len(values)
        sb     strings.Builder
    )
    sb.WriteString(strconv.Itoa(length))
    sb.WriteString(Split)
    for _, val := range values {
        sb.WriteString(strconv.Itoa(len(val)))
        sb.WriteString(Split)
        sb.WriteString(val)
    }
    return sb.String(), nil
}

func (stringsSerializer) Deserialize(data string) (value Serializable, err error) {
    var (
        position  = strings.Index(data, Split)
        length, _ = strconv.Atoi(data[:position])
        values    = make(Strings, length)
    )
    position++
    for i := 0; i < length; i++ {
        var (
            offset    = position + strings.Index(data[position:], Split)
            lenOfS, _ = strconv.Atoi(data[position:offset])
        )
        position = offset + 1 + lenOfS
        values[i] = data[offset+1 : position]
    }
    return values, nil
}
