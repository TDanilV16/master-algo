package internal

type SegmentTree interface {
	Sum(l, r int) int
	Update(index, newVal int)
}

type SegmentTreeSum struct {
	t []int
	n int
}

func NewSegmentTreeSum(a []int) *SegmentTreeSum {
	n := len(a)
	t := make([]int, 4*n)

	build(a, &t, 1, 0, n-1)

	return &SegmentTreeSum{t, n}
}

func (s *SegmentTreeSum) Sum(l, r int) int {
	return s.sum(1, 0, s.n-1, l, r)
}

func (s *SegmentTreeSum) Update(index, newVal int) {
	s.update(1, 0, s.n-1, index, newVal)
}

func build(a []int, t *[]int, v, tl, tr int) {
	if tl == tr {
		(*t)[v] = a[tl]
	} else {
		tm := (tl + tr) / 2
		build(a, t, v*2, tl, tm)
		build(a, t, v*2+1, tm+1, tr)
		(*t)[v] = (*t)[v*2] + (*t)[v*2+1]
	}
}

func (s *SegmentTreeSum) sum(v, tl, tr, l, r int) int {
	if l > r {
		return 0
	}

	if l == tl && r == tr {
		return s.t[v]
	}

	tm := (tl + tr) / 2

	return s.sum(v*2, tl, tm, l, min(r, tm)) + s.sum(v*2+1, tm+1, tr, max(l, tm+1), r)
}

func (s *SegmentTreeSum) update(v, tl, tr, pos, val int) {
	if tl == tr {
		s.t[v] = val
	} else {
		tm := (tl + tr) / 2
		if pos <= tm {
			s.update(v*2, tl, tm, pos, val)
		} else {
			s.update(v*2+1, tm+1, tr, pos, val)
		}

		s.t[v] = s.t[v*2] + s.t[v*2+1]
	}
}
