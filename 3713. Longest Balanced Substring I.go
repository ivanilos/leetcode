const ALPHA = 26

func longestBalanced(s string) int {
    prefixSums := getPrefixSums(s)

    for sz := len(s); sz >= 1; sz-- {
        for start := 0; start + sz <= len(s); start++ {
            if isBalanced(start, start + sz - 1, prefixSums) {
                return sz
            }
        }
    }

    return 0
}

func isBalanced(start, end int, prefixSums [][]int) bool {
    nonZeroMaxFreq := 0
    for i := 0; i < ALPHA; i++ {
        curTotal := prefixSums[i][end]
        if start > 0 {
            curTotal -= prefixSums[i][start - 1]
        }

        if curTotal != 0 {
            if nonZeroMaxFreq == 0 {
                nonZeroMaxFreq = curTotal
            } else if curTotal != nonZeroMaxFreq {
                return false
            }
        }
    }

    return true
}

func getPrefixSums(s string) [][]int {
    prefixSums := make([][]int, ALPHA)
    for i := 0; i < ALPHA; i++ {
        prefixSums[i] = make([]int, len(s))
    }

    firstCharIdx := int(s[0] - 'a')
    prefixSums[firstCharIdx][0] = 1

    for i := 1; i < len(s); i++ {
        curChar := int(s[i] - 'a')

        prefixSums[curChar][i] = 1

        for j := 0; j < ALPHA; j++ {
            prefixSums[j][i] += prefixSums[j][i - 1]
        }
    }

    return prefixSums
}
