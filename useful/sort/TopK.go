package sort

func TopK(data Sortable, k int) (startIndex int) {
	length := data.Len()
	if length == 0 || length <= k {
		return
	}
	for start, end, pivot := 0, length-1, findPivot(data, 0, length-1); length-pivot != k && length-pivot != k+1; pivot = findPivot(data, start, end) {
		if length-pivot < k {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
	}
	return length - k
}
