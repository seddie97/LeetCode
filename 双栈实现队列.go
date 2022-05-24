package main

type CQueue struct {
	s1, s2 []int
}

func Constructor2() CQueue {
	return CQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	v := 0
	if this.s2 != nil {
		v = this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
	} else {
		if this.s1 == nil {
			v = -1
		} else {
			for len(this.s1) > 0 {
				this.s2 = append(this.s2, this.s1[len(this.s1)-1])
				this.s1 = this.s1[:len(this.s1)-1]
			}
			v = this.s2[len(this.s2)-1]
			this.s2 = this.s2[:len(this.s2)-1]
		}
	}

	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
