# Go Currying

Functions in this go module implement curring and uncurring. This module depends on Go's new generics language feature, so Go 1.18 or newer is required.

## Installation


To install go\_curring module, you can use the below command

```sh
go get -u github.com/kandu/go_currying@0.1
```

## Quick start

``` go
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
```

will print out

```
hi
hi
hi
hi
HI
HI
HI
UNCURRIED
UNCURRIED
rearranged
rearranged
```
