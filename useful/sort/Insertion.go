package sort

func Insertion(data []int) {
    for length, i := len(data), 1; i < length; i++ {
        curr := data[i]
        j := i - 1
        for ; j >= 0; j-- {
            if data[j] <= curr {
                break
            }
            data[j+1] = data[j]
        }
        data[j+1] = curr
    }
}
