/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    has map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    has := map[int]bool{}

    DFS(root, 0, has)

    return FindElements{
        has,
    }
}

func DFS(root *TreeNode, val int, has map[int]bool) {
    if root == nil {
        return
    }

    has[val] = true

    DFS(root.Left, val * 2 + 1, has)
    DFS(root.Right, val * 2 + 2, has)
}


func (this *FindElements) Find(target int) bool {
    _, ok := this.has[target]

    return ok
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
