package node

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/sign"
)

type Keypair struct {
	PublicKey  [32]byte `json:"public_key"`
	PrivateKey [64]byte `json:"private_key"`
}

type Keychain struct {
	MainKey   Keypair `json:"main_key"`
	BackupKey Keypair `json:"backup_key"`
}

func GenerateKeychain() (*Keychain, error) {
	keychain := &Keychain{}

	mainPublicKey, mainPrivateKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	backupPublicKey, backupPrivateKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	copy(keychain.MainKey.PublicKey[:], mainPublicKey[:])
	copy(keychain.MainKey.PrivateKey[:], mainPrivateKey[:])

	copy(keychain.BackupKey.PublicKey[:], backupPublicKey[:])
	copy(keychain.BackupKey.PrivateKey[:], backupPrivateKey[:])

	return keychain, nil
}
