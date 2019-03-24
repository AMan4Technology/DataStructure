package maxmin

func OfData(data Data) (iOfMax, iOfMin int) {
    length := data.Len()
    if length == 0 {
        return -1, -1
    }
    for i := 1; i < length; i++ {
        if data.Less(iOfMax, i) {
            iOfMax = i
        }
        if data.Less(i, iOfMin) {
            iOfMin = i
        }
    }
    return
}

func MaxOfData(data Data) (iOfMax int) {
    length := data.Len()
    if length == 0 {
        return -1
    }
    for i := 1; i < length; i++ {
        if data.Less(iOfMax, i) {
            iOfMax = i
        }
    }
    return
}

func MinOfData(data Data) (iOfMin int) {
    length := data.Len()
    if length == 0 {
        return -1
    }
    for i := 1; i < length; i++ {
        if data.Less(i, iOfMin) {
            iOfMin = i
        }
    }
    return
}

type Data interface {
    Len() int
    Less(i, j int) bool
}
