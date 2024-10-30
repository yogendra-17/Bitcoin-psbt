// sign_multisig_psbt.go
package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/btcec/v2"
)

// hexToPrivateKey parses a hex-encoded private key string and returns a btcec private key.
func hexToPrivateKey(hexKey string) (*btcec.PrivateKey, error) {
	privKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}
	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes) // Ignore the public key
	return privKey, nil
}

// SignMultisigPSBT signs each input in the PSBT with both Alice's and Bob's private keys
func SignMultisigPSBT(psbt *wire.MsgTx) {
	// Hardcoded private keys for Alice and Bob
	alicePrivKeyHex := "00c410a3a4d84ba56873142f823e99bec321420c3265e2fc2db8f7671e9dc3ce5c"
	bobPrivKeyHex := "00bf1eedb8b718a18197578d6bfdc3a21c8f70cb624c4eaeab61b3c8c46e74cfa4"

	// Convert hex-encoded private keys to btcec.PrivateKey objects
	alicePrivKey, err := hexToPrivateKey(alicePrivKeyHex)
	if err != nil {
		log.Fatalf("Error decoding Alice's private key: %v", err)
	}
	bobPrivKey, err := hexToPrivateKey(bobPrivKeyHex)
	if err != nil {
		log.Fatalf("Error decoding Bob's private key: %v", err)
	}

	// Example redeem script (from CreateMultisigPSBT)
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_2).
		AddData(alicePrivKey.PubKey().SerializeCompressed()).
		AddData(bobPrivKey.PubKey().SerializeCompressed()).
		AddOp(txscript.OP_2).AddOp(txscript.OP_CHECKMULTISIG).Script()
	if err != nil {
		log.Fatalf("Error creating redeem script: %v", err)
	}

	// Iterate over each input and sign it with both Alice's and Bob's private keys
	for i, txIn := range psbt.TxIn {
		// Generate Alice's signature
		sigAlice, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, alicePrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Alice's signature for input %d: %v", i, err)
		}

		// Generate Bob's signature
		sigBob, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, bobPrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Bob's signature for input %d: %v", i, err)
		}

		// Construct the SignatureScript with OP_FALSE, Alice's and Bob's signatures, and the redeem script
		sigScriptBuilder := txscript.NewScriptBuilder()
		sigScriptBuilder.AddOp(txscript.OP_FALSE) // OP_FALSE for CHECKMULTISIG bug
		sigScriptBuilder.AddData(sigAlice)
		sigScriptBuilder.AddData(sigBob)
		sigScriptBuilder.AddData(redeemScript)
		sigScript, err := sigScriptBuilder.Script()
		if err != nil {
			log.Fatalf("Failed to create signature script for input %d: %v", i, err)
		}

		// Assign the signature script to the transaction input
		txIn.SignatureScript = sigScript
	}

	fmt.Println("PSBT signed successfully by both Alice and Bob.")
	fmt.Print(redeemScript)
}
