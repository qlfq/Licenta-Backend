package data

import (
	"crypto/rsa"
	gormcrypto "github.com/pkasila/gorm-crypto"
	"github.com/pkasila/gorm-crypto/algorithms"
	"github.com/pkasila/gorm-crypto/helpers"
	"github.com/pkasila/gorm-crypto/serialization"
	"io/ioutil"
	"os"
)

func InitializeCrypto() {
	var privateKey *rsa.PrivateKey
	var publicKey *rsa.PublicKey

	// Different behaviour if there is already a PEM with private key or not
	if _, err := os.Stat("private_key.pem"); os.IsNotExist(err) {
		// Do this if there is no PEM file

		// Generate key pairs public/private
		privateKey, publicKey, err = helpers.RSAGenerateKeyPair(4096)

		if err != nil {
			panic(err)
		}

		// Store it
		privateKeyBytes := helpers.RSAPrivateKeyToBytes(privateKey)
		err := ioutil.WriteFile("private_key.pem", privateKeyBytes, 0644)

		if err != nil {
			panic(err)
		}
	} else {
		// There is the PEM file
		bytes, err := ioutil.ReadFile("private_key.pem")

		if err != nil {
			panic(err)
		}

		// Bytes to private key
		privateKey, err = helpers.RSABytesToPrivateKey(bytes)

		if err != nil {
			panic(err)
		}

		publicKey = &privateKey.PublicKey
	}

	// Use privateKey and publicKey to initialize gormcrypto
	gormcrypto.Init(algorithms.NewRSA(privateKey, publicKey), serialization.NewJSON())
}
