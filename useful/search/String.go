package search

import "DataStructure/useful/maxmin"

func BM(main, sub []rune, callback func(index int) bool) {
    var (
        lenOfS   = len(main)
        lenOfSub = len(sub)
    )
    if lenOfS == 0 || lenOfSub == 0 || lenOfS < lenOfSub {
        return
    }
    var (
        indexWithRune                     = indexWithRuneOf(sub)
        indexWithSuffix, prefixWithSuffix = indexWithSuffixOf(sub)
    )
next:
    for i := 0; i <= lenOfS-lenOfSub; {
        for j := lenOfSub - 1; j >= 0; j-- {
            if main[i+j] != sub[j] {
                i += maxmin.MaxOfTwoInt(badCharRule(j, main[i+j], indexWithRune),
                    goodSuffixRule(j, lenOfSub-1-j, indexWithSuffix, prefixWithSuffix))
                continue next
            }
        }
        if !callback(i) {
            break
        }
        i++
    }
}

func indexWithRuneOf(rs []rune) (indexWithRune map[rune]int) {
    indexWithRune = make(map[rune]int)
    for i, r := range rs {
        indexWithRune[r] = i
    }
    return
}

func badCharRule(iOfSub int, r rune, indexWithRune map[rune]int) int {
    if index, ok := indexWithRune[r]; ok {
        return iOfSub - index
    }
    return iOfSub + 1
}

func indexWithSuffixOf(rs []rune) (indexWithSuffix map[int]int, prefixWithSuffix map[int]bool) {
    length := len(rs)
    indexWithSuffix = make(map[int]int, length-1)
    prefixWithSuffix = make(map[int]bool, length-1)
next:
    for i := 0; i < length-1; i++ {
        for k := 1; k <= i+1; k++ {
            if rs[i+1-k] != rs[length-k] {
                continue next
            }
            indexWithSuffix[k] = i + 1 - k
        }
        prefixWithSuffix[i+1] = true
    }
    return
}

func goodSuffixRule(iOfSub, lenOfSuffix int, indexWithSuffix map[int]int, prefixWithSuffix map[int]bool) int {
    if lenOfSuffix == 0 {
        return 1
    }
    if index, ok := indexWithSuffix[lenOfSuffix]; ok {
        return iOfSub - index + 1
    }
    for j := lenOfSuffix - 1; j >= 1; j-- {
        if prefixWithSuffix[j] {
            return iOfSub + 1 + lenOfSuffix - j
        }
    }
    return iOfSub + 1 + lenOfSuffix
}
