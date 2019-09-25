package ysort

type Comparable interface {
	Less(Comparable) bool
}

func swap(arr []Comparable, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}