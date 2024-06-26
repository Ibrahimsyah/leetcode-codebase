/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    var result, head *ListNode
    remainder := 0
    for l1 != nil || l2 != nil {
        val := remainder

        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }
        remainder = val / 10
        if val >= 10 {
            val -= 10
        }

        if result == nil {
            result = &ListNode{Val: val}
            head = result
        } else {
            result.Next = &ListNode{Val: val}
            result = result.Next
        }
    }
    
    if remainder != 0 {
        result.Next = &ListNode{Val: remainder}
    }

    return head
}