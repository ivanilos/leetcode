func countPalindromicSubsequence(s string) int {
    prefSum := make([][]int, 26)
    for i := 0; i < 26; i++ {
        prefSum[i] = make([]int, len(s))
    }

    outerIdx := make([][]int, 26)
    for i := 0; i < 26; i++ {
        outerIdx[i] = []int{len(s), -1}
    }

    for i := 0; i < 26; i++ {
        if s[0] == byte('a' + i) {
            outerIdx[i][0] = 0
            outerIdx[i][1] = 0
            prefSum[i][0] = 1
        }

        for j := 1; j < len(s); j++ {
            if s[j] == byte('a' + i) {
                prefSum[i][j] = 1
                outerIdx[i][0] = min(outerIdx[i][0], j)
                outerIdx[i][1] = max(outerIdx[i][1], j)
            }
            prefSum[i][j] += prefSum[i][j - 1]
        }
    }

    result := 0
    for outer := 0; outer < 26; outer++ {
        if outerIdx[outer][0] < len(s) {
            leftIdx := outerIdx[outer][0]
            rightIdx := outerIdx[outer][1]

            for inner := 0; inner < 26; inner++ {
                if rightIdx - 1 >= 0 && prefSum[inner][rightIdx - 1] - prefSum[inner][leftIdx] > 0 {
                    result++
                }
            }
        }
    }

    return result
}
