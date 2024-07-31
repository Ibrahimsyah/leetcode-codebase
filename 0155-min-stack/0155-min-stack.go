type MinStack struct {
    stk []int
    min []int
}


func Constructor() MinStack {
    return MinStack{}    
}


func (this *MinStack) Push(val int)  {
    this.stk = append(this.stk, val)
    
    if len(this.min) == 0 || val <= this.min[len(this.min) - 1] {
        this.min = append(this.min, val)
    } else {
        this.min = append(this.min, this.min[len(this.min) - 1])
    }
}


func (this *MinStack) Pop()  {
    this.stk = this.stk[:len(this.stk) - 1]
    this.min = this.min[:len(this.min) - 1]
}


func (this *MinStack) Top() int {
    return this.stk[len(this.stk) - 1]
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */