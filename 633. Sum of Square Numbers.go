func judgeSquareSum(c int) bool {
    for a := 0; a * a <= c; a++ {
        if isPerfectSquare(c - a * a) {
            return true
        }
    }
    return false
}

func isPerfectSquare(x int) bool {
    sqRoot := int(math.Sqrt(float64(x)))

    for sqRoot * sqRoot < x {
        sqRoot++
    }
    for sqRoot * sqRoot > x {
        sqRoot--
    }
    return sqRoot * sqRoot == x
}
