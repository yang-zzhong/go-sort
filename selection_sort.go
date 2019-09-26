package ysort

func SelectionSort(a []Comparable) {
	for i := 0; i < len(a) - 1; i++ {
		p := i
		for j := i + 1; j < len(a); j++ {
			if a[j].Less(a[p])  {
				p = j
			}
		}
		if i != p {
			swap(a, i, p)
		}
	}
}