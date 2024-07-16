/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
    pathToStart := []rune{}
    pathToDest := []rune{}
    DFS(root, startValue, destValue, []rune{}, &pathToStart, &pathToDest)

    return solve(pathToStart, pathToDest)
}

func DFS(root *TreeNode, startValue, destValue int, curPath []rune, pathToStart, pathToDest *[]rune) {
    if root == nil {
        return
    }

    if root.Val == startValue {
        *pathToStart = append([]rune(nil), curPath...)
    } else if root.Val == destValue {
        *pathToDest = append([]rune(nil), curPath...)
    }


    curPath = append(curPath, 'L')
    DFS(root.Left, startValue, destValue, curPath, pathToStart, pathToDest)
    curPath = curPath[0 : len(curPath) - 1]

    curPath = append(curPath, 'R')
    DFS(root.Right, startValue,destValue, curPath, pathToStart, pathToDest)
    curPath = curPath[0 : len(curPath) - 1]
}

func solve(pathToStart, pathToDest []rune) string {
    equalPrefixLen := 0

    lenStart := len(pathToStart)
    lenDest := len(pathToDest)

    for i := 0; i < min(lenStart, lenDest); i++ {
        if pathToStart[i] == pathToDest[i] {
            equalPrefixLen++
        } else {
            break;
        }
    }

    result := []rune{}
    for i := equalPrefixLen; i < lenStart; i++ {
        result = append(result, 'U')
    }
    for i := equalPrefixLen; i < lenDest; i++ {
        result = append(result, pathToDest[i])
    }

    return string(result)
}
