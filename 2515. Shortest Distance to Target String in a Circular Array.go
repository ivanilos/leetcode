func closestTarget(words []string, target string, startIndex int) int {
    result := len(words)

    for i, word := range words {
        if target == word {
            result = min(result, abs(startIndex - i), len(words) - abs(startIndex - i))
        }
    }

    if result >= len(words) {
        return -1
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
