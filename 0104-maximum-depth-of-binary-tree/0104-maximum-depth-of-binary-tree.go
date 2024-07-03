/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    max := 0
    var dfs func (n *TreeNode, d int)
    dfs = func(n *TreeNode, d int) {
        if n == nil {
            return
        }

        if d > max {
            max = d
        }

        dfs(n.Left, d + 1)
        dfs(n.Right, d + 1)
    }

    dfs(root, 1)
    return max
}