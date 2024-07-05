/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    // traverse the list is way faster than multiple 2 pointer
    nums := []int{}
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }

    var d func(n []int) *TreeNode
    d = func (n []int) *TreeNode{
        if len(n) == 0 {
            return nil
        }

        m := len(n)/2
        return &TreeNode{
            Val: n[m],
            Left: d(n[:m]),
            Right: d(n[m+1:]),
        }
    }

    return d(nums)
}