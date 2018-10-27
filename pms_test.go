package pms

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{2, 1, 7, 3, 9, 6, 0, 8, 4, 5}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mergesort(s)
	if !reflect.DeepEqual(s, expected) {
		t.Error("did not match expected sorting")
	}
}
