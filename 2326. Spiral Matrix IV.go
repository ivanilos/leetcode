/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    result := make([][]int, m)
    for i := 0; i < m; i++ {
        result[i] = make([]int, n)
        for j := 0; j < n; j++ {
            result[i][j] = -1
        }
    }

    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    dir := 0
    nextHor := n - 1
    nextVer := m - 1

    x := 0
    y := 0

    for head != nil {
        if head != nil {
            result[x][y] = head.Val
            head = head.Next
        }
        for i := 0; i < nextHor && head != nil; i++ {
            x += dx[dir]
            y += dy[dir]
            result[x][y] = head.Val
            head = head.Next
        }
        dir = (dir + 1) % 4
        for i := 0; i < nextVer && head != nil; i++ {
            x += dx[dir]
            y += dy[dir]
            result[x][y] = head.Val
            head = head.Next
        }

        nextVer--

        dir = (dir + 1) % 4
        for i := 0; i < nextHor && head != nil; i++ {
            x += dx[dir]
            y += dy[dir]
            result[x][y] = head.Val
            head = head.Next
        }
        dir = (dir + 1) % 4
        for i := 0; i < nextVer && head != nil; i++ {
            x += dx[dir]
            y += dy[dir]
            result[x][y] = head.Val
            head = head.Next
        }
        dir = (dir + 1) % 4

        x += dx[dir]
        y += dy[dir]
        m = m - 2
        n =  n - 2
        nextHor = n - 1
        nextVer = m - 1

    }
    return result
}
