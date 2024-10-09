func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    return reversed(x) == x
}

func reversed(x int) int {
    result := 0
    for x > 0 {
        result = 10 * result + x % 10
        x /= 10
    }
    return result
}
