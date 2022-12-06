package main

func maxArea(height []int) int {
    if len(height) <= 1 {
        return 0
    }
    var result int
    var start int
    end := len(height)-1
    for ; start < end; {
        area := area(start, end, height)
        if area > result {
            result = area
        }
        if height[start] < height[end] {
            start++
        } else {
            end--
        }
    }
    return result
}

func area(i, j int, height []int) int {
    min := min(height[i], height[j])
    return min * (j-i)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
