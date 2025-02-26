package main


import "fmt"
import "go-education/my_module"

func main() {
    fmt.Println("Hello, Go!")
    m := my_module.NewModule()
    m.Module_func()
}

