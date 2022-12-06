package main

func lengthOfLongestSubstring(s string) int {
    var result int
    set := make(map[rune]bool)
    var start int
    var end int
    for ; end < len(s); {
        if v, ok := set[rune(s[end])]; v && ok {
            set[rune(s[start])] = false
            start++
            continue
        }
        set[rune(s[end])] = true
        end++
        if end - start > result {
            result = end - start
        }
    }
    return result
}
