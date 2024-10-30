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
	fmt.Printf("Private Key Bytes: %x\n", privKeyBytes)
	return privKey, nil
}

func SignMultisigPSBT(psbt *wire.MsgTx) {
	// Replace these placeholders with the actual hex-encoded private keys
	alicePrivKeyHex := "" // placeholder,replace with alice's private key
	bobPrivKeyHex := ""     // placeholder, replace with bob's private key

	if alicePrivKeyHex == "" || bobPrivKeyHex == "" {
		log.Fatalf("Error: Private keys for Alice or Bob are not set. Please replace placeholders with actual keys.")
	}

	alicePrivKey, err := hexToPrivateKey(alicePrivKeyHex)
	if err != nil {
		log.Fatalf("Error decoding Alice's private key: %v", err)
	}
	bobPrivKey, err := hexToPrivateKey(bobPrivKeyHex)
	if err != nil {
		log.Fatalf("Error decoding Bob's private key: %v", err)
	}

	// Redeem script for Alice and Bob's 2-of-2 multisig
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_2).
		AddData(alicePrivKey.PubKey().SerializeCompressed()).
		AddData(bobPrivKey.PubKey().SerializeCompressed()).
		AddOp(txscript.OP_2).AddOp(txscript.OP_CHECKMULTISIG).Script()
	if err != nil {
		log.Fatalf("Error creating redeem script: %v", err)
	}

	// Sign each input
	for i, txIn := range psbt.TxIn {
		
		sigAlice, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, alicePrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Alice's signature for input %d: %v", i, err)
		}

		sigBob, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, bobPrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Bob's signature for input %d: %v", i, err)
		}

		fmt.Printf("Alice's Signature for input %d: %x\n", i, sigAlice)
		fmt.Printf("Bob's Signature for input %d: %x\n", i, sigBob)

		// signature script with OP_FALSE, Alice's and Bob's signatures, and the redeem script
		sigScriptBuilder := txscript.NewScriptBuilder()
		sigScriptBuilder.AddOp(txscript.OP_FALSE) // OP_FALSE for CHECKMULTISIG bug
		sigScriptBuilder.AddData(sigAlice)
		sigScriptBuilder.AddData(sigBob)
		sigScriptBuilder.AddData(redeemScript)
		sigScript, err := sigScriptBuilder.Script()
		if err != nil {
			log.Fatalf("Failed to create signature script for input %d: %v", i, err)
		}
		txIn.SignatureScript = sigScript
		fmt.Printf("Signature script: %x\n", sigScript)
	}
	
	fmt.Println("PSBT signed successfully by both Alice and Bob.")
}
