var INF int = int(1e9)
var LEFT_STRING int = 0
var RIGHT_STRING int = 1
var BOTH int = 2

func shortestCommonSupersequence(str1 string, str2 string) string {
    dp, change := initDP(str1, str2)

    _ = solve(0, 0, str1, str2, dp, change)
    result := getAnswer(dp, change, str1, str2)

    return result
}

func initDP(str1, str2 string) ([][]int, [][]int) {
    dp := make([][]int, len(str1) + 1)
    change := make([][]int, len(str1) + 1)
    for i := 0; i <= len(str1); i++ {
        dp[i] = make([]int, len(str2) + 1)
        change[i] = make([]int, len(str2) + 1)

        for j := 0; j <= len(str2); j++ {
            dp[i][j] = -1
            change[i][j] = -1
        }
    }

    return dp, change

}

func getAnswer(dp, change [][]int, str1, str2 string) string {
    pos1 := 0
    pos2 := 0

    result := []byte{}
    for pos1 < len(str1) || pos2 < len(str2) {
        if change[pos1][pos2] == LEFT_STRING {
            result = append(result, str1[pos1])
            pos1++
        } else if change[pos1][pos2] == RIGHT_STRING {
            result = append(result, str2[pos2])
            pos2++
        } else if change[pos1][pos2] == BOTH {
            result = append(result, str1[pos1])
            pos1++
            pos2++
        }
    }

    return string(result)
}

func solve(pos1, pos2 int, str1, str2 string, dp, change [][]int) int {
    if pos1 >= len(str1) && pos2 >= len(str2) {
        return 0
    }
    if dp[pos1][pos2] != -1 {
        return dp[pos1][pos2]
    }

    dp[pos1][pos2] = INF

    if pos1 < len(str1) {
        aux := 1 + solve(pos1 + 1, pos2, str1, str2, dp, change)
        if aux < dp[pos1][pos2] {
            dp[pos1][pos2] = aux
            change[pos1][pos2] = LEFT_STRING
        }
    }
    if pos2 < len(str2) {
        aux := 1 + solve(pos1, pos2 + 1, str1, str2, dp, change)
        if aux < dp[pos1][pos2] {
            dp[pos1][pos2] = aux
            change[pos1][pos2] = RIGHT_STRING
        }
    }

    if  pos1 < len(str1) && pos2 < len(str2) && str1[pos1] == str2[pos2] {
        aux := 1 + solve(pos1 + 1, pos2 + 1, str1, str2, dp, change)
        if aux < dp[pos1][pos2] {
            dp[pos1][pos2] = aux
            change[pos1][pos2] = BOTH
        }
    }

    return dp[pos1][pos2]
}
