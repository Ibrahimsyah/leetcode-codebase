/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2 
    }

    if list2 == nil {
        return list1
    }

    current := list1
    if list1.Val > list2.Val {
        current = list2
        list2 = list2.Next
    } else {
        list1 = list1.Next
    }
    head := current

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }

        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    }

    if list2 != nil {
        current.Next = list2
    }

    return head
}