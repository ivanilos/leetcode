func minFlips(s string) int {
    sz := len(s)

    s += s

    parity := make([][]int, 2)
    parity[0] = []int{0, 0}
    parity[1] = []int{0, 0}

    for i := 0; i < sz; i++ {
        val := int(s[i] - '0')
        parity[val][i % 2]++
    }

    result := min(parity[0][0] + parity[1][1], parity[0][1] + parity[1][0])
    for i := sz; i < len(s); i++ {
        val := int(s[i] - '0')
        parity[val][i % 2]++

        prevIdx := i - sz
        prevVal := int(s[prevIdx] - '0')
        parity[prevVal][prevIdx % 2]--

        result = min(result, parity[0][0] + parity[1][1], parity[0][1] + parity[1][0])
    }

    return result
}
