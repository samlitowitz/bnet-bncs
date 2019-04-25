package main

import (
	"fmt"
	"log"

	"github.com/samlitowitz/bnet-bncs/pkg/bncs"
	"github.com/samlitowitz/bnet-bncs/pkg/bncs/server"
	"github.com/samlitowitz/bnet-encoding/pkg/encoding/bnet"
)

func main() {
	data := []byte{0xff, byte(bncs.SidPing), 0x08, 0x00, 0x01, 0x00, 0x00, 0x00}

	var header bncs.Header
	err := bnet.Unmarshal(data[:4], &header)
	if err != nil {
		log.Fatal(err)
	}

	var ping server.Ping
	err = bnet.Unmarshal(data[4:], &ping)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n%+v\n", header, ping)
}
