func majorityElement(nums []int) int {
    candidate := nums[0]
    qt := 0
    for _, elem := range nums {
        if elem == candidate {
            qt++
        } else {
            qt--
            if qt < 0 {
                candidate = elem
                qt = 1
            }
        }
    }
    return candidate
}
