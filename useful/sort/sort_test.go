package sort

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	data := ints{3, 5, 2, 5, 2, 8, 9}
	Insertion(data)
	fmt.Println(data)
}

func TestQuick(t *testing.T) {
	data := ints{3, 5, 2, 5, 2, 8, 9}
	Quick(data)
	fmt.Println(data)
}

func TestTopK(t *testing.T) {
	data := ints{3, 5, 2, 5, 2, 8, 9}
	fmt.Println(data[TopK(data, 3):])
}

func TestBucket(t *testing.T) {
	data := ints{3, 5, 2, 5, 2, 8, 9}
	Bucket(data)
	fmt.Println(data)
}

func TestBottomK(t *testing.T) {
	data := ints{3, 5, 2, 5, 2, 8, 9}
	fmt.Println(data[:BottomK(data, 3)])
}

type ints []int

func (is ints) Len() int {
	return len(is)
}

func (is ints) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is ints) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
