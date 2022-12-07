package main

import "strings"

func intToRoman(num int) string {
    integers :=  [...]int{1000, 900 , 500, 400 , 100,   90, 50,    40, 10 ,    9,  5 , 4   , 1}
    romans := [...]string{"M" , "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV","I",}
    var i int
    var builder strings.Builder
    for ;num > 0; {
        if num - integers[i] < 0 {
            i++
            continue
        }
        builder.WriteString(romans[i])
        num -= integers[i]
    }
    return builder.String()
}
