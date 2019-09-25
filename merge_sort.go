package ysort

type MergeSort struct {
	aux []Comparable
}

func (s *MergeSort)TopDown(arr []Comparable) {
	s.aux = make([]Comparable, len(arr))
	s.topdownMerge(arr, 0, len(arr) - 1)
}

func (s *MergeSort)BottomUp(arr []Comparable) {
	s.aux = make([]Comparable, len(arr))
}

func (s *MergeSort)topdownMerge(arr []Comparable, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi - lo) / 2 + 1
	s.topdownMerge(arr, lo, mid)
	s.topdownMerge(arr, mid + 1, hi)
	s.merge(arr, mid, lo, hi)
}

func (s *MergeSort)merge(arr []Comparable, mid, lo, hi int) {
	for i := lo; i <= hi; i++ {
		s.aux[i] = arr[i]
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			// [lo ... mid] is out
			arr[k] = s.aux[j]
			i++
		}  else if j > hi {
			// [mid + 1 ... hi] is out
			arr[k] = s.aux[i]
			j++
		} else if s.aux[i].Less(s.aux[j]) {
			arr[k] = s.aux[i]
			i++
		} else {
			arr[k] = s.aux[j]
			j++
		}
	}
}