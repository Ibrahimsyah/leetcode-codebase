/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    var solve func(node *TreeNode, temp int) int
    solve = func(node *TreeNode, temp int) int{
        if node == nil {
            return 0
        }

        res := temp * 2 + node.Val
        if node.Left == nil && node.Right == nil {
            return res
        }

        return solve(node.Left, res) + solve(node.Right, res)
    }


    return solve(root, 0)
}