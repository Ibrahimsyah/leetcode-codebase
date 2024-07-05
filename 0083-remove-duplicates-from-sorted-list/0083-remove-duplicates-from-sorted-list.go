/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    p := head
    unique := head
    curr := head.Next
    for curr != nil {
        if curr.Val != unique.Val {
            unique.Next = curr
            unique = curr
        }
        curr = curr.Next
    }
    unique.Next = nil

    return p
}