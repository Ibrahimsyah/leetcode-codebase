/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    var s func(node *TreeNode)
    s = func(node *TreeNode) {
        if node == nil {
            return
        }

        s(node.Left)
        res = append(res, node.Val)
        s(node.Right)
    }

    s(root)
    return res
}