const INF = int64(1e15)
const MOD = int64(1e9 + 7)

func nthMagicalNumber(n int, a int, b int) int {
    low := int64(1)
    high := INF
    result := low

    d := gcd(a, b)

    for low <= high {
        mid := (low + high) / 2

        if hasSufficient(n, a, b, d, mid) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return int(result % MOD)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func hasSufficient(n, a, b, d int, maxVal int64) bool {
    lcm := int64((a / d) * b)
    total := maxVal / int64(a) + maxVal / int64(b) - maxVal / lcm

    return total >= int64(n)
}
