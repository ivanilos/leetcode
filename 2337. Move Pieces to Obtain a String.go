func canChange(start string, target string) bool {
    posStart := getPos(start)
    posTarget := getPos(target)

    if len(posStart) != len(posTarget) {
        return false
    }

    for i := 0; i < len(posStart); i++ {
        if posStart[i][1] != posTarget[i][1] {
            return false
        }

        if posStart[i][1] == -1 && posStart[i][0] < posTarget[i][0] {
            return false
        }
        if posStart[i][1] == 1 && posStart[i][0] > posTarget[i][0] {
            return false
        }
    }
    return true
}

func getPos(s string) [][]int {
    result := [][]int{}

    for i, c := range s {
        if c == 'L' {
            result = append(result, []int{i, -1})
        } else if c == 'R' {
            result = append(result, []int{i, 1})
        }
    }
    return result
}
