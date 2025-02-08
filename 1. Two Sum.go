func twoSum(nums []int, target int) []int {
    elemToIdx := map[int]int{}

    for i, num := range nums {
        need := target - num

        if val, ok := elemToIdx[need]; ok {
            return []int{val, i}
        }
        elemToIdx[num] = i
    }
    panic("Should not reach here")
}
