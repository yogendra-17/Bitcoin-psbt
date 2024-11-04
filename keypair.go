package main

import (
	"encoding/hex"
	"log"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/btcutil"
)

// keypair generates a new private and public key pair and returns the corresponding testnet address
func keypair() (string, string, string) {
	// Generate a new private key
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Convert the private key to a hexadecimal string
	privateKeyHex := hex.EncodeToString(privateKey.Serialize())

	// Get the public key in compressed format
	publicKey := privateKey.PubKey()
	publicKeyHex := hex.EncodeToString(publicKey.SerializeCompressed())

	// Generate the testnet address (P2WPKH, Bech32 format)
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatalf("Failed to create testnet address: %v", err)
	}

	return privateKeyHex, publicKeyHex, address.EncodeAddress()
}
