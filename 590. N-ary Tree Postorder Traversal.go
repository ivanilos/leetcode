/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    
    stack := []*Node{}

    result := []int{}

    stack = append(stack, root)
    processed := map[*Node]bool{}

    for len(stack) > 0 {
        next := stack[len(stack) - 1]
        stack = stack[0 : len(stack) - 1]

        if _, ok := processed[next]; ok {
            result = append(result, next.Val)
        } else {
            stack = append(stack, next)
            processed[next] = true
            for i := len(next.Children) - 1; i >= 0; i-- {
                stack = append(stack, next.Children[i])
            }
        }
    }
    return result
}
