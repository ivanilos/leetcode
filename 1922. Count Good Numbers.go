var MOD = int64(1e9 + 7)

func countGoodNumbers(n int64) int {
    even := (n + 1) / 2
    odd := n - even

    result := fastExp(even, 5) * fastExp(odd, 4)

    return int(result % MOD)
}

func fastExp(p, base int64) int64 {
    if p == 0 {
        return 1
    }

    result := int64(1)
    for p > 0 {
        if p & 1 != 0 {
            result = (result * base) % MOD
        }
        base = (base * base) % MOD
        p /= 2
    }
    return result
}
