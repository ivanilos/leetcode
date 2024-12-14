func continuousSubarrays(nums []int) int64 {
    result := int64(0)

    freq := map[int]int{}
    l := 0
    for r := 0; r < len(nums); r++ {
        freq[nums[r]]++

        for l < r && !isContinuous(freq) {
            freq[nums[l]]--
            if freq[nums[l]] == 0 {
                delete(freq, nums[l])
            }
            l++
        }
        result += int64(r - l + 1)
    }
    return result
}

func isContinuous(freq map[int]int) bool {
    mini := -1
    maxi := -1
    for key, _ := range freq {
        if mini == -1 || key < mini {
            mini = key
        }
        if maxi == -1 || key > maxi {
            maxi = key
        }
    }
    return maxi - mini <= 2
}
