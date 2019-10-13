package search

import (
    "github.com/AMan4Technology/DataStructure/useful/common"

    "github.com/AMan4Technology/DataStructure/useful/compare"
)

func BinarySearch(data Searchable, value common.Value) (exist bool, index int) {
    length := data.Len()
    if length == 0 {
        return false, 0
    }
    for start, end := 0, length-1; start <= end; {
        index = start + (end-start)/2
        switch data.Compare(index, value) {
        case compare.EqualTo:
            return true, index
        case compare.MoreThan:
            end = index - 1
        default:
            index++
            start = index

        }
    }
    return false, index
}

func FirstIndex(data Searchable, value common.Value) (exist bool, index int) {
    if exist, index = BinarySearch(data, value); !exist {
        return
    }
    for i := index - 1; i >= 0; i-- {
        if data.Compare(i, value) != compare.EqualTo {
            return
        }
        index = i
    }
    return
}

func LastIndex(data Searchable, value common.Value) (exist bool, index int) {
    if exist, index = BinarySearch(data, value); !exist {
        return
    }
    for length, i := data.Len(), index+1; i < length; i++ {
        if data.Compare(i, value) != compare.EqualTo {
            return
        }
        index = i
    }
    return
}
