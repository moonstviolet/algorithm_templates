package data_structures

type BIT struct {
	n    int
	pre  []int64
	data []int64
}

func lowbit(x int) int {
	return x & -x
}
func (b *BIT) Build(n int) {
	b.n = n
	b.data = make([]int64, b.n+1)
}
func (b *BIT) Init(arr []int) {
	b.pre = make([]int64, b.n+1)
	for i := 1; i <= b.n; i++ {
		b.pre[i] = b.pre[i-1] + int64(arr[i])
		b.data[i] = b.pre[i] - b.pre[i-lowbit(i)]
	}
}
func (b *BIT) Add(x, v int) {
	for x <= b.n {
		b.data[x] += int64(v)
		x += lowbit(x)
	}

}
func (b *BIT) Query(x int) (sum int64) {
	for x > 0 {
		sum += b.data[x]
		x -= lowbit(x)
	}
	return
}
