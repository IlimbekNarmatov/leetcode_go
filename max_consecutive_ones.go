func findMaxConsecutiveOnes(nums []int) int {
    var max int
    lastZero := -1
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            lastZero = i
        } else if lastZero < 0 {
            max = i + 1
        } else if max < i - lastZero {
            max = i - lastZero
        }
    }
    return max
}
