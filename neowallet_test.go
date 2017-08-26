package neowallet

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestGenKey(t *testing.T) {
	privateKey := "0C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D"
	wallet, _ := GeneratePublicKeyFromPrivateKey(privateKey)

	log.Printf("%+v", wallet)
}

func TestGenFromWIF(t *testing.T) {
	wif := "KzULqzStT2tseGnqogXnTLG5NCT1YXa3F9Wp1Kdv9xMxFhvV6H2A"
	wallet, err := GenerateFromWIF(wif)
	if err != nil {
		log.Printf("%+v", err)
		t.Fail()
	}

	log.Printf("private key %+v", hex.EncodeToString(wallet.PrivateKey))
	log.Printf("public key %+v (%d)", hex.EncodeToString(wallet.PublicKey), len(wallet.PublicKey))
	log.Printf("wallet%+v", wallet)
	log.Printf("wallet address %+v %d", wallet.Address, len(wallet.Address))
}
