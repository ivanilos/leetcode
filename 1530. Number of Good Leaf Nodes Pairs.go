/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
    adjList := map[*TreeNode][]*TreeNode{}
    leafMap := map[*TreeNode]bool{}

    createAdjacencyList(root, adjList, leafMap)

    result := 0
    for node, _ := range(leafMap) {
        result += DFS(node, adjList, distance, leafMap, nil)
        fmt.Println(result) 
    }
    return result / 2
}

func createAdjacencyList(root *TreeNode, adjList map[*TreeNode][]*TreeNode, leafMap map[*TreeNode]bool) {
    if root == nil {
        return
    }

    if root.Left != nil {
        adjList[root] = append(adjList[root], root.Left)
        adjList[root.Left] = append(adjList[root.Left], root)
        createAdjacencyList(root.Left, adjList, leafMap)
    }

    
    if root.Right != nil {
        adjList[root] = append(adjList[root], root.Right)
        adjList[root.Right] = append(adjList[root.Right], root)
        createAdjacencyList(root.Right, adjList, leafMap)
    }

    if root.Left == nil && root.Right == nil {
        leafMap[root] = true
    }
}

func DFS(node *TreeNode, adjList map[*TreeNode][]*TreeNode, distance int, leafMap map[*TreeNode]bool, par *TreeNode) int {
    if distance < 0 {
        return 0
    }

    result := 0
    if distance >= 0 && leafMap[node] && par != nil {
        result++
    }

    for _, viz := range(adjList[node]) {
        if viz != par {
            result += DFS(viz, adjList, distance - 1, leafMap, node)
        }
    }
    return result
}
