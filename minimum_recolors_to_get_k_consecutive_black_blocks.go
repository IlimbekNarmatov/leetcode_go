func minimumRecolors(blocks string, k int) int {
    m := -1
    var bc int
    var s int
    var e int
    b := 'B'
    for ; e < len(blocks); {
        if e - s >= k {
            if rune(blocks[s]) == b {
                bc--
            }
            s++
            continue
        }
        if rune(blocks[e]) == b {
            bc++
        }
        if e - s + 1 < k {
            e++
            continue
        }
        c := e - s - bc + 1
        if m == -1 || m > c {
            m = c
        }
        e++
    }
    return m
}
