type RecentCounter struct {
    maxMemory int
    requestTimes []int
}


func Constructor() RecentCounter {
    return RecentCounter {
        3000,
        []int{},
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.requestTimes = append(this.requestTimes, t)

    for this.requestTimes[0] < t - this.maxMemory {
        this.requestTimes = this.requestTimes[1 : len(this.requestTimes)]
    }
    return len(this.requestTimes)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
