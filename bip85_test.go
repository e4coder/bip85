package bip85_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/e4coder/bip85"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip32"
)

const MASTER_BIP32_ROOT_KEY = "xprv9s21ZrQH143K2LBWUUQRFXhucrQqBpKdRRxNVq2zBqsx8HVqFk2uYo8kmbaLLHRdqtQpUm98uKfu3vca1LqdGhUtyoFnCNkfmXRyPXLjbKb"

func TestEntropyFromKey(t *testing.T) {
	mrootkey, err := bip32.B58Deserialize(MASTER_BIP32_ROOT_KEY)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	pathStr1 := "m/83696968'/0'/0'"
	path1, err := accounts.ParseDerivationPath(pathStr1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	// derive the child key
	var currKey1 *bip32.Key = mrootkey

	for _, childIdX := range path1 {
		k, err := currKey1.NewChildKey(childIdX)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		currKey1 = k
	}

	childKey1 := currKey1

	t.Log(hex.EncodeToString(childKey1.Key))
	if strings.Compare(hex.EncodeToString(childKey1.Key), "cca20ccb0e9a90feb0912870c3323b24874b0ca3d8018c4b96d0b97c0e82ded0") != 0 {
		t.Log("childKey1 does not match")
		t.FailNow()
	}

	entropyNew1, err := bip85.EntropyFromKey(childKey1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if strings.Compare(hex.EncodeToString(entropyNew1), "efecfbccffea313214232d29e71563d941229afb4338c21f9517c41aaa0d16f00b83d2a09ef747e7a64e8e2bd5a14869e693da66ce94ac2da570ab7ee48618f7") != 0 {
		t.Log("entropyNew1 does not match")
		t.FailNow()
	}

}

func TestNewBip39FromEntropy(t *testing.T) {
	mrootkey, err := bip32.B58Deserialize(MASTER_BIP32_ROOT_KEY)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	// pathStr := "m/83696968'/0'/0'"
	pathStr := "m/83696968'/39'/0'/12'/0'"
	path, err := accounts.ParseDerivationPath(pathStr)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	// derive the child key
	var currKey *bip32.Key = mrootkey

	for _, childIdX := range path {
		k, err := currKey.NewChildKey(childIdX)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		currKey = k
	}

	childKey := currKey

	t.Log(hex.EncodeToString(childKey.Key))
	entropyNew, err := bip85.EntropyFromKey(childKey)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(hex.EncodeToString(entropyNew[:16]))

	mnemonic, err := bip85.NewBip39FromEntropy(entropyNew)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(mnemonic)

	if strings.Compare(mnemonic, "girl mad pet galaxy egg matter matrix prison refuse sense ordinary nose") != 0 {
		t.Log("Mnemonic does not match")
		t.FailNow()
	}
}
