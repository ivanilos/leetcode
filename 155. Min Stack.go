type MinStack struct {
    elems [][]int
}


func Constructor() MinStack {
    return MinStack{
        [][]int{},
    }
}


func (this *MinStack) Push(val int)  {
    mini := val
    if len(this.elems) > 0 {
        mini = min(mini, this.elems[len(this.elems) - 1][1])
    }

    this.elems = append(this.elems, []int{val, mini})
}


func (this *MinStack) Pop()  {
    this.elems = this.elems[:len(this.elems) - 1]
}


func (this *MinStack) Top() int {
    return this.elems[len(this.elems) - 1][0]
}


func (this *MinStack) GetMin() int {
    return this.elems[len(this.elems) - 1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
