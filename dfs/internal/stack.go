package internal

type Stack struct {
	head   *node
	length int
}

type node struct {
	position *position
	next     *node
}

func NewStack() *Stack {
	return &Stack{}
}

func (q *Stack) Push(position *position) {
	n := &node{
		position: position,
		next:     q.head,
	}

	q.head = n
	q.length++
}

func (q *Stack) Pop() *position {
	if q.length == 0 {
		return nil
	}

	candidate := q.head
	q.head = candidate.next
	q.length--

	return candidate.position
}

func (q *Stack) IsEmpty() bool {
	return q.length == 0
}
