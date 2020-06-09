package BST

var (
	insertCases = []struct {
		description string
		element     Structure
		tree        BST
		err         bool
	}{
		{
			description: "insert into empty tree",
			element:     Structure{Index: 5},
			tree:        BST{},
			err:         false,
		},
		{
			description: "insert repeated element",
			element:     Structure{Index: 5},
			tree: BST{
				root: &Node{
					Apex:  Structure{Index: 5},
					left:  nil,
					right: nil,
				},
			},
			err: true,
		},
		{
			description: "insert smaller element",
			element:     Structure{Index: 4},
			tree: BST{
				root: &Node{
					Apex:  Structure{Index: 5},
					left:  nil,
					right: nil,
				},
			},
			err: false,
		},
		{
			description: "insert bigger element",
			element:     Structure{Index: 6},
			tree: BST{
				root: &Node{
					Apex:  Structure{Index: 5},
					left:  nil,
					right: nil,
				},
			},
			err: false,
		},
		{
			description: "insert bigger element",
			element:     Structure{Index: 7},
			tree: BST{
				root: &Node{
					Apex: Structure{Index: 5},
					right: &Node{
						Apex:  Structure{Index: 6},
						left:  nil,
						right: nil,
					},
					left: nil,
				},
			},
			err: false,
		},
	}

	getCases = []struct {
		description string
		element     Structure
		tree        BST
		err         bool
	}{
		{
			description: "Empty tree",
			element: Structure{Index: 4},
			tree: BST{},
			err: true,
		},
		{
			description: "Find element that does not exist",
			element: Structure{Index: 4},
			tree: BST{
				root: &Node{
					Apex: Structure{Index: 5},
					right: &Node{
						Apex:  Structure{Index: 6},
						left:  nil,
						right: nil,
					},
					left: nil,
				},
			},
			err: true,
		},
		{
			description: "Find root element",
			element: Structure{Index: 5},
			tree: BST{
				root: &Node{
					Apex: Structure{Index: 5},
					right: &Node{
						Apex:  Structure{Index: 6},
						left:  nil,
						right: nil,
					},
					left: nil,
				},
			},
			err: false,
		},
		{
			description: "Find bigger element",
			element: Structure{Index: 6},
			tree: BST{
				root: &Node{
					Apex: Structure{Index: 5},
					right: &Node{
						Apex:  Structure{Index: 6},
						left:  nil,
						right: nil,
					},
					left: &Node{
						Apex:  Structure{Index: 4},
						left:  nil,
						right: nil,
					},
				},
			},
			err: false,
		},
		{
			description: "Find smaller element",
			element: Structure{Index: 4},
			tree: BST{
				root: &Node{
					Apex: Structure{Index: 5},
					right: &Node{
						Apex:  Structure{Index: 6},
						left:  nil,
						right: nil,
					},
					left: &Node{
						Apex:  Structure{Index: 4},
						left:  nil,
						right: nil,
					},
				},
			},
			err: false,
		},
	}
)
