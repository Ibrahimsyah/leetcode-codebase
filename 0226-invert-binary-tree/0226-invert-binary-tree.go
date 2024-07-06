/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    var d func(node *TreeNode)
    d = func (node *TreeNode) {
        if node == nil {
            return
        }

        temp := node.Left
        node.Left = node.Right
        node.Right = temp
        d(node.Left)
        d(node.Right)
    }

    d(root)
    return root
}