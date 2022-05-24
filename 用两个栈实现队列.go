package main

type CQueue1 struct {
	s1, s2 []int
}

func main() {
	Constructor1()
	AppendTail(3)
}

func Constructor1() CQueue1 {
	return CQueue1{
		s1: []int{},
		s2: []int{},
	}
}

func (this *CQueue1) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue1) DeleteHead() int {
	v := 0
	if len(this.s2) != 0 {
		v = this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
	} else {
		if len(this.s1) == 0 {
			v = -1
		} else {
			for len(this.s1) != 0 {
				this.s2 = append(this.s2, this.s1[len(this.s1)-1])
				this.s1 = this.s1[:len(this.s1)-1]
			}
			v = this.s2[len(this.s2)-1]
			this.s2 = this.s2[:len(this.s2)-1]
		}
	}
	return v
}
