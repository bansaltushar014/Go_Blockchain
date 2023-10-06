package main

import (
	"time"

	"github.com/bansaltushar014/go-blockchain-l2/network"
)

func main() {
	trLocal := network.NewLockTransport("Local")
	trRemote := network.NewLockTransport("Remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	// addr := trLocal.GetAddress()
	go func() {
		for {
			trRemote.SendMessage(trLocal.GetAddress(), "Hello!")
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
