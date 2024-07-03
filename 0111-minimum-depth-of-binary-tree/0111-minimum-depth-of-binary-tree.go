/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    min := math.MaxInt
    var dfs func(n *TreeNode, d int)
    dfs = func(n *TreeNode, d int) {
        if n == nil {
            return
        }

        if n.Left == nil && n.Right == nil {
            if d < min {
                min = d
            }
            return
        }

        dfs(n.Left, d + 1)
        dfs(n.Right, d + 1)
    }

    dfs(root, 1)
    return min
}