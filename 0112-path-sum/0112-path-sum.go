/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    res := false
    var dfs func(n *TreeNode, sum int)
    dfs = func(n *TreeNode, sum int) {
        sum += n.Val
        if n.Left == nil && n.Right == nil {
            if sum == targetSum {
                res = true
            }
            return
        }

        if n.Left != nil {
            dfs(n.Left, sum)
        }

        if n.Right != nil {
            dfs(n.Right, sum)
        }
    }

    dfs(root, 0)

    return res
}