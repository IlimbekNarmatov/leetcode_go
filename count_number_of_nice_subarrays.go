func numberOfSubarrays(nums []int, k int) int {
    var start int
    var end int
    var result int
    var oddCount int
    var temp int
    for end < len(nums) {
        if nums[end] % 2 != 0 {
            oddCount
            temp = 0
        }
        for oddCount == k {
            temp++
            if nums[start] % 2 != 0 { oddCount-- }
            start++
        }
        result += temp
        end++
    }
    return result
}
