package main

func main() {
	name := "hamed"
	namePtr := &name
	namePtr2Ptr := &namePtr

}

func search1() {
	t1 := &Tree[int]{
		cmp: func(a int, b int) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			} else {
				return 0
			}

		},
		root: &leaf[int]{val: 12},
	}

	for i := 0; i < 1000; i++ {
		root := t1.root
		var left = root.left
		var right = root.right
		left.val = i
	}
}

func search2() {

}

func (bt *Tree[T]) find1(val T) *leaf[T] {
	pl := bt.root
	for pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = (*pl).left
		case cmp > 0:
			pl = (*pl).right
		default:
			return pl
		}
	}
	return pl
}

func (bt *Tree[T]) find(val T) **leaf[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *leaf[T]
}

type leaf[T any] struct {
	val         T
	left, right *leaf[T]
}
