package network

import (
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
	// fmt.Println(<-rpc)
	assert.Equal(t, (<-rpc).Payload, "Hello Buddy!")
}

func TestBroadcast(t *testing.T) {
	tra := NewLockTransport("A")
	trb := NewLockTransport("B")
	trc := NewLockTransport("C")

	tra.Connect(trb)
	tra.Connect(trc)

	msg := []byte("foo")
	assert.Nil(t, tra.Broadcast(msg))

	rpcb := <-trb.Consume()
	// b, err := ioutil.ReadAll(rpcb.Payload)
	// assert.Nil(t, err)
	assert.Equal(t, rpcb.Payload, string(msg))

	rpcC := <-trc.Consume()
	// b, err = ioutil.ReadAll(rpcC.Payload)
	// assert.Nil(t, err)
	assert.Equal(t, rpcC.Payload, string(msg))
}
