package internal

type stack struct {
	head   *node
	length int
}

type node struct {
	position *position
	next     *node
}

func NewStack() *stack {
	return &stack{}
}

func (q *stack) push(position *position) {
	n := &node{
		position: position,
		next:     q.head,
	}

	q.head = n
	q.length++
}

func (q *stack) pop() *position {
	if q.length == 0 {
		return nil
	}

	candidate := q.head
	q.head = candidate.next
	q.length--

	return candidate.position
}

func (q *stack) isEmpty() bool {
	return q.length == 0
}
