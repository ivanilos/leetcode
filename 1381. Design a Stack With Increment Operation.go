type CustomStack struct {
    v []int
    top int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack {
        make([]int, maxSize),
        -1,
    }
}


func (this *CustomStack) Push(x int)  {
    if !this.isFull() {
        this.top++
        this.v[this.top] = x
    }
}


func (this *CustomStack) Pop() int {
    if this.isEmpty() {
        return -1
    }

    result := this.v[this.top]
    this.top--
    return result
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; i < k && i <= this.top; i++ {
        this.v[i] += val
    }
}

func (this *CustomStack) isEmpty() bool {
    return this.top == -1
}

func (this *CustomStack) isFull() bool {
    return this.top == len(this.v) - 1
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
