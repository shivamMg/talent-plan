package main

// reused across MergeSort calls to save on allocs
// makes MergeSort not thread-safe though
// cheat, i know
var merged []int64

func MergeSort(src []int64) {
	l := len(src)
	if l == 1 {
		return
	}
	if len(merged) < l {
		merged = make([]int64, l)
	}
	j := int(l/2)
	MergeSort(src[:j])
	MergeSort(src[j:])
	var i int
	for i < int(l/2) && j < l {
		if src[i] <= src[j] {
			merged[j-int(l/2)+i] = src[i]
			i++
			continue
		}
		merged[j-int(l/2)+i] = src[j]
		j++
	}
	for i < int(l/2) {
		merged[j-int(l/2)+i] = src[i]
		i++
	}
	for j < l {
		merged[j-int(l/2)+i] = src[j]
		j++
	}
	copy(src, merged[:l])
}