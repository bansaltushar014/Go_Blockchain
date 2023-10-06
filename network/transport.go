package network

type RPC struct {
	netAddress string
	payload    string
}

type Transport interface {
	Consume() <-chan RPC
	Connect(*LocalTransport)
	SendMessage(string, string) error
	GetAddress() string
}
