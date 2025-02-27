func lenLongestFibSubseq(arr []int) int {
    result := 0

    inArr := map[int]bool{}
    for _, val := range arr {
        inArr[val] = true
    }

    for firstIdx := 0; firstIdx < len(arr); firstIdx++ {
        for secondIdx := firstIdx + 1; secondIdx < len(arr); secondIdx++ {
            a := arr[firstIdx]
            b := arr[secondIdx]

            cur := 2
            ok := inArr[a + b]
            for ok {
                cur++
                a, b = b, a + b
                ok = inArr[a + b]
            }

            if cur >= 3 {
                result = max(result, cur)
            }
        }
    }
    return result
}
