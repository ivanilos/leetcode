func numberOfAlternatingGroups(colors []int, k int) int {
    sz := len(colors)

    result := 0

    left := 0
    alternatingSize := 0
    last := 1 - colors[0]
    for right := 0; right < sz + k - 1; right++ {
        if colors[right % sz] != last {
            alternatingSize++
        } else {
            alternatingSize = 1
            left = right
        }
        last = colors[right % sz]


        for left < right && right - left + 1 > k {
            alternatingSize--
            left++
        }

        if alternatingSize == k {
            result++
        }
    }

    return result
}
