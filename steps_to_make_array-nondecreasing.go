func totalSteps(nums []int) int {
    var steps int
    var isNondecreasing bool
    marked := make([int]bool)
    for ; !isNondecreasing; {
        var foundDecreasing
        for i := 0; i < len(nums); i++ {
            if _, ok := marked[i]; ok {
                continue
            }
            minIndex = findNondecreasing(i, &nums, &marked)
            if minIndex >= 0 {
                foundDecreasing = true
                marked[nums[i]] = true
                i++
            }
        }
        steps++
        isNondecreasing = !foundDecreasing
    }
    return steps
}

func findNondecreasing(i int, nums *[]int, marked *map[int]bool) int {
    for j := i+1; j < len(*nums); j++ {
        if (*marked)[(*nums)[j]] {
            continue
        } else if (*nums)[j] > (*nums)[i] {
            return -1
        } else if (*nums)[j] < (*nums)[i] {
            return j
        }
    }
    return -1
}
