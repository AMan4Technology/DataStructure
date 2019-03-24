package compare

const (
    LessThan = iota - 1
    EqualTo
    MoreThan
)

func Compare(a, b Value, name string) int8 {
    return Float64(a.Value(name), b.Value(name))
}

func Float64(a, b float64) int8 {
    if a > b {
        return MoreThan
    } else if a < b {
        return LessThan
    }
    return EqualTo
}
