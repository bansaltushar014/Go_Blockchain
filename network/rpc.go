package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/bansaltushar014/go-blockchain-l2/core"
	"github.com/sirupsen/logrus"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBock
)

type RPC struct {
	From    string
	Payload string
}

type Message struct {
	Header MessageType
	Data   []byte
}

type DecodedMessage struct {
	From string
	Data interface{}
}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

func (msg *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

type RPCDecodeFunc func(RPC) (*DecodedMessage, error)

func DefaultRPCDecodeFunc(rpc RPC) (*DecodedMessage, error) {
	msg := Message{}
	if err := gob.NewDecoder(strings.NewReader(rpc.Payload)).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to decode message from %s: %s", rpc.From, err)
	}

	logrus.WithFields(logrus.Fields{
		"from": rpc.From,
		"type": msg.Header,
	}).Debug("new incoming message")

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		// if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
		// 	return err
		// }
		// return h.p.ProcessTransaction(rpc.From, tx)
		return &DecodedMessage{
			From: rpc.From,
			Data: tx,
		}, nil
	default:
		return nil, fmt.Errorf("invalid message header %x", msg.Header)
	}
}

type RPCProcessor interface {
	ProcessMessage(*DecodedMessage) error
}
