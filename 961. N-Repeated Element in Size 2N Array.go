func repeatedNTimes(nums []int) int {
    freq := map[int]int{}

    for _, num := range nums {
        freq[num]++

        if freq[num] > 1 {
            return num
        }
    }

    panic("Should not reach here")
}
