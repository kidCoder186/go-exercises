package binarysearchtree

// SearchTreeData structure
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns a SearchTreeData Object.
func Bst(key int) *SearchTreeData {
	return &SearchTreeData{nil, key, nil}
}

// Insert inserts a element into BST.
func (bst *SearchTreeData) Insert(x int) {
	node := &SearchTreeData{nil, x, nil}
	if bst.data >= x {
		if bst.left == nil {
			bst.left = node
		} else {
			bst.left.Insert(x)
		}
	} else {
		if bst.right == nil {
			bst.right = node
		} else {
			bst.right.Insert(x)
		}
	}
}

// MapString returns all values of bst that are mapped to string.
func (bst *SearchTreeData) MapString(f func(int) string) []string {
	if bst == nil {
		return []string{}
	}
	res := []string{}
	res = append(res, bst.left.MapString(f)...)
	res = append(res, f(bst.data))
	res = append(res, bst.right.MapString(f)...)
	return res
}

// MapInt returns all values of bst that are mapped to int.
func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	if bst == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, bst.left.MapInt(f)...)
	res = append(res, f(bst.data))
	res = append(res, bst.right.MapInt(f)...)
	return res
}
