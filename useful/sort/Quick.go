package sort

func Quick(data Sortable) {
    quick(data, 0, data.Len()-1)
}

func quick(data Sortable, start, end int) {
    if start >= end {
        return
    }
    pivot := findPivot(data, start, end)
    quick(data, start, pivot-1)
    quick(data, pivot+1, end)
}

func findPivot(data Sortable, start, end int) (pivot int) {
    pivot = start + (end-start)/2
    compare(data, start, end)
    compare(data, start, pivot)
    compare(data, end, pivot)
    for i, j := start, start; i < end; i++ {
        if !data.Less(end, i) {
            continue
        }
        for j = i + 1; j < end; j++ {
            if data.Less(j, end) {
                break
            }
        }
        data.Swap(i, j)
        if j == end {
            pivot = i
            break
        }
    }
    return
}

func compare(data Sortable, i, j int) {
    if data.Less(j, i) {
        data.Swap(i, j)
    }
}
