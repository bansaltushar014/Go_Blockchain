package network

type Transport interface {
	Consume() <-chan RPC
	Connect(*LocalTransport)
	SendMessage(string, string) error
	Broadcast([]byte) error
	GetAddress() string
}
