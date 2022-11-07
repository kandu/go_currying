package main

import . "github.com/kandu/go_currying/curry0"
import "strings"

func main() {
    var print_str_upcase_n_times= func(str string, upcase bool, n int) {
        for i := 0; i < n; i++ {
            if upcase {
                println(strings.ToUpper(str))
            } else {
                println(str)
            }
        }
    }

    curried:= Curry3(print_str_upcase_n_times)

    hi:= curried("hi")(false)
    HI:= curried("hi")(true)

    hi(4)
    HI(3)

    uncurried:= Uncurry3(curried)
    rearranged:= Curry2(func(str string, n int) {
        uncurried(str, false, n)
    })

    uncurried("uncurried", true, 2)
    rearranged("rearranged")(2)
}
