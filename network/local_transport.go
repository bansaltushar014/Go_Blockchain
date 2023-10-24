package network

import (
	"fmt"
)

type LocalTransport struct {
	address   string
	consumeCh chan RPC
	peers     map[string]*LocalTransport
}

func NewLockTransport(add string) *LocalTransport {
	return &LocalTransport{
		address:   add,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[string]*LocalTransport),
	}
}

func (a *LocalTransport) GetAddress() string {
	return a.address
}

func (a *LocalTransport) Consume() <-chan RPC {
	return a.consumeCh
}

func (a *LocalTransport) Connect(b *LocalTransport) {
	a.peers[b.address] = b
}

func (a *LocalTransport) SendMessage(addr string, message string) error {
	peer, ok := a.peers[addr]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", a.address, addr)
	}

	peer.consumeCh <- RPC{
		From:    addr,
		Payload: message,
	}
	return nil
}

func (a *LocalTransport) Broadcast(payload []byte) error {
	for _, peer := range a.peers {
		if err := a.SendMessage(peer.GetAddress(), string(payload)); err != nil {
			return err
		}
	}
	return nil
}
