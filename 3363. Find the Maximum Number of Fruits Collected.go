func maxCollectedFruits(fruits [][]int) int {
    return getDiagonal(fruits) + getStartFromBottomLeft(fruits) + getStartFromTopRight(fruits)
}

func getDiagonal(fruits [][]int) int {
    result := 0

    for i := 0; i < len(fruits); i++ {
        result += fruits[i][i]
    }

    return result
}

func getStartFromBottomLeft(fruits [][]int) int {
    n := len(fruits)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }

    return DFSToRight(n - 1, 0, n, fruits, dp)
}

func DFSToRight(x, y, n int, fruits [][]int, dp [][]int) int {
    if x == n - 1 && y == n - 1 {
        return 0
    }

    if dp[x][y] != -1 {
        return dp[x][y]
    }


    dp[x][y] = fruits[x][y]
    for i := -1; i <= 1; i++ {
        nx := x + i
        ny := y + 1

        if nx < n && nx > ny {
            dp[x][y] = max(dp[x][y], fruits[x][y] + DFSToRight(nx, ny, n, fruits, dp))
        }
    }

    return dp[x][y]
}

func getStartFromTopRight(fruits [][]int) int {
    n := len(fruits)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }

    return DFSToBottom(0, n - 1, n, fruits, dp)
}

func DFSToBottom(x, y, n int, fruits [][]int, dp [][]int) int {
    if x == n - 1 && y == n - 1 {
        return 0
    }

    if dp[x][y] != -1 {
        return dp[x][y]
    }


    dp[x][y] = fruits[x][y]
    for j := -1; j <= 1; j++ {
        nx := x + 1
        ny := y + j

        if ny < n && nx < ny {
            dp[x][y] = max(dp[x][y], fruits[x][y] + DFSToBottom(nx, ny, n, fruits, dp))
        }
    }

    return dp[x][y]
}
