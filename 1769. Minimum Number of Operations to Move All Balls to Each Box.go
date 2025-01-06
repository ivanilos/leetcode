func minOperations(boxes string) []int {
    result := make([]int, len(boxes))

    curBalls := 0
    curMovements := 0
    for i := 0; i < len(boxes); i++ {
        curMovements += curBalls
        
        if boxes[i] == '1' {
            curBalls++
        }
    }

    result[len(boxes) - 1] = curMovements
    rightBalls := 0

    if boxes[len(boxes) - 1] == '1' {
        rightBalls = 1
        curBalls--
    }
    for i := len(boxes) - 2; i >= 0; i-- {
        curMovements += rightBalls
        curMovements -= curBalls
        
        result[i] = curMovements

        if boxes[i] == '1' {
            rightBalls++
            curBalls--
        }
    }

    return result
}
