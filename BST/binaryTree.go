// Package BST implements a Binary Search Tree for a Structure,
// you may want to change the Structure type
package BST

import (
	"errors"
	"strconv"
)

type Structure struct {
	Index int
}

type Node struct {
	Apex  Structure
	left  *Node
	right *Node
}

type BST struct {
	root   *Node
	Cursor *Node
	aux    *Node
}

func (bst *BST) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BST) Find(element int) bool {
	bst.Cursor = bst.root
	for bst.Cursor != nil && bst.Cursor.Apex.Index != element {
		bst.aux = bst.Cursor
		if bst.Cursor.Apex.Index > element {
			bst.Cursor = bst.Cursor.left
		} else {
			bst.Cursor = bst.Cursor.right
		}
	}
	if bst.Cursor != nil {
		return true
	}
	return false
}

func (bst *BST) Insert(element Structure) error {
	temp := &Node{
		Apex:  element,
		left:  nil,
		right: nil,
	}
	if bst.Find(element.Index) {
		return errors.New("already exists")
	}
	if bst.IsEmpty() {
		bst.root = temp
		bst.Cursor = bst.root
		bst.aux = bst.root
	} else {
		bst.Cursor = temp
		if bst.aux.Apex.Index > temp.Apex.Index {
			bst.aux.left = bst.Cursor
		} else {
			bst.aux.right = bst.Cursor
		}
	}
	return nil
}

func (bst BST) Show() (view string) {
	if bst.root == nil {
		return
	}
	bst.Cursor = bst.root
	bst.aux = bst.root
	view += "(" + strconv.Itoa(bst.Cursor.Apex.Index) + ")\n"
	view += "Left son\n"
	if bst.Cursor.left == nil {
		view += "<nil>\n"
	} else {
		view += "{" + strconv.Itoa(bst.Cursor.left.Apex.Index) + "}\n"
	}
	view += "Right son\n"
	if bst.Cursor.right == nil {
		view += "<nil>\n"
	} else {
		view += "{" + strconv.Itoa(bst.Cursor.right.Apex.Index) + "}\n"
	}
	bst.root = bst.Cursor.left
	view += bst.Show()
	bst.root = bst.Cursor.right
	view += bst.Show()
	return
}

func (bst BST) GetElement(index int) (element Structure, err error){
	if !bst.Find(index){
		return element, errors.New("element not found")
	}
	return bst.Cursor.Apex, nil
}