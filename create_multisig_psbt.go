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

// CreateMultisigPSBT generates a PSBT for locking funds into a 2-of-2 multisig address.
func CreateMultisigPSBT() *wire.MsgTx {
	// Public keys for Alice and Bob
	alicePubKey, _ := hex.DecodeString("0356404a1d39773a23c6fcf28e9eb248e6b5a9c5bd8ff9211f928a178f3481b7ce") // replace with actual public key
	bobPubKey, _ := hex.DecodeString("023d33218fe35eb4aec6fe81c8acb58dbc51f98708a664c9702a16cfb883f7617e") // replace with actual public key

	// Address Z 
	addressZ, _ := btcutil.DecodeAddress("bcrt1qzefpec3w9f5jr743pxt3x8q0p5m9p9r85gryes", &chaincfg.RegressionNetParams)

	fmt.Println("Creating a 2-of-2 multisig PSBT...")

	// 1. Create a 2-of-2 multisig redeem script for Alice and Bob
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_2).
		AddData(alicePubKey).AddData(bobPubKey).AddOp(txscript.OP_2).AddOp(txscript.OP_CHECKMULTISIG).Script()
	if err != nil {
		log.Fatalf("Failed to create multisig redeem script: %v", err)
	}
	fmt.Printf("Multisig redeem script: %x\n", redeemScript)

	// 2. Assume Alice's UTXO value and calculate outputs
	utxoValue := int64(10000) //placeholder for utxo value
	multisigOutput := int64(utxoValue * 99 / 100) // 99% of UTXO
	feeRecipientOutput := int64(utxoValue * 1 / 100) // 1% of UTXO to Address Z

	// 3. Create PSBT with Alice’s UTXO as input
	psbt := wire.NewMsgTx(wire.TxVersion)
	utxoHash, _ := chainhash.NewHashFromStr("5e0ba95a787fe54b1ebb2a1cda55186152e63421bc0adaa6e0d052c56f2d8dfc") // Replace with actual UTXO hash
	psbt.AddTxIn(wire.NewTxIn(wire.NewOutPoint(utxoHash, 0), nil, nil)) // UTXO from Alice
	fmt.Println("Added Alice's UTXO as input.")

	// 4. Add output 1: 99% to the multisig address
	multisigAddress, err := btcutil.NewAddressScriptHash(redeemScript, &chaincfg.RegressionNetParams)
	println(multisigAddress)
	if err != nil {
		log.Fatalf("Failed to create multisig address: %v", err)
	}
	scriptPubKey, err := txscript.PayToAddrScript(multisigAddress)
	if err != nil {
		log.Fatalf("Failed to create script for multisig address: %v", err)
	}
	psbt.AddTxOut(wire.NewTxOut(multisigOutput, scriptPubKey))
	fmt.Printf("Added multisig output with value: %d satoshis\n", multisigOutput)

	// 5. Add output 2: 1% to Address Z
	scriptAddrZ, err := txscript.PayToAddrScript(addressZ)
	if err != nil {
		log.Fatalf("Failed to create script for Address Z: %v", err)
	}
	psbt.AddTxOut(wire.NewTxOut(feeRecipientOutput, scriptAddrZ))
	fmt.Printf("Added fee output to Address Z with value: %d satoshis\n", feeRecipientOutput)

	fmt.Println("Generated PSBT for locking funds into multisig address.")
	
	return psbt
}
