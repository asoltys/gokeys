package main

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/btcec"

	zmq "github.com/pebbe/zmq4"
	"github.com/vulpemventures/go-elements/block"
	"github.com/vulpemventures/go-elements/network"
	"github.com/vulpemventures/go-elements/payment"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.SUB)
	s.Connect("tcp://localhost:18606")
	s.SetSubscribe("rawblock")

  s.Recv(0)
	for {
		msg, err := s.Recv(0)
		if err != nil {
			panic(err)
		}
		data := []byte(msg)
		x := bytes.NewBuffer(data)
		b, err := block.NewFromBuffer(x)
		if err != nil {
			fmt.Println(err)
			continue
		}
    for _, tx := range b.TransactionsData.Transactions {
      fmt.Println("TXN");
      for _, output := range tx.Outputs {
        fmt.Println("OUTPUT");
        if output.Asset != nil && len(output.Asset) > 0 {
          fmt.Println("HAS ASSET");

          blindPrivkey, _ := btcec.NewPrivateKey(btcec.S256())
          blindPubkey := blindPrivkey.PubKey()
          p, err := payment.FromScript(output.Script, &network.Regtest, blindPubkey)
          if err != nil {
            fmt.Println(err)
            continue
          }
          address, _ := p.WitnessPubKeyHash()
          fmt.Println(address)
        }
      }
    }
	}
}
