package ysort

func QuickSort(a []Comparable) {
	quicksort(a, 0, len(a) - 1)
}

func quicksort(a []Comparable, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := partition(a, lo, hi)
	quicksort(a, lo, mid)
	quicksort(a, mid + 1, hi)
}

func partition(a []Comparable, lo, hi int) int {
	i := lo + 1
	j := hi
	v := a[lo]
	for {
		// find the index of > v
		for i <= hi && a[i].Less(v) {
			i++
		}
		// find the index of < v
		for j >= lo && v.Less(a[j]) {
			j--
		}
		if j <= i {
			break
		}
		swap(&a[i], &a[j])
	}
	swap(&a[lo], &a[j])
	return j
}