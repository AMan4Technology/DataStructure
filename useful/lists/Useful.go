package lists

import (
    "github.com/AMan4Technology/DataStructure/list/doubly"
    "github.com/AMan4Technology/DataStructure/list/linked"
    "github.com/AMan4Technology/DataStructure/useful/lists/doublylists"
    "github.com/AMan4Technology/DataStructure/useful/lists/linkedlists"
)

func MergeTwoList(l1, l2 List) List {
    switch list1 := l1.(type) {
    case *linkedlist.List:
        return linkedlists.MergeTwoList(list1, l2.(*linkedlist.List))
    case *doublylist.List:
        return doublylists.MergeTwoList(list1, l2.(*doublylist.List))
    }
    return nil
}

func MergeLists(lists Lists) List {
    switch realLists := lists.(type) {
    case []*linkedlist.List:
        return linkedlists.MergeLists(realLists)
    case []*doublylist.List:
        return doublylists.MergeLists(realLists)
    }
    return nil
}
