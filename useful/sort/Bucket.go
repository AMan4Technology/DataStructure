package sort

import (
    . "DataStructure/useful/maxmin"
)

func Bucket(data []int) {
    length := len(data)
    if length < 2 {
        return
    }
    var (
        max, min = OfInt(data[0], data[1:]...)
        i        int
    )
    if max-min <= 1<<10 {
        buckets := make([]int, max-min+1)
        for _, value := range data {
            buckets[value-min]++
        }
        for value, nums := range buckets {
            for ; nums != 0; nums-- {
                data[i] = value + min
                i++
            }
        }
        return
    }
    var (
        count   = MinOfTwoInt(1<<10, length)
        size    = (max-min+1)/count + 1
        buckets = make([]bucket, count)
    )
    for _, value := range data {
        buckets[value/size] = append(buckets[value/size], value)
    }
    for _, bucket := range buckets {
        Quick(bucket)
        for _, value := range bucket {
            data[i] = value
            i++
        }
    }
}

type bucket []int

func (b bucket) Len() int {
    return len(b)
}

func (b bucket) Less(i, j int) bool {
    return b[i] < b[j]
}

func (b bucket) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}
