package network

import (
	"fmt"
	"time"

	core "github.com/bansaltushar014/go-blockchain-l2/core"
	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	RPCDecodeFunc RPCDecodeFunc
	RPCProcessor  RPCProcessor
	Transports    []Transport
	BlockTime     time.Duration
	PrivateKey    *cryptoX.PrivateKey
}

type Server struct {
	ServerOpts
	mempool     *TxPool
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}
	if opts.RPCDecodeFunc == nil {
		opts.RPCDecodeFunc = DefaultRPCDecodeFunc
	}

	s := &Server{
		ServerOpts:  opts,
		mempool:     NewTxPool(),
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}

	if s.RPCProcessor == nil {
		s.RPCProcessor = s
	}

	// s.ServerOpts = opts

	return s
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.BlockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				logrus.Error(err)
			}

			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
				logrus.Error(err)
			}
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do stuff every x seconds")
			if s.isValidator {
				s.createNewBlock()
			}
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) ProcessMessage(msg *DecodedMessage) error {
	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.ProcessTransaction(t)
	}

	return nil
}

func (s *Server) broadcast(payload []byte) error {
	for _, tr := range s.Transports {
		if err := tr.Broadcast(payload); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) ProcessTransaction(tx *core.Transaction) error {
	hash := tx.Hash(tx.Data)

	if s.mempool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already in mempool")

		return nil
	}

	// if !tx.Verify() {
	// 	return errors.New("Verification Failed!")
	// }

	tx.SetFirstSeen(time.Now())

	logrus.WithFields(logrus.Fields{
		"hash":           hash,
		"mempool length": s.mempool.Len(),
	}).Info("adding new tx to the mempool")

	go s.broadcastTx(tx)

	return s.mempool.Add(tx)
}

func (s *Server) broadcastTx(tx *core.Transaction) error {
	// buf := &bytes.Buffer{}
	// if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
	// 	return err
	// }

	msg := NewMessage(MessageTypeTx, tx.Data)

	return s.broadcast(msg.Bytes())
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating a new block")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
