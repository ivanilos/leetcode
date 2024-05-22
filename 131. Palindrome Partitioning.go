func partition(s string) [][]string {
    result := solve(s)

    for i := 0; i < len(result); i++ {
        result[i] = reverseString(result[i])
    }
    return result
}

func solve(s string) [][]string {
    if s == "" {
        return [][]string{}
    }
    result := [][]string{}
    for i := 0; i < len(s); i++ {
        if isPalindrome(s[0 : i + 1]) {
            if i + 1 >= len(s) {
                result = append(result, []string{s[0 : i + 1]})
            } else {
                restDivision := solve(s[i + 1 : len(s)])

                for _, rest := range(restDivision) {
                    newPartition := append(rest, s[0 : i + 1])
                    result = append(result, newPartition)
                }
            }
        }
    }
    return result
}

func reverseString(partition []string) []string {
    for i, j := 0, len(partition) - 1; i < j; i, j = i + 1, j - 1 {
        partition[i], partition[j] = partition[j], partition[i]
    }
    return partition
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}
