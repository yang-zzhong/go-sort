package ysort

type Comparable interface {
	Less(interface{}) bool
}

func swap(a []Comparable, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}