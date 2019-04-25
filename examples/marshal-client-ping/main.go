package main

import (
	"fmt"
	"log"

	"github.com/samlitowitz/bnet-bncs/pkg/bncs"
	"github.com/samlitowitz/bnet-bncs/pkg/bncs/client"
	"github.com/samlitowitz/bnet-encoding/pkg/encoding/bnet"
)

func main() {
	buffer := make([]byte, 8)

	header := &bncs.Header{
		Fixed:     0xff,
		MessageID: bncs.SidPing,
		Length:    0x08,
	}

	data, err := bnet.Marshal(&header)
	if err != nil {
		log.Fatal(err)
	}

	if copy(buffer, data) != 4 {
		log.Fatal("Expected 4 bytes copied.")
	}

	ping := &client.Ping{PingValue: 0x01}
	data, err = bnet.Marshal(&ping)
	if err != nil {
		log.Fatal(err)
	}

	if copy(buffer[4:], data) != 4 {
		log.Fatal("Expected 4 bytes copied.")
	}

	fmt.Printf("%+v\n", buffer)
}
