package lists

import (
    "github.com/AMan4Technology/DataStructure/list/doubly"
    "github.com/AMan4Technology/DataStructure/list/linked"
    "github.com/AMan4Technology/DataStructure/useful/lists/doubly_lists"
    "github.com/AMan4Technology/DataStructure/useful/lists/linked_lists"
)

func MergeTwoList(l1, l2 List) List {
    switch list1 := l1.(type) {
    case *linked_list.List:
        return linked_lists.MergeTwoList(list1, l2.(*linked_list.List))
    case *doubly_list.List:
        return doubly_lists.MergeTwoList(list1, l2.(*doubly_list.List))
    }
    return nil
}

func MergeLists(lists Lists) List {
    switch realLists := lists.(type) {
    case []*linked_list.List:
        return linked_lists.MergeLists(realLists)
    case []*doubly_list.List:
        return doubly_lists.MergeLists(realLists)
    }
    return nil
}
