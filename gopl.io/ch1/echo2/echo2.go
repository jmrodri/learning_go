package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    s, sep := "", ""
    for idx, arg := range os.Args[0:] {
        s += strconv.Itoa(idx) + sep + arg + "\n"
        sep = " "
    }
    fmt.Println(s)
}
