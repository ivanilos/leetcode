type RandomizedSet struct {
    elemsMap map[int]int
    elems []int
}


func Constructor() RandomizedSet {
    return RandomizedSet {
        map[int]int{},
        []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.elemsMap[val]; !ok {
        this.elemsMap[val] = len(this.elems)
        this.elems = append(this.elems, val)
        return true
    }
    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.elemsMap[val]; !ok {
        return false
    }
    valPosInElems := this.elemsMap[val]

    lastElem := this.elems[len(this.elems) - 1]
    this.elems[valPosInElems] = lastElem
    
    this.elemsMap[lastElem] = valPosInElems

    delete(this.elemsMap, val)

    this.elems = this.elems[0 : len(this.elems) - 1]

    return true
}


func (this *RandomizedSet) GetRandom() int {
    idx := rand.IntN(len(this.elems))
    return this.elems[idx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
