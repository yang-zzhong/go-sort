package ysort

func InsertionSort(a []Comparable) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j].Less(a[j - 1]) {
				swap(a, j - 1, j)
			}
		}
	}
}