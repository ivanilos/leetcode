func frequencySort(nums []int) []int {
    freq := map[int]int{}

    for _, num := range(nums) {
        freq[num]++
    }

    sort.Slice(nums, func(a, b int) bool {
        return freq[nums[a]] < freq[nums[b]] || 
                freq[nums[a]] == freq[nums[b]] && nums[a] > nums[b]
    })

    return nums
}
