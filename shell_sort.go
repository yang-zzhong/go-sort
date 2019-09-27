package ysort

func ShellSort(a []Comparable) {
	N := len(a)
	h := 1
	for h < N / 3 {
		h = h * 3 + 1
	}
	for h > 0 {
		for i := h; i < N; i++ {
			for j := i; j >= h; j -= h {
				if a[j].Less(a[j - h]) {
					swap(&a[j], &a[j - h])
				} else {
					break
				}
			}
		}
		h = h / 3
	}
}