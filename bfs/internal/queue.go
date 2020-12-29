package internal

type Queue struct {
	tail   *node
	head   *node
	length int
}

type node struct {
	position *position
	next     *node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(position *position) {
	n := &node{position: position}

	switch q.length {
	case 0:
		q.head = n
	case 1:
		q.head.next = n
	default:
		q.tail.next = n
	}

	q.tail = n
	q.length++
}

func (q *Queue) Pop() *position {
	if q.length == 0 {
		return nil
	}

	candidate := q.head
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	q.length--
	return candidate.position
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}
