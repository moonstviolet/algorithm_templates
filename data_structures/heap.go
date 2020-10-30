package data_structures

import "sort"

type Heap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func down(h Heap, root, n int) {
	for {
		son := root*2 + 1
		if son >= n || son < 0 {
			break
		}
		if son < n-1 && h.Less(son+1, son) {
			son++
		}
		if h.Less(son, root) {
			h.Swap(son, root)
			root = son
		} else {
			break
		}
	}
}
func up(h Heap, idx int) {
	for idx != 0 {
		pre := (idx - 1) / 2
		if h.Less(idx, pre) {
			h.Swap(idx, pre)
			idx = pre
		} else {
			break
		}
	}
}
func Init(h Heap) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}
func Push(h Heap, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}
func Pop(h Heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
