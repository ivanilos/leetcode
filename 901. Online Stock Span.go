type StockSpanner struct {
    st [][]int
    day int
}


func Constructor() StockSpanner {
    return StockSpanner {
        [][]int{[]int{1000000000, -1}},
        0,
    }
}


func (this *StockSpanner) Next(price int) int {
    for this.st[len(this.st) - 1][0] <= price {
        this.st = this.st[: len(this.st) - 1]
    }
    result := this.day - this.st[len(this.st) - 1][1]

    this.st = append(this.st, []int{price, this.day})
    this.day++

    return result
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
