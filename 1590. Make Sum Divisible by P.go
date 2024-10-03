func minSubarray(nums []int, p int) int {
    sumModP := 0
    for _, num := range nums {
        sumModP = (sumModP + num) % p
    }
    if sumModP == 0 {
        return 0
    }

    curMod := 0
    modIdx := map[int]int{}
    modIdx[0] = -1
    result := len(nums)
    for i := 0; i < len(nums); i++ {
        curMod = (curMod + nums[i]) % p
        need := (curMod - sumModP + p) % p

        if _, ok := modIdx[need]; ok {
            result = min(result, i - modIdx[need])
        }
        modIdx[curMod] = i
    }
    if result >= len(nums) {
        return -1
    } else {
        return result
    }
}
