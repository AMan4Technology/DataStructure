package sort

func BottomK(data Sortable, k int) (end int) {
	length := data.Len()
	if length == 0 || length <= k {
		return
	}
	for start, end, pivot := 0, length-1, findPivot(data, 0, length-1); pivot != k && pivot != k-1; pivot = findPivot(data, start, end) {
		if pivot < k {
			start = pivot + 1
		} else {
			end = pivot - 1
		}
	}
	return k
}
