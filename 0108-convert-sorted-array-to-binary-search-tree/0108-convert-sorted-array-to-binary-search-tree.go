/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    var d func(nums []int) *TreeNode
    d = func(n []int) *TreeNode {
        if len(n) == 0 {
            return nil
        }   
        m := len(n) / 2
        return &TreeNode{
            Val: n[m],
            Left: d(n[:m]),
            Right: d(n[m+1:]),
        }
    }

    return d(nums)
}