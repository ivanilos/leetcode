type MyCircularDeque struct {
    front int
    back int
    sz int
    cap int
    dq []int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        0,
        0,
        0,
        k,
        make([]int, k),
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.sz >= this.cap {
        return false
    }
    if this.sz > 0 {
        this.front = this.prev(this.front)
    }
    this.dq[this.front] = value
    this.sz++

    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.sz >= this.cap {
        return false
    }
    if this.sz > 0 {
        this.back = this.next(this.back)
    }
    this.dq[this.back] = value
    this.sz++;

    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.sz == 0 {
        return false
    } else if this.sz > 1 {
        this.front = this.next(this.front)
    }
    this.sz--
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.sz == 0 {
        return false
    } else if this.sz > 1 {
        this.back = this.prev(this.back)
    }
    this.sz--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    fmt.Println(this.front)
    if this.sz == 0 {
        return -1
    }
    return this.dq[this.front]
}


func (this *MyCircularDeque) GetRear() int {
    if this.sz == 0 {
        return -1
    }
    return this.dq[this.back]
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.sz == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.sz == this.cap
}

func (this *MyCircularDeque) prev(idx int) int {
    return (idx - 1 + this.cap) % this.cap
}

func (this *MyCircularDeque) next(idx int) int {
     return (idx + 1) % this.cap
}



/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
