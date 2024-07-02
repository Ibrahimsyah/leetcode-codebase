/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var solve func(p *TreeNode, q *TreeNode) bool 
    solve = func (p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        }

        if (p != nil && q == nil) || (p == nil && q != nil) {
            return false
        }

        if p.Val != q.Val {
            return false
        }

        return solve(p.Left, q.Left) && solve(p.Right, q.Right)
    }

    return solve(p, q)
}