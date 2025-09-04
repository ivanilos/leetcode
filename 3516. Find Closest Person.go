func findClosest(x int, y int, z int) int {
    xDist := abs(z - x)
    yDist := abs(z - y)

    if xDist < yDist {
        return 1
    } else if xDist > yDist {
        return 2
    } else {
        return 0
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
