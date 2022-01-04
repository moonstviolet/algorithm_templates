package data_structures

import "errors"

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

func (q *Queue) Pop() (ret interface{}, err error) {
	if q.count <= 0 {
		err = errors.New("queue: Pop() called on empty queue")
		return
	}
	ret = q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) & (len(q.data) - 1)
	q.count--
	if len(q.data) > minSize && q.count<<2 == len(q.data) {
		q.resize()
	}
	return
}

func (q *Queue) Push(elem interface{}) {
	if q.count == len(q.data) {
		q.resize()
	}
	q.data[q.tail] = elem
	q.tail = (q.tail + 1) & (len(q.data) - 1)
	q.count++
}

func (q *Queue) Front() (ret interface{}, err error) {
	if q.count <= 0 {
		err = errors.New("queue: Front() called on empty queue")
		return
	}
	ret = q.data[q.head]
	return
}

func (q *Queue) Get(idx int) (ret interface{}, err error) {
	if idx < 0 {
		idx += q.count
	}
	if idx < 0 || idx > q.count {
		err = errors.New("queue: Get() called with index out of range")
		return
	}
	ret = q.data[(q.head+idx)&(len(q.data)-1)]
	return
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
