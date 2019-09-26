package ysort

import (
	"math/rand"
	"testing"
	"fmt"
)

func TestSelectionSort(t * testing.T) {
	test_sort(t, SelectionSort, "selection sort", 100, 10000888, false)
}

func TestInsertionSort(t *testing.T) {
	test_sort(t, InsertionSort, "insertion sort", 100, 1000000, false)
}

func TestShellSort(t *testing.T) {
	test_sort(t, ShellSort, "shell sort", 100, 100000000, false)
}

func TestTopDownMergeSort(t *testing.T) {
	test_sort(t, TopDownMergeSort, "top down merge", 100, 100000000, false)
}

func TestBottomUpMergeSort(t *testing.T) {
	test_sort(t, BottomUpMergeSort, "bottom up merge", 100, 100000000, false)
}

type cint int

func (c cint) Less( a interface{}) bool {
	return c < a.(cint)
}

func input(l, max int) []Comparable {
	arr := make([]Comparable, l)
	for i := 0; i < l; i++ {
		arr[i] = cint(rand.Intn(max))
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

func is_sorted(a []Comparable) bool {
	for i := 0; i < len(a) - 1; i++ {
		if a[i + 1].Less(a[i]) {
			return false
		}
	}
	return true
}

type sortable func ([]Comparable)
func test_sort(t *testing.T, s sortable, name string, len, max int, isshow bool) {
	a := input(len, max)
	if isshow {
		fmt.Printf("before %s: \n", name)
		show(a)
		defer func() {
			fmt.Printf("after %s: \n", name)
			show(a)
		}()
	}
	s(a)
	if !is_sorted(a) {
		t.Fatalf("%s error!\n", name)
	}
}