/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    var s func(node *TreeNode)
    s = func(node *TreeNode) {
        if node == nil {
            return
        }

        s(node.Left)
        s(node.Right)
        res = append(res, node.Val)
    }

    s(root)
    return res
}