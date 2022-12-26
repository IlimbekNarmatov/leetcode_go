func countGoodSubstrings(s string) int {
    var r int
    for i := 0; i < len(s) - 2; i++ {
        if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
            r++
        }
    }
    return r
}
