/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := getNodesLength(head)
    target := length - n
    if target == 0 {
        return head.Next
    }

    current := head
    for i := 0; i < target - 1; i++ {
        current = current.Next
    }
    current.Next = current.Next.Next

    return head
}

func getNodesLength(head *ListNode) int {
    result := 0
    current := head
    for current != nil{
        result ++
        current = current.Next
    }

    return result
}