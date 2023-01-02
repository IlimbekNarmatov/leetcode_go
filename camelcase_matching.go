func camelMatch(queries []string, pattern string) []bool {
    result := make([]bool, len(queries))
    for i := 0; i < len(queries); i++ {
        result[i] = isQueryCamelMatch(queries[i], pattern)
    }
    return result
}

func isQueryCamelMatch(query, pattern string) bool {
    var p int
    for i := 0; i < len(query); i++{
        qu := unicode.IsUpper(rune(query[i])) 
        if p >= len(pattern) {
            if qu {
                return false
            } 
            continue
        } 
        pu := unicode.IsUpper(rune(pattern[p]))
        if qu && pu {
            if query[i] != pattern[p] {
                return false
            }
            p++
        } else if qu && !pu {
            return false
        } else if !qu && !pu {
            if query[i] == pattern[p] {
                p++
            }
        } 
    }
    return p == len(pattern)
}
