/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverFromPreorder(traversal string) *TreeNode {
    values, dashCount := parse(traversal)

    root := &TreeNode{values[0], nil, nil}
    st := []*TreeNode{root}

    nextNodeIdx := 1
    nextDashCountIdx := 0
    for nextNodeIdx < len(values) {
        for len(st) > dashCount[nextDashCountIdx] {
            st = st[: len(st) - 1]
        }

        newNode := &TreeNode{values[nextNodeIdx], nil, nil}

        lastElem := st[len(st) - 1]
        if lastElem.Left == nil {
            lastElem.Left = newNode
        } else {
            lastElem.Right = newNode
        }

        st = append(st, newNode)
        nextNodeIdx++
        nextDashCountIdx++
    }

    return root
}

func parse(traversal string) ([]int, []int) {
    values := []int{}
    dashCount := []int{}

    curValue := 0
    curDashCount := 0
    for _, c := range traversal {
        if unicode.IsDigit(c) {
            updateDashCount(curDashCount, &dashCount)
            curDashCount = 0
            curValue = 10 * curValue + int(c - '0')
        } else {
            updateValues(curValue, &values)
            curValue = 0
            curDashCount++
        }
    }

    updateValues(curValue, &values)
    updateDashCount(curDashCount, &dashCount)

    return values, dashCount
}

func updateValues(curValue int, values *[]int) {
    if curValue > 0 {
        *values = append(*values, curValue)
    }
}

func updateDashCount(curDashCount int, dashCount *[]int) {
    if curDashCount > 0 {
        *dashCount = append(*dashCount, curDashCount)
    }
}
