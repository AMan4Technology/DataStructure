package maxmin

func OfInt(base int, values ...int) (max, min int) {
    max, min = base, base
    for i, length := 0, len(values); i < length; i++ {
        max = MaxOfTwoInt(max, values[i])
        min = MinOfTwoInt(min, values[i])
    }
    return
}

func MaxOfInt(base int, values ...int) (max int) {
    max = base
    for i, length := 0, len(values); i < length; i++ {
        max = MaxOfTwoInt(max, values[i])
    }
    return
}

func MinOfInt(base int, values ...int) (min int) {
    min = base
    for i, length := 0, len(values); i < length; i++ {
        min = MinOfTwoInt(min, values[i])
    }
    return
}

func MaxOfTwoInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func MinOfTwoInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}
