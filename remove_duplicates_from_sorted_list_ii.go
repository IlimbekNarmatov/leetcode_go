func deleteDuplicates(head *ListNode) *ListNode {
    r := &ListNode{}
    runner := r
    for ;head != nil; {
        if isDistinct(head) {
            runner.Next = head
            runner = runner.Next
        }
        head = nextNode(head)
    }
    runner.Next = nil
    return r.Next
}

func isDistinct(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    return head.Val != head.Next.Val
}

func nextNode(head *ListNode) *ListNode {
    for runner := head.Next; runner != nil; runner = runner.Next {
        if runner.Val != head.Val {
            return runner
        }
    }
    return nil
}
