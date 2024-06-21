/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }

	if head.Next == nil {
        return head
    }

    var length int
    current := head
    for current != nil{
        length++
        current = current.Next
    }

    n := 1
    k %= length
    resultHead := head
    for n <= k {
        beforeLastNode := getBeforeLastNode(resultHead)
        lastNode := beforeLastNode.Next
        lastNode.Next = resultHead
        resultHead = lastNode
        beforeLastNode.Next = nil
        n++
    }

    return resultHead
}

func getBeforeLastNode(head *ListNode) *ListNode {
    current := head
    for {
        if current.Next.Next == nil {
            return current
        }
        current = current.Next
    }
}