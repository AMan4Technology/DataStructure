package sort

import (
    "fmt"
    "testing"
)

func TestInsertion(t *testing.T) {
    data := []int{3, 5, 2, 5, 2, 8, 9}
    Insertion(data)
    fmt.Println(data)
}

func TestQuick(t *testing.T) {
    data := []int{3, 5, 2, 5, 2, 8, 9}
    Quick(data)
    fmt.Println(data)
}

func TestTopK(t *testing.T) {
    data := []int{3, 5, 2, 5, 2, 8, 9}
    fmt.Println(data[TopK(data, 3):])
}

func TestBucket(t *testing.T) {
    data := []int{3, 5, 2, 5, 2, 8, 9}
    Bucket(data)
    fmt.Println(data)
}
