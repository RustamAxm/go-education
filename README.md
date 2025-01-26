# go education
## first programm 
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
## package

```
rustam@rustam-zenbook:~/go-education$ tree .
.
├── go.mod
├── main
├── main.go
├── my_module
│   └── my_module.go
└── README.md

2 directories, 5 files

```
imported func name start with UpperCase

Create mod file
```
go mod init <project name >
```
example
```
rustam@rustam-zenbook:~/go-education$ cat go.mod 
module go-education

go 1.23.5
```
after create files in module correct import in main
```
package main


import "fmt"
import "go-education/my_module"

func main() {
    fmt.Println("Hello, Go!")
    m := my_module.NewModule()
    m.Module_func()
}
```
after build created runnable bin
```
rustam@rustam-zenbook:~/go-education$ go build 
rustam@rustam-zenbook:~/go-education$ ls -la
total 2112
drwxrwxr-x  4 rustam rustam    4096 Jan 26 22:45 .
drwxr-x--- 41 rustam rustam    4096 Jan 26 22:44 ..
drwxrwxr-x  8 rustam rustam    4096 Jan 26 22:45 .git
-rwxrwxr-x  1 rustam rustam 2130991 Jan 26 22:45 go-education
-rw-rw-r--  1 rustam rustam      31 Jan 26 22:25 go.mod
-rw-rw-r--  1 rustam rustam     158 Jan 26 22:37 main.go
drwxrwxr-x  2 rustam rustam    4096 Jan 26 22:36 my_module
-rw-rw-r--  1 rustam rustam    1237 Jan 26 22:43 README.md
rustam@rustam-zenbook:~/go-education$ ./go-education 
Hello, Go!
in module

```
