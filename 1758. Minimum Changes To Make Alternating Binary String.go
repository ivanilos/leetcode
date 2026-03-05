func minOperations(s string) int {
    even := []int{0, 0}
    odd := []int{0, 0}

    for i, c := range s {
        val := int(c - '0')

        if i % 2 == 0 {
            even[val]++
        } else {
            odd[val]++
        }
    }

    return min(even[0] + odd[1], even[1] + odd[0])
}
