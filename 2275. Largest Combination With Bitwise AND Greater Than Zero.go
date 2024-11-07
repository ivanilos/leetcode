func largestCombination(candidates []int) int {
    freqOnBits := map[int]int{}

    for _, num := range candidates {
        update(freqOnBits, num)
    }

    result := 0
    for _, val := range freqOnBits {
        result = max(result, val)
    }
    return result
}

func update(freqOnBits map[int]int, num int) {
    for i := 0; (1 << i) <= num; i++ {
        bit := (num >> i) & 1

        if bit == 1 {
            freqOnBits[i]++
        }
    }
}
