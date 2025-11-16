const MOD = int64(1e9 + 7)

func numSub(s string) int {
    result := int64(0)

    cur := int64(0)
    for _, c := range s {
        if c == '1' {
            cur++
        } else {
            result += (cur * (cur + 1)) / 2
            cur = 0
        }
    }

    result += (cur * (cur + 1)) / 2

    return int(result % MOD)
}
