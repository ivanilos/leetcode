func subarraysDivByK(nums []int, k int) int {
    prefix := 0
    prefixMap := map[int]int{}
    prefixMap[0] = 1

    result := 0

    for _, num := range(nums) {
        prefix = (((prefix + num) % k) + k) % k
        result += prefixMap[prefix]

        prefixMap[prefix]++
    }
    return result
}
