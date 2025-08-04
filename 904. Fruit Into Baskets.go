func totalFruit(fruits []int) int {
    freq := map[int]int{}

    left := 0
    curFruits := 0
    result := 0
    for right := 0; right < len(fruits); right++ {
        freq[fruits[right]]++
        curFruits++

        for len(freq) > 2 && left <= right {
            curFruits--
            freq[fruits[left]]--
            if freq[fruits[left]] == 0 {
                delete(freq, fruits[left])
            }
            left++
        }

        result = max(result, curFruits)
    }

    return result
}
