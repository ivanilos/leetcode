const COLORINGS = 27 // 3^3
const MOD = int(1e9 + 7)

func numOfWays(n int) int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, COLORINGS) // 3^3
    }

    validColorings := getValidColorings()

    // base cases
    for _, base := range validColorings {
        dp[0][base] = 1
    }
    for i := 1; i < n; i++ {
        for _, cur := range validColorings {
            for _, prev := range validColorings {
                if notMatchColors(cur, prev) {
                    dp[i][cur] = add(dp[i][cur], dp[i - 1][prev])
                }
            }
        }
    }

    result := 0
    for _, last := range validColorings {
        result = add(result, dp[n - 1][last])
    }
    
    return result
}

func getValidColorings() []int {
    result := []int{}
    for i := 0; i < COLORINGS; i++ {
        a := i % 3
        b := (i / 3) % 3
        c := (i / 9)

        if a != b && b != c {
            result = append(result, i)
        }
    }

    return result
}

func notMatchColors(x, y int) bool {
    a1 := x % 3
    a2 := y % 3

    if a1 == a2 {
        return false
    }

    b1 := (x / 3) % 3
    b2 := (y / 3) % 3

    if b1 == b2 {
        return false
    }

    c1 := (x / 9)
    c2 := (y / 9)

    if c1 == c2 {
        return false
    }

    return true
}

func add(a, b int) int {
    return (a + b) % MOD
}
