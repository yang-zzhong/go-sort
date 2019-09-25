package ysort

import (
	"testing"
	"fmt"
)

func TestTopDownMergeSort(t *testing.T) {
	arr := input(100)
	fmt.Printf("before: \n")
	show(arr)
	merge := new(MergeSort)
	merge.TopDown(arr)
	fmt.Printf("after top down merge: \n")
	show(arr)
}

func TestBottomUpMergeSort(t *testing.T) {
	arr := input(100)
	fmt.Printf("before: \n")
	show(arr)
	merge := new(MergeSort)
	merge.BottomUp(arr)
	fmt.Printf("after bottom up merge: \n")
	show(arr)
}