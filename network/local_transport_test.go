package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnections(t *testing.T) {
	A := NewLockTransport("A")
	B := NewLockTransport("B")

	A.Connect(B)
	B.Connect(A)

	assert.Equal(t, A, B.peers[A.address])
	assert.Equal(t, A.address, B.peers[A.address].address)
}

func TestSendingMessage(t *testing.T) {
	A := NewLockTransport("A")
	B := NewLockTransport("B")

	A.Connect(B)
	B.Connect(A)

	A.SendMessage(B.address, "Hello Buddy!")
	rpc := B.Consume()
	fmt.Println(<-rpc)
	assert.Equal(t, (<-rpc).payload, "Hello Buddy!")
}
