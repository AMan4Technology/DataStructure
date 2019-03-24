package search

import (
    "fmt"
    "testing"
)

func TestBM(t *testing.T) {
    BM([]rune("ababab"), []rune("abab"), func(index int) bool {
        fmt.Println(index)
        return true
    })
}
