package data_structures

type Node struct {
	l, r int
	val  int64
}

type SGT struct {
	n    int
	tree []Node
}

func (s *SGT) build(p, l, r int, arr []int) {
	s.tree[p].l = l
	s.tree[p].r = r
	if l == r {
		s.tree[p].val = int64(arr[l])
		return
	}
	mid, ls, rs := (l+r)/2, p*2, p*2+1
	s.build(ls, l, mid, arr)
	s.build(rs, mid+1, r, arr)
	s.tree[p].val = s.tree[ls].val + s.tree[rs].val
}
func (s *SGT) Build(arr []int, n int) {
	s.tree = make([]Node, n*4+1)
	s.build(1, 1, n, arr)
}

func (s *SGT) update(p, x, v int) {
	if s.tree[p].l == s.tree[p].r {
		s.tree[p].val += int64(v)
		return
	}
	mid, ls, rs := (s.tree[p].l+s.tree[p].r)/2, p*2, p*2+1
	if x <= mid {
		s.update(ls, x, v)
	} else {
		s.update(rs, x, v)
	}
	s.tree[p].val = s.tree[ls].val + s.tree[rs].val
}
func (s *SGT) Update(x, v int) {
	s.update(1, x, v)
}

func (s *SGT) query(p, l, r int) (ans int64) {
	if l <= s.tree[p].l && s.tree[p].r <= r {
		ans = s.tree[p].val
		return
	}
	mid, ls, rs := (s.tree[p].l+s.tree[p].r)/2, p*2, p*2+1
	if l <= mid {
		ans += s.query(ls, l, r)
	}
	if r > mid {
		ans += s.query(rs, l, r)
	}
	return
}
func (s *SGT) Query(l, r int) (ans int64) {
	ans = s.query(1, l, r)
	return
}
