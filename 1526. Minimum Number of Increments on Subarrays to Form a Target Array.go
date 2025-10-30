const LOG = 20

func minNumberOperations(target []int) int {
    rmq := buildRMQ(target)

    result := solve(target, rmq, 0, len(target) - 1, 0)

    return result
}

func buildRMQ(target []int) [][]int {
    n := len(target)

    rmq := make([][]int, n)
    for i := 0; i < n; i++ {
        rmq[i] = make([]int, LOG)
        rmq[i][0] = i
    }

    for j := 1; j < LOG; j++ {
        for i := 0; i < n; i++ {
            next := min(n - 1, i + (1 << (j - 1)))

            leftPart := target[rmq[i][j - 1]]
            rightPart := target[rmq[next][j - 1]]

            if leftPart <= rightPart {
                rmq[i][j] = rmq[i][j - 1]
            } else {
                rmq[i][j] = rmq[next][j - 1]
            }
        }
    }

    return rmq
}

func solve(target []int, rmq [][]int, left, right, operationsDone int) int {
    if left > right {
        return 0
    }

    minIdx := getMinValueIdx(target, rmq, left, right)

    spend := max(0, target[minIdx] - operationsDone)

    newOperationsDone := max(target[minIdx], operationsDone)
    return spend + solve(target, rmq, left, minIdx - 1, newOperationsDone) + 
            solve(target, rmq, minIdx + 1, right, newOperationsDone)
}

func getMinValueIdx(target []int, rmq [][]int, left, right int) int {
    len := right - left + 1

    log := 0
    for (1 << (log + 1)) < len {
        log++
    }

    next := right - (1 << log) + 1

    leftPart := target[rmq[left][log]]
    rightPart := target[rmq[next][log]]

    if leftPart <= rightPart {
        return rmq[left][log]
    } else {
        return rmq[next][log]
    }
}
