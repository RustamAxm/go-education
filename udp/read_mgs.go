package main
import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
   // ваш код
    conn, err := net.Dial("tcp", "127.0.0.1:8081")
    if err != nil {
        log.Println(err)
    }
    defer conn.Close()
    for i:=0 ; i < 3; i++ {
        message := make([]byte, 1024) // создадим буфер
        n, err := conn.Read(message)
        if err != nil {
            log.Println(err)
        }
        fmt.Println(strings.ToUpper(string(message[:n]))) 
    }

}
