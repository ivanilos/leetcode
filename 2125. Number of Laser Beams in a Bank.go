func numberOfBeams(bank []string) int {
    result := 0
    last := 0

    for _, row := range bank {
        cur := 0
        for _, c := range row {
            if c == '1' {
                cur++
            }
        }

        if cur > 0 {
            result += last * cur
            last = cur
        }
    }

    return result
}
