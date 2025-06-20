func maxDistance(s string, k int) int {
    x := 0
    y := 0

    result := 0

    dir := map[rune]int{}
    for _, c := range s {
        x, y = getPos(x, y, c)

        dir[c]++

        result = max(result, getMaxManhattan(x, y, dir, k))
    }

    return result
}

func getPos(x, y int, c rune) (int, int) {
    if c == 'N' {
        return x - 1, y
    } else if c == 'S' {
        return x + 1, y
    } else if c == 'W' {
        return x, y - 1
    } else {
        return x, y + 1
    }
}

func getMaxManhattan(x, y int, dir map[rune]int, k int) int {
    return max(getMaxManhattanNorth(x, y, dir, k), getMaxManhattanSouth(x, y, dir, k))
}

func getMaxManhattanNorth(x, y int, dir map[rune]int, k int) int {
    maxNorth := abs(x - 2 * min(dir['S'], k))
    k -= min(dir['S'], k)

    maxLeft := abs(y - 2 * min(dir['E'], k))
    maxRight := abs(y + 2 * min(dir['W'], k))

    return maxNorth + max(maxLeft, maxRight)
}

func getMaxManhattanSouth(x, y int, dir map[rune]int, k int) int {
    maxNorth := abs(x + 2 * min(dir['N'], k))
    k -= min(dir['N'], k)

    maxLeft := abs(y - 2 * min(dir['E'], k))
    maxRight := abs(y + 2 * min(dir['W'], k))

    return maxNorth + max(maxLeft, maxRight)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
