type AllOne struct {
    elems map[int]map[string]bool
    counter map[string]int
    minVal int
    maxVal int
}


func Constructor() AllOne {
    ds := AllOne {
        map[int]map[string]bool{},
        map[string]int{},
        0,
        0,
    }

    ds.elems[0] = map[string]bool{}
    return ds
}


func (this *AllOne) Inc(key string)  {
    if _, ok := this.counter[key]; ok {
        delete(this.elems[this.counter[key]], key)
        this.counter[key]++

        if len(this.elems[this.counter[key]]) == 0 {
            this.elems[this.counter[key]] = map[string]bool{}
        }

        this.elems[this.counter[key]][key] = true
    } else {
        if len(this.elems[1]) == 0 {
            this.elems[1] = map[string]bool{}
        }
        this.elems[1][key] = true
        this.counter[key] = 1
    }
    this.maxVal = max(this.maxVal, this.counter[key])
    this.minVal = min(this.minVal, this.counter[key])
}


func (this *AllOne) Dec(key string)  {
    if _, ok := this.counter[key]; ok {
        delete(this.elems[this.counter[key]], key)
        this.counter[key]--

        if this.counter[key] != 0 {
            this.elems[this.counter[key]][key] = true
        } else {
            delete(this.counter, key)
        }
    }
    this.maxVal = max(this.maxVal, this.counter[key])
    this.minVal = min(this.minVal, this.counter[key])
}


func (this *AllOne) GetMaxKey() string {
    if len(this.counter) == 0 {
        return ""
    }

    for len(this.elems[this.maxVal]) == 0 {
        this.maxVal--
    }
    for key, _ := range(this.elems[this.maxVal]) {
        return key
    }
    return ""
}


func (this *AllOne) GetMinKey() string {
    if len(this.counter) == 0 {
        return ""
    }

    for len(this.elems[this.minVal]) == 0 {
        this.minVal++
    }
    for key, _ := range(this.elems[this.minVal]) {
        return key
    }
    return ""
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
