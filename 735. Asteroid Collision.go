func asteroidCollision(asteroids []int) []int {
    result := []int{}

    for _, asteroid := range asteroids {
        isNextBroke := false
        for len(result) > 0 && sign(result[len(result) - 1]) == 1 && sign(asteroid) == -1 {
            if abs(asteroid) > result[len(result) - 1] {
                result = result[: len(result) - 1]
            } else if abs(asteroid) < result[len(result) - 1] {
                isNextBroke = true
                break
            } else {
                isNextBroke = true
                result = result[: len(result) - 1]
                break
            }
        }

        if !isNextBroke {
            result = append(result, asteroid)
        }
    }

    return result
}

func sign(x int) int {
    if x < 0 {
        return -1
    } else {
        return 1
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
