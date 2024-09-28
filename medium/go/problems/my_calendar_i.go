package problems

type Tree struct {
	l *Tree
	r *Tree
	s int
	e int
}

func (root *Tree) insert(start, end int) bool {
	c := root
	for true {
		if end <= c.s {
			if c.l == nil {
				c.l = &Tree{s: start, e: end}
				return true
			}
			c = c.l
		} else if start >= c.e {
			if c.r == nil {
				c.r = &Tree{s: start, e: end}
				return true
			}
			c = c.r
		} else {
			return false
		}
	}
	return false
}

type MyCalendar struct {
	root *Tree
}

func Constructor() MyCalendar {
	return MyCalendar{root: nil}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.root == nil {
		this.root = &Tree{s: start, e: end}
		return true
	}
	return this.root.insert(start, end)
}
