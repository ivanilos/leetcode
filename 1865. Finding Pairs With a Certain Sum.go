type FindSumPairs struct {
    nums1 []int
    nums2Map map[int]int
    freqNums2 map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    nums2Map := map[int]int{}
    freqNums2 := map[int]int{}

    for i, num := range nums2 {
        nums2Map[i] = num
        freqNums2[num]++
    }

    return FindSumPairs{
        nums1,
        nums2Map,
        freqNums2,
    }
}


func (this *FindSumPairs) Add(index int, val int)  {
    curValue := this.nums2Map[index]
    this.freqNums2[curValue]--

    newVal := curValue + val
    this.nums2Map[index] = newVal
    this.freqNums2[newVal]++ 
}


func (this *FindSumPairs) Count(tot int) int {
    result := 0

    for _, val := range this.nums1 {
        result += this.freqNums2[tot - val]
    }

    return result
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
