package main

import "inalpha/sort/mergesort"

func main() {
	unsorted := []int{2, 1, 7, 3, 9, 6, 0, 8, 4, 5}
	mergesort.Sequential(unsorted)
}
