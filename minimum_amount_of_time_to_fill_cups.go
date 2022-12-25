func fillCups(amount []int) int {    
    var c int
    var isAllFilled bool
    for ; !isAllFilled ; {
        isAllFilled = amount[0] == 0 && amount[1] == 0 && amount[2] == 0
        if isAllFilled {
            return c
        }
        var filled int
        if c % 2 == 0 && amount[0] > 0 && amount[1] > 0 {
            amount[0]--
            amount[1]--
        } else if c % 2 != 0 && amount[1] > 0 && amount[2] > 0 {
            amount[1]--
            amount[2]--
        } else {
            if filled < 3 && amount[0] > 0 {
                amount[0]--
                filled++
            }
            if filled < 3 && amount[1] > 0 {
                amount[1]--
                filled++
            } 
            if filled < 3 && amount[2] > 0 {
                amount[2]--
                filled++
            } 

        }
        c++
    }
    return c
}
