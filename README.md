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

## test demo
### std tests
```
rustam@rustam-zenbook:~/go-education$ tree my_module/
my_module/
├── my_module.go
└── my_module_test.go

1 directory, 2 files
rustam@rustam-zenbook:~/go-education$ cat my_module/my_module_test.go 

package my_module

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	fmt.Println("in test")
	m := NewModule()
	if got := m.Module_func(); got != "in module" {
		t.Errorf("error in module test = %q", got)
	}
}

func TestSimpleErr(t *testing.T) {
        fmt.Println("in test")
        m := NewModule()
        m.Module_func()
        if got := m.Module_func(); got != "in module 2 " {
                t.Errorf("error in module test = %q", got)
        }
}

```
run tests 
```
rustam@rustam-zenbook:~/go-education$ go test -v go-education/my_module
=== RUN   TestSimple
in test
in module
--- PASS: TestSimple (0.00s)
=== RUN   TestSimpleErr
in test
in module
in module
    my_module_test.go:20: error in module test = "in module"
--- FAIL: TestSimpleErr (0.00s)
FAIL
FAIL	go-education/my_module	0.005s
FAIL

```
### gotestsum 
install 
```
go install -v  gotest.tools/gotestsum
```

```
rustam@rustam-zenbook:~/go-education$ gotestsum --junitfile unit-tests.xml
∅  .
∅  go_doc
∅  gorutine
✖  my_module (19ms)
∅  net_http
∅  proto_test
∅  udp
∅  udp/client
∅  udp/server

=== Failed
=== FAIL: my_module TestSimpleErr (0.00s)
in test
in module
in module
    my_module_test.go:20: error in module test = "in module"

DONE 2 tests, 1 failure in 2.715s
rustam@rustam-zenbook:~/go-education$ cat unit-tests.xml 
<?xml version="1.0" encoding="UTF-8"?>
<testsuites tests="2" failures="1" errors="0" time="2.806015">
	<testsuite tests="0" failures="0" time="0.000000" name="go-education" timestamp="2025-02-26T20:58:23+03:00">

```


