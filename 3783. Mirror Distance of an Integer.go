func mirrorDistance(n int) int {
    mirror := getMirror(n)

    return max(n, mirror) - min(n, mirror)
}

func getMirror(n int) int {
    result := 0

    for n > 0 {
        digit := n % 10
        result = 10 * result + digit

        n /= 10
    }

    return result
}
