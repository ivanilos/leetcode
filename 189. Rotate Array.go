func rotate(nums []int, k int)  {
    d := gcd(len(nums), k)

    for i := 0; i < d; i++ {
        rotateCycle(i, len(nums) / d, k, nums)
    }
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func rotateCycle(start, times, k int, nums []int) {
    pos := start
    for i := 0; i < times; i++ {
        nextPos := (pos + k) % len(nums)
        nums[start], nums[nextPos] = nums[nextPos], nums[start]
        pos = nextPos
    }
}
