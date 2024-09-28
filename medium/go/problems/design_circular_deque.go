package problems

type MyCircularDeque struct {
	f     int
	l     int
	queue []int
	k     int
}

func Constructor(k int) MyCircularDeque {
	queue := make([]int, k)
	for i := range queue {
		queue[i] = -1
	}
	return MyCircularDeque{f: 0, l: 0, queue: queue, k: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.queue[this.f] != -1 {
		return false
	}
	this.queue[this.f] = value
	this.f = moveBack(this.f, this.k)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.queue[this.l] != -1 {
		return false
	}
	this.queue[this.l] = value
	this.l = moveFront(this.l, this.k)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	p := moveFront(this.f, this.k)
	if this.queue[p] == -1 {
		return false
	}
	this.queue[p] = -1
	this.f = p
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	pointer := moveBack(this.l, this.k)
	if this.queue[pointer] == -1 {
		return false
	}
	this.queue[pointer] = -1
	this.l = pointer
	return true
}

func (this *MyCircularDeque) GetFront() int {
	return this.queue[moveFront(this.f, this.k)]
}

func (this *MyCircularDeque) GetRear() int {
	return this.queue[moveBack(this.l, this.k)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.queue[moveBack(this.l, this.k)] == -1
}

func (this *MyCircularDeque) IsFull() bool {
	if this.f == -1 {
		return false
	}
	return this.queue[this.f] == this.queue[this.l]
}

func moveFront(p, k int) int {
	return p + 1%k
}

func moveBack(p, k int) int {
	val := p - 1
	if val < 0 {
		val = k - 1
	}
	return val
}
