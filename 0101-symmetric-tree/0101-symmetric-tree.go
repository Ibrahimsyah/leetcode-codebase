/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var solve func (n1 *TreeNode, n2 *TreeNode) bool
    solve = func (n1 *TreeNode, n2 *TreeNode) bool {
        if n1 == nil && n2 == nil {
            return true
        }

        if n1 == nil || n2 == nil {
            return false
        }

        return n1.Val == n2.Val && solve(n1.Left, n2.Right) && solve(n1.Right, n2.Left)
    }   

    if root == nil {
        return true
    }

    return solve(root.Left, root.Right)

}