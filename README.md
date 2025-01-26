# go education
first programm 
```
package main


import "fmt"


func main() {
    fmt.Println("Hello, Go!")
}
```

```
rustam@rustam-zenbook:~/go-education$ nano main.go
rustam@rustam-zenbook:~/go-education$ go run main.go 
Hello, Go!
rustam@rustam-zenbook:~/go-education$ ls -la
total 12
drwxrwxr-x  2 rustam rustam 4096 Jan 26 21:57 .
drwxr-x--- 40 rustam rustam 4096 Jan 26 21:57 ..
-rw-rw-r--  1 rustam rustam   76 Jan 26 21:57 main.go
rustam@rustam-zenbook:~/go-education$ go build main.go 
rustam@rustam-zenbook:~/go-education$ ls
main  main.go
rustam@rustam-zenbook:~/go-education$ ./main 
Hello, Go!
```
