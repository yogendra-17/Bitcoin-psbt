package main

import (

	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
)

// GenerateP2SHAddress creates a testnet P2SH address from a given redeem script
func GenerateP2SHAddress(redeemScript []byte) (string, error) {
    // Hash the redeem script using hash160
    redeemHash := hash160(redeemScript)

    // Add the version byte for testnet (0xC4) at the start
    versionedHash := append([]byte{0xC4}, redeemHash...)

    // Double SHA-256 for checksum calculation
    checksum := dblSha256(versionedHash)[:4]

    // Append the checksum to versionedHash
    fullHash := append(versionedHash, checksum...)

    // Base58 encode the full hash to get the P2SH address
    p2shAddress := base58.Encode(fullHash)
    return p2shAddress, nil
}

func mainP2SH() {
	// Assuming we have the redeem script from our 2-of-2 multisig setup
	redeemScript, err := BuildMultiSigRedeemScript(
		[]byte{0x03, 0x9a, 0xa0, 0x52, 0xf9, 0x44, 0xa1, 0x86, 0xbf, 0xca, 0x84, 0xfc, 0x7d, 0x90, 0x20, 0x41, 0xf7, 0x52, 0x6d, 0x96, 0xbc, 0x70, 0xee, 0x49, 0x40, 0xad, 0x74, 0x87, 0x69, 0xd5, 0x38, 0x4f, 0xf4},
		[]byte{0x03, 0x1d, 0x1b, 0xe6, 0x13, 0xb6, 0x28, 0x33, 0x0d, 0x42, 0x2f, 0x93, 0x88, 0x85, 0x4c, 0x2d, 0xc7, 0x54, 0x86, 0x1a, 0x3d, 0x19, 0x77, 0x3a, 0xa6, 0x85, 0xe4, 0x53, 0x6d, 0xd0, 0xd9, 0x54, 0x64},
	)
	if err != nil {
		log.Fatalf("Failed to create redeem script: %v", err)
	}

	// Generate the P2SH address
	address, err := GenerateP2SHAddress(redeemScript)
	if err != nil {
		log.Fatalf("Failed to create P2SH address: %v", err)
	}

	fmt.Println("MultiSig P2SH Address:", address)
}
