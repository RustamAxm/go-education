package main

import (
    "fmt"
    "github.com/golang/protobuf/proto"
)

func main() {

    // using the profo created struct
    p := &Application{ Payload: 10000, }

    // Serializing the struct and assigning it to body
    body, _ := proto.Marshal(p)

    // De-serializing the body and saving it to p1 for testing
    p1 := &Application{}
    _ = proto.Unmarshal(body, p1)

    fmt.Println("Original struct loaded from proto file:", p)
    fmt.Println("Marshalled proto data: ", body)
    fmt.Println("Unmarshalled struct: ", p1)
}
