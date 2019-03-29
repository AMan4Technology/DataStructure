package sort

func Insertion(data Sortable) {
	for length, i := data.Len(), 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if !data.Less(j, j-1) {
				break
			}
			data.Swap(j, j-1)
		}
	}
}
