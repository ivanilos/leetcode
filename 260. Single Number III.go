func singleNumber(nums []int) []int {
    xor := 0
    for _, num := range(nums) {
        xor ^= num
    }

    result := []int{0, 0}
    lastSetBit := findLastSetBit(xor)

    for _, num := range(nums) {
        if ((num >> lastSetBit) & 1) == 1 {
            result[0] ^= num
        } else {
            result[1] ^= num
        }
    }
    return result
}

func findLastSetBit(x int) int {
    for i := 0; i < 32; i++ {
        if ((x >> i) & 1) == 1 {
            return i
        }
    }
    return -1
}
