package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/tyler-smith/go-bip32"
	"github.com/vulpemventures/go-elements/network"
	"github.com/vulpemventures/go-elements/payment"
)

func main() {
	bip32.PrivateWalletVersion, _ = hex.DecodeString("04358394")
	bip32.PublicWalletVersion, _ = hex.DecodeString("043587cf")
	root, _ := bip32.B58Deserialize("tprv8ZgxMBicQKsPdM1ejJrR8D3qr3h7wCPvDLcWmrBb5Miyrh34j3XzgghV2bvS1baNbLveYFCi87jpTdD7EUkxoJYaE6mS982ac9EdcBJoqVG")

  // m/0'
	x, _ := root.NewChildKey(bip32.FirstHardenedChild)

  // m/0'/0'
	y, _ := x.NewChildKey(bip32.FirstHardenedChild)

  // m/0'/0'/91586'
	z, _ := y.NewChildKey(bip32.FirstHardenedChild + 91587)
	fmt.Println(hex.EncodeToString(z.Key));

	pubkey := z.PublicKey()

	blindPrivkey, _ := btcec.NewPrivateKey(btcec.S256())
	blindPubkey := blindPrivkey.PubKey()

	pk, _ := btcec.ParsePubKey(pubkey.Key, btcec.S256())
	p2wpkh := payment.FromPublicKey(pk, &network.Regtest, blindPubkey)
	address, _ := p2wpkh.PubKeyHash()
	fmt.Println(address)
}
