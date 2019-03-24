package bitmap

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bM := New(10)
	fmt.Println(bM.Have(1))
	bM.Save(1)
	bM.Save(9)
	bM.Save(16)
	fmt.Println(bM.Have(16))
	bM.Remove(16)
	fmt.Println(bM.length, bM.capacity)
	fmt.Println(bM.Have(16))
}
