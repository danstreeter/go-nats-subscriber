package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var version string = "0.0.1"

func main() {
	fmt.Println("Go NATS Subscriber", version)
	var natsServer string
	var subjectName string
	if len(os.Args) == 3 {
		natsServer = os.Args[1]
		subjectName = os.Args[2]
	} else {
		fmt.Print("Enter NATS server to connect to (Default: nats://demo.nats.io:4222): ")
		fmt.Scanln(&natsServer)
		if natsServer == "" {
			// fmt.Println("Using default: nats://demo.nats.io:4222")
			natsServer = "nats://demo.nats.io:4222"
		}

		fmt.Print("Enter NATS subject to subscribe to (Default: >): ")
		fmt.Scanln(&subjectName)
		if subjectName == "" {
			// fmt.Println("Using default: >")
			subjectName = ">"
		}
	}
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Println("Subscribing to", natsServer, "on subject:", subjectName)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(subjectName, ch)
	defer sub.Unsubscribe()

	for {
		// Wait for incoming messages
		req := <-ch
		timeNow := time.Now().Format("2006-01-02 15:04:05.000000")
		fmt.Printf("%s - Received message on '%s': %s\n", string(timeNow), string(req.Subject), string(req.Data))
	}

}
