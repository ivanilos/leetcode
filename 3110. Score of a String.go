func scoreOfString(s string) int {
    result := 0
    for i := 1; i < len(s); i++ {
        result += abs(int(s[i]) - int(s[i - 1]))
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
