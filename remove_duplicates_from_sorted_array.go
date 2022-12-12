func removeDuplicates(nums []int) int {
    var k int
    start, end := 0, 1
    for ;end < len(nums); {
        if nums[end] < nums[end-1] {
            break
        }
        if nums[start] != nums[end] {
            nums[start+1] = nums[end]
            k++
            start++
        }
        end++
    }
    if start <= len(nums)-1 {
        k++
    }
    return k
}
