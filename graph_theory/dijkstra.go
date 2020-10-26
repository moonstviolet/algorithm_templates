package graph_theory

import "container/heap"

type Node struct {
	idx int
	dis int64
}
type Heap []Node

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Heap) empty() bool {
	return len(h) == 0
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

const (
	N   = 1e5 + 10
	M   = 2e5 + 10
	INF = 0x3f3f3f3f3f3f3f3f
)

type Edge struct {
	to, w, next int
}

var (
	cnt  int
	head [N]int
	Dis  [N]int64
	vis  [N]bool
	edge [M]Edge
)

func Add(u, v, w int) {
	cnt++
	edge[cnt].next = head[u]
	edge[cnt].to = v
	edge[cnt].w = w
	head[u] = cnt
}

func Dijkstra(n, st int) {
	for i := 1; i <= n; i++ {
		Dis[i] = INF
	}
	Dis[st] = 0
	var pq = &Heap{}
	heap.Init(pq)
	heap.Push(pq, Node{st, 0})
	for !pq.empty() {
		t := heap.Pop(pq).(Node)
		u := t.idx
		if vis[u] {
			continue
		}
		vis[u] = true
		for i := head[u]; i > 0; i = edge[i].next {
			v, w := edge[i].to, edge[i].w
			if Dis[u]+int64(w) < Dis[v] {
				Dis[v] = Dis[u] + int64(w)
				heap.Push(pq, Node{v, Dis[v]})
			}
		}
	}
}
