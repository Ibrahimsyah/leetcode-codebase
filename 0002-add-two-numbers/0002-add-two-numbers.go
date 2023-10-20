/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var r *ListNode
    var c *ListNode
    var s int
    for {
        if l1 == nil && l2 == nil && s == 0 {
            break
        }

        if c != nil {
            c.Next = &ListNode{}
            c = c.Next
        } else {
            c = &ListNode{}
        }

        val := s
        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }

        s = val / 10
        val = val % 10
        c.Val = val
        if r == nil {
            r = c
        }
    }

    return r
}