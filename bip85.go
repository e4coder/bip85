package bip85

import (
	"crypto/hmac"
	"crypto/sha512"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func EntropyFromKey(key *bip32.Key) ([]byte, error) {
	hash := hmac.New(sha512.New, []byte("bip-entropy-from-k"))
	_, err := hash.Write(key.Key)
	if err != nil {
		return []byte(""), err
	}
	hashedKey := hash.Sum(nil)
	return hashedKey, nil
}

func NewBip39FromEntropy(entropy []byte) (string, error) {
	entropyBIP39 := entropy[:16]
	return bip39.NewMnemonic(entropyBIP39)
}
