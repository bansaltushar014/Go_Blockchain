package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bansaltushar014/go-blockchain-l2/core"
	"github.com/bansaltushar014/go-blockchain-l2/network"
	"github.com/sirupsen/logrus"
)

func main() {
	trLocal := network.NewLockTransport("Local")
	trRemote := network.NewLockTransport("Remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	// addr := trLocal.GetAddress()
	go func() {
		for {
			// trRemote.SendMessage(trLocal.GetAddress(), "Hello!")
			if err := sendTransaction(trRemote, trLocal.GetAddress()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}

func sendTransaction(tr network.Transport, to string) error {
	// privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000000000)), 10))
	t1 := core.Trans{}
	tx := t1.CreateRandomTx(data)
	// tx.Sign(privKey)
	// buf := &bytes.Buffer{}
	// if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
	// 	return err
	// }
	// encoded := core.ExcodeTx(tx.Data)
	encoded := tx.Encoded(tx.Data)

	msg := network.NewMessage(network.MessageTypeTx, []byte(encoded))

	return tr.SendMessage(to, string(msg.Bytes()))
}
