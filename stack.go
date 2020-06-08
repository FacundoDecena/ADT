// Package ADT implements abstract data types
package ADT

import "errors"

//Stack
type Stack struct {
	top  int
	tail []interface{}
}

func (s *Stack) Init() {
	var arr []interface{}
	s.top = -1
	s.tail = arr
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(item interface{}) {
	s.top++
	s.tail = append(s.tail, item)
}

func (s *Stack) Pop() (item interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	item = s.tail[s.top]
	s.top--
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	return s.tail[s.top], nil
}
