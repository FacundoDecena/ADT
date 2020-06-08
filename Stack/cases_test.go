package Stack

var (
	pushTestCases = []struct {
		description string
		stack       Stack
		item        interface{}
		itemWanted  interface{}
		stackWanted Stack
	}{
		{
			description: "Push to empty stack",
			stack:       Stack{top: -1},
			item:        "New Item",
			stackWanted: Stack{top: 0, tail: []interface{}{"New Item"}},
		},
		{
			description: "Push to stack string",
			stack:       Stack{top: 0, tail: []interface{}{"Prev Item"}},
			item:        "New Item",
			stackWanted: Stack{top: 1, tail: []interface{}{"Prev Item", "New Item"}},
		},
		{
			description: "Push to stack int",
			stack:       Stack{top: 0, tail: []interface{}{7}},
			item:        8,
			stackWanted: Stack{top: 1, tail: []interface{}{7, 8}},
		},
		{
			description: "Push to stack different types",
			stack:       Stack{top: 0, tail: []interface{}{7}},
			item:        "seven",
			stackWanted: Stack{top: 1, tail: []interface{}{7, "seven"}},
		},
	}

	peekTestCases = []struct {
		description string
		stack       Stack
		stackWanted Stack
		itemWanted  interface{}
		err         bool
	}{
		{
			description: "Peek empty stack",
			stack:       Stack{-1, []interface{}{}},
			stackWanted: Stack{-1, []interface{}{}},
			itemWanted:  nil,
			err:         true,
		},
		{
			description: "Peek stack with one element",
			stack:       Stack{0, []interface{}{"Item"}},
			stackWanted: Stack{0, []interface{}{"Item"}},
			itemWanted:  "Item",
			err:         false,
		},
		{
			description: "Peek stack with more elements",
			stack:       Stack{2, []interface{}{"Item", 5, "Other item"}},
			stackWanted: Stack{2, []interface{}{"Item", 5, "Other item"}},
			itemWanted:  "Other item",
			err:         false,
		},
	}
	popTestCases = []struct {
		description string
		stack       Stack
		stackWanted Stack
		itemWanted  interface{}
		err         bool
	}{
		{
			description: "Pop empty stack",
			stack:       Stack{-1, []interface{}(nil)},
			stackWanted: Stack{-1, []interface{}(nil)},
			itemWanted:  nil,
			err:         true,
		},
		{
			description: "Pop stack with one element",
			stack:       Stack{0, []interface{}{"Item"}},
			stackWanted: Stack{-1, []interface{}(nil)},
			itemWanted:  "Item",
			err:         false,
		},
		{
			description: "Pop stack with more elements",
			stack:       Stack{2, []interface{}{"Item", 5, "Other item"}},
			stackWanted: Stack{1, []interface{}{"Item", 5}},
			itemWanted:  "Other item",
			err:         false,
		},
	}
)
