/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodeMap := map[int]*TreeNode{}
    childrenMap := map[int]bool{}

    for _, description := range(descriptions) {
        par := description[0]
        val := description[1]
        isLeft := description[2]

        if _, ok := nodeMap[val]; !ok {
            nodeMap[val] = &TreeNode{val, nil, nil}
        }
        node := nodeMap[val]
        childrenMap[val] = true

        if _, ok := nodeMap[par]; !ok {
            nodeMap[par] = &TreeNode{par, nil, nil}
        }
        parentNode := nodeMap[par]
        if isLeft == 1 {
            parentNode.Left = node
        } else {
            parentNode.Right = node
        }
    }

    for key, _ := range(nodeMap) {
        if _, ok := childrenMap[key]; !ok {
            return nodeMap[key]
        }
    }
    panic("should not reach this")
}
