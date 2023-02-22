package elevator

type TreeSet struct {
	tree   []Request
	length int
}

type TreeSets interface {
	Push()
	Pop()
}

func NewTreeSet() TreeSet {
	return TreeSet{}
}

func (t *TreeSet) Push(r Request) {
	// TO DO: Check if already exist to remove uplication
	t.tree = append(t.tree, r)
	t.length++
}

func (t *TreeSet) Pop() {
	if len(t.tree) > 1 {
		t.tree = t.tree[1:]
		t.length--
	}
}
