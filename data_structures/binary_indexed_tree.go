package data_structures

type BIT struct {
	len  int
	arr  []int
	data []int
}

func (b *BIT) Lowbit(x int) int {
	return x & -x
}
func (b *BIT) Init() {
	for i := 1; i <= b.len; i++ {
		b.arr[i] = b.arr[i] + b.arr[i-1]
		b.data[i] = b.arr[i] - b.arr[i-b.lowbit(i)]
	}
}
func (b *BIT) Add(x, v int) {
	for x <= b.len {
		b.data[x] += v
		x += b.lowbit(x)
	}

}
func (b *BIT) Query(x int) (sum int) {
	for x > 0 {
		sum += b.data[x]
		x -= b.lowbit(x)
	}
	return
}
