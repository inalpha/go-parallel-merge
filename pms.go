package pms

func mergesort(s []int) {
	if len(s) > 1 {
		m := len(s) / 2
		mergesort(s[:m])
		mergesort(s[m:])
		merge(s, m)
	}
}

func merge(s []int, m int) {
	c := make([]int, len(s))
	copy(c, s)
	cl := 0
	cr := m
	i := 0
	h := len(s) - 1

	for cl <= m-1 && cr <= h {
		if c[cl] <= c[cr] {
			s[i] = c[cl]
			cl++
		} else {
			s[i] = c[cr]
			cr++
		}
		i++
	}

	for cl <= m-1 {
		s[i] = c[cl]
		cl++
		i++
	}
}
