func findSubarrays(nums []int) bool {
    m := make(map[int]bool)
    for i := 1; i < len(nums); i++ {
        if v, ok := m[nums[i] + nums[i-1]]; ok && v {
            return true
        }
        m[nums[i] + nums[i-1]] = true
    }
    return false
}
