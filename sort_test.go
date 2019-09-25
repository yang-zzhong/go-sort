package ysort

import (
	"math/rand"
	"fmt"
)

type cint int

func (c cint) Less( a interface{}) bool {
	return c < a.(cint)
}

func input(l int) []Comparable {
	arr := make([]Comparable, l)
	for i := 0; i < l; i++ {
		arr[i] = cint(rand.Intn(1000))
	}
	return arr
}

func show(arr []Comparable) {
	for i := 0; i < len(arr); i++ {
		if i > 0 && i % 10 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d\t", arr[i])
	}
	fmt.Printf("\n")
}