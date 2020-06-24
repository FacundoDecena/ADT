// Package BST implements a Binary Search Tree for a Structure
package BST

import (
	"errors"
	"strconv"
)

type Structure struct {
	Index int
	Element interface{}
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

// IsEmpty checks if the Tree is empty
func (bst *BST) IsEmpty() bool {
	return bst.root == nil
}

// Find search for an elements and return true if exists
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

//Insert insert a Structure in the Tree
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

//Show returns a string with basic information of the tree
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

//GetElement returns an element with the given index <or id>
func (bst BST) GetElement(index int) (element Structure, err error) {
	if !bst.Find(index) {
		return element, errors.New("element not found")
	}
	return bst.Cursor.Apex, nil
}

//DeleteElement deletes the element with the given index or return an error if it does not exist
func (bst *BST) DeleteElement(index int) error {
	if !bst.Find(index) {
		return errors.New("element not found")
	}
	if bst.Cursor.left == nil && bst.Cursor.right == nil {
		bst.deleteNoSons()
	} else if bst.Cursor.left == nil || bst.Cursor.right == nil {
		bst.deleteOneSon()
	} else {
		bst.deleteTwoSons()
	}
	return nil
}

func (bst *BST) deleteNoSons() {
	if bst.aux.right == bst.Cursor {
		bst.aux.right = nil
		bst.Cursor = nil
	} else if bst.aux.left == bst.Cursor {
		bst.aux.left = nil
		bst.Cursor = nil
	} else {
		bst.Cursor = nil
		bst.root = nil
		bst.aux = nil
	}
}

func (bst *BST) deleteOneSon() {
	if bst.Cursor.right != nil {
		if bst.aux.right == bst.Cursor {
			bst.aux.right = bst.Cursor.right
		} else if bst.aux.left == bst.Cursor {
			bst.aux.left = bst.Cursor.right
		} else {
			bst.root = bst.Cursor.right
		}
	} else {
		if bst.aux.right == bst.Cursor {
			bst.aux.right = bst.Cursor.left
		} else if bst.aux.left == bst.Cursor {
			bst.aux.left = bst.Cursor.left
		} else {
			bst.root = bst.Cursor.left
		}
	}
	bst.Cursor = nil
}

func (bst *BST) deleteTwoSons() {
	deleteCursor := bst.Cursor
	bst.aux = bst.Cursor
	bst.aux = bst.Cursor.left
	for bst.Cursor.right != nil {
		bst.aux = bst.Cursor
		bst.Cursor = bst.Cursor.right
	}
	deleteCursor.Apex = bst.Cursor.Apex

	if bst.Cursor.left == nil {
		if bst.aux.right == bst.Cursor {
			bst.aux.right = nil
		} else {
			bst.aux.left = nil
		}
	} else {
		if bst.aux.right == bst.Cursor {
			bst.aux.right = bst.Cursor.left
		} else {
			bst.aux.left = bst.Cursor.left
		}
	}
	bst.Cursor = nil
}
