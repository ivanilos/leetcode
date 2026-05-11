func separateDigits(nums []int) []int {
    result := []int{}
    for _, num := range nums {
        cur := []int{}
        for num > 0 {
            cur = append(cur, num % 10)
            num /= 10
        }

        slices.Reverse(cur)
        result = append(result, cur...)
    }

    return result
}
