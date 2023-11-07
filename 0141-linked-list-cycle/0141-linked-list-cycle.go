/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    t := head
    h := head
    for {
        if t == nil || h == nil || h.Next == nil {
            return false
        }

        t = t.Next
        h = h.Next.Next

        if t != nil && h != nil && t == h {
            return true
        }
    }

    return false
}