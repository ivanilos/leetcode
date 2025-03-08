func minimumRecolors(blocks string, k int) int {
    result := len(blocks)

    left := 0
    curWhite := 0

    for right := 0; right < len(blocks); right++ {
        if blocks[right] == 'W' {
            curWhite++
        }

        if right - left + 1 > k {
            if blocks[left] == 'W' {
                curWhite--
            }
            left++
        }

        if right - left + 1 >= k {
            result = min(result, curWhite)
        }
    }

    return result
}
