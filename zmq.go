package main

import (
  "fmt"
  zmq "github.com/pebbe/zmq4"
)

func main() {
    zctx, _ := zmq.NewContext()

    s, _ := zctx.NewSocket(zmq.SUB)
    s.Connect("tcp://localhost:18606")
    s.SetSubscribe("rawblock")

    for {
        msg, err := s.Recv(0)
        if err != nil {
            panic(err)
        }
        fmt.Println(msg);
    }
}
