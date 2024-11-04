package main

import (
	"fmt"
	"log"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// createUnlockTx builds a PSBT to unlock funds from a multisig UTXO.
func createUnlockTx(multisigUTXO string) *wire.MsgTx {
	// Placeholder addresses for fee recipient (Address Z) and Alice’s address
	feeAddressStr := "tb1q4943lwh4ey0z95qh7ywgkxllnapu8sjct9nwvm"     // replace with actual Address Z (testnet format)
	aliceAddressStr := "tb1q8syrtgcwptvanrejzpdvuezwdy3dq9hgqq90qg"    // replace with Alice’s actual testnet address
	utxoValue := int64(1000) // placeholder UTXO value in satoshis

	// Check for required placeholders
	if feeAddressStr == "" || aliceAddressStr == "" || multisigUTXO == "" {
		log.Fatalf("Error: All placeholders (fee address, Alice's address, multisig UTXO hash) must be set.")
	}

	// Decode Address Z and Alice's address
	feeAddress, err := btcutil.DecodeAddress(feeAddressStr, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatalf("Failed to decode fee recipient address (Address Z): %v", err)
	}
	aliceAddress, err := btcutil.DecodeAddress(aliceAddressStr, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatalf("Failed to decode Alice's address: %v", err)
	}

	// Decode the multisig UTXO hash
	utxoHash, err := chainhash.NewHashFromStr(multisigUTXO)
	if err != nil {
		log.Fatalf("Failed to decode multisig UTXO hash: %v", err)
	}

	// Calculate output amounts
	feeAmount := utxoValue * 1 / 100           // 1% to Address Z as fee
	aliceAmount := utxoValue - feeAmount - 1000 // Alice's amount after subtracting fees

	// Create a new transaction
	tx := wire.NewMsgTx(wire.TxVersion)

	// Add multisig UTXO as input (replace 0 with actual output index if needed)
	tx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(utxoHash, 0), nil, nil))
	fmt.Println("Added multisig UTXO as input.")

	// Add output 1: 1% fee to Address Z
	feeScript, err := txscript.PayToAddrScript(feeAddress)
	if err != nil {
		log.Fatalf("Failed to create script for fee address: %v", err)
	}
	tx.AddTxOut(wire.NewTxOut(feeAmount, feeScript))
	fmt.Printf("Added fee output to Address Z: %d satoshis\n", feeAmount)

	// Add output 2: Remaining amount to Alice
	aliceScript, err := txscript.PayToAddrScript(aliceAddress)
	if err != nil {
		log.Fatalf("Failed to create script for Alice's address: %v", err)
	}
	tx.AddTxOut(wire.NewTxOut(aliceAmount, aliceScript))
	fmt.Printf("Added output to Alice: %d satoshis\n", aliceAmount)

	fmt.Println("Generated PSBT for unlocking multisig funds.")
	return tx
}

