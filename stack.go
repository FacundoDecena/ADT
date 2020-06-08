// Package ADT implements abstract data types
package main

import "errors"

//Stack
type Stack struct {
	top   int
	array []interface{}
}

func (Stack) Init() Stack {
	var arr []interface{}
	return Stack{-1, arr}
}

func (s Stack) IsEmpty() bool {
	return s.top == -1
}

func (s Stack) Push(item interface{}) {
	s.top++
	s.array[s.top] = item
}

func (s Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	s.top--
	return s.array[s.top], nil
}

func (s Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	return s.array[s.top], nil
}
