package linkedList

// https://golangbyexample.com/singly-linked-list-in-golang/

type element struct {
	value string
	next  *element
}

type singleList struct {
	len  int
	head *element
}

func newSingleList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(value string) {
	ele := &element{value: value}

	if s.head == nil {
		s.head = ele
	} else {
		ele.next = s.head
		s.head = ele
	}

	s.len++
}

func (s *singleList) AddBack(value string) {
	ele := &element{value: value}

	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}

	s.len++
}
