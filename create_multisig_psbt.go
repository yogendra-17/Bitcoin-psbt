package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/btcutil"
)

// createMultisigTx builds a PSBT to lock funds into a 2-of-2 multisig address.
func createMultisigTx() *wire.MsgTx {
	// Placeholder public keys for Alice and Bob
	alicePubKeyHex := "" // replace with Alice's public key (hex)
	bobPubKeyHex := ""
    // replace with Bob's public key (hex)
	
	// Placeholder for Address Z (fee recipient)
	feeAddressStr := "" // replace with Address Z in testnet format
	
	// Placeholder for UTXO hash (transaction ID for Alice's input)
	utxoTxID := "" // replace with actual UTXO transaction ID

	// Check for required values
	if alicePubKeyHex == "" || bobPubKeyHex == "" {
		log.Fatalf("Error: Public keys for Alice and Bob must be set.")
	}
	if feeAddressStr == "" {
		log.Fatalf("Error: Fee recipient address (Address Z) must be set.")
	}
	if utxoTxID == "" {
		log.Fatalf("Error: UTXO transaction ID for Alice's input must be set.")
	}

	// Decode Alice's and Bob's public keys from hex
	alicePubKey, err := hex.DecodeString(alicePubKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode Alice's public key: %v", err)
	}
	bobPubKey, err := hex.DecodeString(bobPubKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode Bob's public key: %v", err)
	}

	// Decode Address Z (fee recipient)
	feeAddress, err := btcutil.DecodeAddress(feeAddressStr, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatalf("Failed to decode fee recipient address (Address Z): %v", err)
	}

	fmt.Println("Building 2-of-2 multisig transaction...")

	//Create the 2-of-2 multisig redeem script
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_2).
		AddData(alicePubKey).AddData(bobPubKey).AddOp(txscript.OP_2).AddOp(txscript.OP_CHECKMULTISIG).Script()
	if err != nil {
		log.Fatalf("Failed to create redeem script: %v", err)
	}
	fmt.Printf("Multisig redeem script: %x\n", redeemScript)

	// Set the UTXO value and calculate outputs
	utxoValue := int64(10000) // Placeholder value in satoshis
	multisigOutput := utxoValue * 99 / 100 // 99% to multisig
	feeOutput := utxoValue * 1 / 100       // 1% to Address Z

	// Create PSBT with Alice’s UTXO as input
	tx := wire.NewMsgTx(wire.TxVersion)
	utxoHash, err := chainhash.NewHashFromStr(utxoTxID)
	if err != nil {
		log.Fatalf("Failed to parse UTXO hash: %v", err)
	}
	tx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(utxoHash, 0), nil, nil)) // Alice's UTXO
	fmt.Println("Added Alice's UTXO as input.")

	// Add 99% output to the multisig address
	multisigAddr, err := btcutil.NewAddressScriptHash(redeemScript, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatalf("Failed to create multisig address: %v", err)
	}
	scriptPubKey, err := txscript.PayToAddrScript(multisigAddr)
	if err != nil {
		log.Fatalf("Failed to create script for multisig address: %v", err)
	}
	tx.AddTxOut(wire.NewTxOut(multisigOutput, scriptPubKey))
	fmt.Printf("Added multisig output: %d satoshis\n", multisigOutput)

	// Add 1% output to Address Z
	feeScript, err := txscript.PayToAddrScript(feeAddress)
	if err != nil {
		log.Fatalf("Failed to create script for fee address: %v", err)
	}
	tx.AddTxOut(wire.NewTxOut(feeOutput, feeScript))
	fmt.Printf("Added fee output to Address Z: %d satoshis\n", feeOutput)

	fmt.Println("Generated PSBT for locking funds into multisig address.")
	fmt.Println("Transaction ID:", tx.TxHash().String())
	return tx
}
