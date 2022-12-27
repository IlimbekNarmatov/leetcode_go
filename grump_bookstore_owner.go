func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    maxUnsatisfiedIdx := -1
    maxUnsatisfied := -1
    var start int
    var end int
    var unsatisfied int
    for end < len(customers) {
        if end - start + 1 > minutes {
            if grumpy[start] == 1 {
                unsatisfied -= customers[start]
            }
            start++
            continue
        } 
        if grumpy[end] == 1 {
            unsatisfied += customers[end]
        }
        if end - start + 1 == minutes && unsatisfied > maxUnsatisfied {
            maxUnsatisfied = unsatisfied
            maxUnsatisfiedIdx = start
        }
        end++
    }
    for i := maxUnsatisfiedIdx; i < len(grumpy) && i < maxUnsatisfiedIdx + minutes; i++ {
        grumpy[i] = 0
    }
    var result int
    for i := 0; i < len(customers); i++ {
        if grumpy[i] == 0 {
            result += customers[i]
        }
    }
    return result
}
