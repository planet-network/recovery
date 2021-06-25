package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/planet-platform/recovery/node"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
)

const (
	password = "example_password123"
)

func store(cmd *cobra.Command, args []string) {
	fmt.Println("-> generating keychain")
	keychain, err := node.GenerateKeychain()
	if err != nil {
		fail(err)
	}

	fmt.Println("-> serializing keychain")
	data, err := json.Marshal(keychain)
	if err != nil {
		fail(err)
	}

	fmt.Println("-> generating 8 byte salt")
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		fail(err)
	}

	fmt.Printf("-> creating derived key using pbkdf2 with sha512 using %q as password\n", password)
	dk := pbkdf2.Key([]byte(password), salt, 4096, 32, sha512.New)
	dkKey32 := [32]byte{}
	copy(dkKey32[:], dk[:])

	fmt.Println("-> encrypting keychain")
	encryptedKeychain, err := Encrypt(data, dkKey32)
	if err != nil {
		fail(err)
	}

	fmt.Printf("-> storing keychain to %q\n", args[0])
	if err := ioutil.WriteFile(args[0], encryptedKeychain, 0600); err != nil {
		fail(err)
	}

}

// Encrypt encrypts data using symmetric algorithm
func Encrypt(data []byte, key [32]byte) ([]byte, error) {
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, err
	}

	return secretbox.Seal(nonce[:], data, &nonce, (*[32]byte)(&key)), nil
}
