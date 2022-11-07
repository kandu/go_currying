package main

import "github.com/kandu/go_currying/curry0"
import "github.com/kandu/go_currying/curry"
import "fmt"
import "strings"

func main() {
    print_str_upcase_n_times:= func(str string, upcase bool, n int) {
        for i := 0; i < n; i++ {
            if upcase {
                println(strings.ToUpper(str))
            } else {
                println(str)
            }
        }
    }

    curried:= curry0.Curry3(print_str_upcase_n_times)

    hi:= curried("hi")(false)
    HI:= curried("hi")(true)

    hi(4)
    HI(3)

    uncurried:= curry0.Uncurry3(curried)
    rearranged:= curry0.Curry2(func(str string, n int) {
        uncurried(str, false, n)
    })

    uncurried("uncurried", true, 2)
    rearranged("rearranged")(2)

    println("\n---\n")

    checkLen:= func(length int, str string) bool {
        return length == len(str)
    }

    lenIsTen:= curry.Curry2(checkLen)(10)
    
    str:= "hi"
    fmt.Printf("the length of \"%s\" is 10, %v\n", str, lenIsTen(str))
    str= "0123456789"
    fmt.Printf("the length of \"%s\" is 10, %v\n", str, lenIsTen(str))
}
