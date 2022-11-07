# Go Currying

Functions in this go module implement curring. This module depends on Go's new generics language feature, so Go 1.18 or newer is required.

## Installation


To install go\_curring module, you can use the below command

```sh
go get -u github.com/kandu/go_curring
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

    hi(3)
    HI(2)
}
```

will print out

```
hi
hi
hi
HI
HI
```
