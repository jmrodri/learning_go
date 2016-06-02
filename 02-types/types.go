package main

import "fmt"
import "time"

type User struct {
    fname string
    lname string
    age int
    birthdate time.Time
}

func main() {
    fmt.Println("This is a string")
    fmt.Println("Con" + "catenated string")
    var str = "String variable"
    fmt.Println(str)

    user := User{}
    user.fname = "John"
    user.lname = "Doe"
    fmt.Println(user)
}
