package data_structures

const minSize = 16

type Queue struct {
	data              []interface{}
	head, tail, count int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, minSize),
	}
}

func (q *Queue) Pop() interface{} {
	if q.count <= 0 {
		panic("queue: Pop() called on empty queue")
	}
	ret := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) & (len(q.data) - 1)
	q.count--
	if len(q.data) > minSize && q.count<<2 == len(q.data) {
		q.resize()
	}
	return ret
}

func (q *Queue) Push(elem interface{}) {
	if q.count == len(q.data) {
		q.resize()
	}
	q.data[q.tail] = elem
	q.tail = (q.tail + 1) & (len(q.data) - 1)
	q.count++
}

func (q *Queue) Front() interface{} {
	if q.count <= 0 {
		panic("queue: Front() called on empty queue")
	}
	return q.data[q.head]
}

func (q *Queue) Get(idx int) interface{} {
	if idx < 0 {
		idx += q.count
	}
	if idx < 0 || idx > q.count {
		panic("queue: Get() called with index out of range")
	}
	return q.data[(q.head+idx)&(len(q.data)-1)]
}

func (q *Queue) Size() int {
	return q.count
}

func (q *Queue) Empty() bool {
	return q.count == 0
}

func (q *Queue) Clear() {
	q.head = 0
	q.tail = 0
	q.count = 0
	q.data = make([]interface{}, minSize)
}

func (q *Queue) resize() {
	newData := make([]interface{}, q.count<<1)
	if q.tail > q.head {
		copy(newData, q.data[q.head:q.tail])
	} else {
		n := copy(newData, q.data[q.head:])
		copy(newData[n:], q.data[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.data = newData
}
