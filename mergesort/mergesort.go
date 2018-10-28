package mergesort

func Sequential(s []int) {
	if len(s) > 1 {
		m := len(s) / 2
		Sequential(s[:m])
		Sequential(s[m:])
		merge(s, m)
	}
}

func Parallel(s []int) {

}

func merge(s []int, m int) {
	cl, cr, i, h := 0, m, 0, len(s)-1
	c := make([]int, len(s))
	copy(c, s)

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
