func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }
    slow := head
    fast := head.Next
    for ;fast.Next != nil && fast.Next.Next != nil; {
        fast = fast.Next.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return head
}
