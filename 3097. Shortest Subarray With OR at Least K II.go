const wordSize int = 31

func minimumSubarrayLength(nums []int, k int) int {
    onBits := getOnBits(nums)

    l := 0
    curOr := 0
    freq := make([]int, wordSize)

    result := len(nums) + 1
    for r := 0; r < len(nums); r++ {
        for i := 0; i < wordSize; i++ {
            freq[i] += onBits[r][i]
        }
        curOr |= nums[r]

        for l < r && curOr >= k {
            result = min(result, r - l + 1)
            for i := 0; i < wordSize; i++ {
                freq[i] -= onBits[l][i]
                if freq[i] == 0 && onBits[l][i] == 1 {
                    curOr ^= 1 << i
                }
            }
            l++
        }

        if curOr >= k {
            result = min(result, r - l + 1)
        }
    }

    if result > len(nums) {
        return -1
    }
    return result
}

func getOnBits(nums []int) [][]int {
    onBits := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        onBits[i] = make([]int, wordSize)
    }

    for i := 0; i < len(nums); i++ {
        for j := 0; j < wordSize; j++ {
            bit := (nums[i] >> j) & 1

            onBits[i][j] = bit
        }
    }

    return onBits
}
