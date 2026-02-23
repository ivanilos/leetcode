func hasAllCodes(s string, k int) bool {
    cur := 0

    for i := 0; i < len(s) && i < k; i++ {
        cur = 2 * cur + int(s[i] - '0')
    }

    found := map[int]bool{}

    found[cur] = true
    for i := k; i < len(s); i++ {
        cur &= (1 << (k - 1)) - 1

        cur = 2 * cur + int(s[i] - '0')
        found[cur] = true
    }

    return len(found) == (1 << k)
}
