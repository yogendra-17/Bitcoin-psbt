package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/txscript"
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

// SignMultisigPSBT signs the PSBT with Alice's and Bob's private keys and returns the raw signed transaction.
func SignMultisigPSBT(psbt *wire.MsgTx) string {
	alicePrivKeyHex := "" // Placeholder: Replace with Alice's private key
	bobPrivKeyHex := ""  // Placeholder: Replace with Bob's private key

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

	for i, txIn := range psbt.TxIn {
		sigAlice, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, alicePrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Alice's signature for input %d: %v", i, err)
		}

		sigBob, err := txscript.RawTxInSignature(psbt, i, redeemScript, txscript.SigHashAll, bobPrivKey)
		if err != nil {
			log.Fatalf("Failed to generate Bob's signature for input %d: %v", i, err)
		}

		// Create the signature script
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
	}

	// Serialize the signed transaction
	var buf bytes.Buffer
	err = psbt.Serialize(&buf)
	if err != nil {
		log.Fatalf("Failed to serialize signed transaction: %v", err)
	}

	// Convert to hex string
	rawTxHex := hex.EncodeToString(buf.Bytes())
	fmt.Printf("Raw Signed Transaction: %s\n", rawTxHex)

	return rawTxHex
}

// SendRawTransaction sends the raw signed transaction to the Bitcoin Testnet.
// func SendRawTransaction(rawTxHex string) {
// 	// Connect to the local Bitcoin Testnet node
// 	connCfg := &rpcclient.ConnConfig{
// 		Host:         "localhost:18332", // Replace with your node's address and port
// 		User:         "yogendra",     // Replace with your RPC username
// 		Pass:         "yogendra", // Replace with your RPC password
// 		HTTPPostMode: true,              // Bitcoin Core supports only HTTP POST mode
// 		DisableTLS:   true,              // Disable TLS for simplicity
// 	}

// 	client, err := rpcclient.New(connCfg, nil)
// 	if err != nil {
// 		log.Fatalf("Error creating new Bitcoin RPC client: %v", err)
// 	}
// 	defer client.Shutdown()

// 	// Decode the raw transaction from hex
// 	rawTxBytes, err := hex.DecodeString(rawTxHex)
// 	if err != nil {
// 		log.Fatalf("Failed to decode raw transaction hex: %v", err)
// 	}

// 	// Create a wire.MsgTx from the raw bytes
// 	tx := wire.NewMsgTx(wire.TxVersion)
// 	err = tx.Deserialize(bytes.NewReader(rawTxBytes))
// 	if err != nil {
// 		log.Fatalf("Failed to deserialize raw transaction: %v", err)
// 	}

// 	// Send the raw transaction
// 	txHash, err := client.SendRawTransaction(tx, false)
// 	if err != nil {
// 		log.Fatalf("Failed to send raw transaction: %v", err)
// 	}

// 	fmt.Printf("Transaction successfully sent! TxHash: %s\n", txHash.String())
// }

