package stack

type Stack struct {
	top    *node
	length int
}

type node struct {
	val  interface{}
	next *node
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Push(val interface{}) {
	n := &node{
		val:  val,
		next: s.top,
	}
	s.top = n
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	n := s.top
	s.top = n.next
	s.length--
	return n.val
}
