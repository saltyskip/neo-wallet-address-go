package neowallet

import (
	"encoding/hex"

	"github.com/apisit/btckeygenie/btckey"
)

type Wallet struct {
	PublicKey       []byte
	PrivateKey      []byte
	Address         string
	WIF             string
	HashedSignature []byte
}

func hex2bytes(hexstring string) (b []byte) {
	b, _ = hex.DecodeString(hexstring)
	return b
}

func GeneratePublicKeyFromPrivateKey(privateKey string) (*Wallet, error) {
	pb := hex2bytes(privateKey)
	var priv btckey.PrivateKey
	err := priv.FromBytes(pb)
	if err != nil {
		return &Wallet{}, err
	}
	wallet := &Wallet{
		PublicKey:       priv.PublicKey.ToBytes(),
		PrivateKey:      priv.ToBytes(),
		Address:         priv.ToNeoAddress(),
		WIF:             priv.ToWIF(),
		HashedSignature: priv.ToNeoSignature(),
	}
	return wallet, nil
}

func GenerateFromWIF(wif string) (*Wallet, error) {
	var priv btckey.PrivateKey
	err := priv.FromWIF(wif)
	if err != nil {
		return &Wallet{}, err
	}

	wallet := &Wallet{
		PublicKey:       priv.PublicKey.ToBytes(),
		PrivateKey:      priv.ToBytes(),
		Address:         priv.ToNeoAddress(),
		WIF:             priv.ToWIF(),
		HashedSignature: priv.ToNeoSignature(),
	}
	return wallet, nil
}
