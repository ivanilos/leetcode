func decrypt(code []int, k int) []int {
    result := make([]int, len(code))
    for i := 0; i < len(code); i++ {
        sum := 0
        delta := 1
        if k < 0 {
            delta = -1
        }
        for j := 1; j <= k * delta; j++ {
            idx :=((i + j * delta) % len(code) + len(code)) % len(code)
            sum += code[idx]
        }

        result[i] = sum
    }

    return result
}
