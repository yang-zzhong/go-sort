package ysort

type Comparable interface {
	Less(interface{}) bool
}

func swap(a, b *Comparable) {
	tmp := *a
	*a = *b
	*b = tmp
}