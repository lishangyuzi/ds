package test

import (
	"testing"
	"errors"
	"ds/sort"
)

var arr []int = []int{1, 2, 3, 3, 5, 0}
var sortedArr []int = []int{0, 1, 2, 3, 3, 5}

func TestBubbleSort(t *testing.T) {
	bubbleSource := arr
	sort.QuickSort(bubbleSource, 0, len(bubbleSource)-1)

	err := checkSeq(bubbleSource, sortedArr)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestQuickSort(t *testing.T) {
	quickSource := arr
	sort.BubbleSort(quickSource)

	err := checkSeq(quickSource, sortedArr)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func checkSeq(a []int, b []int) error {
	for idx := 0; len(a) > idx+1 && len(a) > idx+1; idx++ {
		if a[idx] != b[idx] {
			return errors.New("seq not equal")
		}
	}
	return nil
}
