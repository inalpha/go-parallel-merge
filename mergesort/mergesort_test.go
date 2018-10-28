package mergesort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

const size = 9999

func TestSequential(t *testing.T) {
	s := []int{2, 1, 7, 3, 9, 6, 0, 8, 4, 5}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Sequential(s)
	if !reflect.DeepEqual(s, expected) {
		t.Error("did not match expected sorting")
	}
}
func TestParallel(t *testing.T) {
	s := random(size)
	// t.Logf("Source: %v", s)

	expected := make([]int, size)
	copy(expected, s)
	Sequential(expected)
	// t.Logf("Expected: %v", expected)
	Parallel(s)
	// t.Logf("Result: %v", s)
	if !reflect.DeepEqual(s, expected) {
		t.Error("did not match expected sorting")
	}
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		Sequential(s)
		b.StopTimer()
	}
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		Parallel(s)
		b.StopTimer()
	}
}

func random(n int) []int {
	s := make([]int, n)
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}
	return s
}
