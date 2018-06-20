package main

import (
	"log"

	"storj.io/storj/pkg/netstate"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/process"
)

func main() {
	err := process.Main(&overlay.Service{}, &netstate.Service{})

	if err != nil {
		log.Fatal(err)
	}
}