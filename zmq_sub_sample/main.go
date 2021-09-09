package main

import (
	"fmt"
	"log"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatal(err)
	}
	defer subscriber.Close()

	if len(os.Getenv("PUB_ADDRESS")) == 0 {
		subscriber.Connect("tcp://localhost:64230")
	} else {
		subscriber.Connect(fmt.Sprintf("tcp://%s:64230", os.Getenv("PUB_ADDRESS")))
	}

	subscriber.SetSubscribe("")

	for {
		msg, err := subscriber.Recv(0)
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Println(msg)
	}
}
