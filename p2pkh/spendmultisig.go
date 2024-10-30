// spendmultisig.go
package main

import (
	"bytes"
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcec/v2"
)

// hexToPrivateKey parses a hex-encoded private key string and returns a btcec private key.
func hexToPrivateKey1(hexKey string) (*btcec.PrivateKey, error) {
	privKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}
	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes) // Ignore the public key
	return privKey, nil
}

// SpendMultiSig constructs and signs a 2-of-2 multisig transaction
func SpendMultiSig() (string, error) {
	// Use the private keys for Wallet 1 and Wallet 2
	privKeyHex1 := "00c410a3a4d84ba56873142f823e99bec321420c3265e2fc2db8f7671e9dc3ce5c"
	privKeyHex2 := "00bf1eedb8b718a18197578d6bfdc3a21c8f70cb624c4eaeab61b3c8c46e74cfa4"

	privKey1, err := hexToPrivateKey(privKeyHex1)
	if err != nil {
		return "", err
	}
	privKey2, err := hexToPrivateKey(privKeyHex2)
	if err != nil {
		return "", err
	}

	// Public keys for Wallet 1 and Wallet 2
	pk1 := privKey1.PubKey().SerializeCompressed()
	pk2 := privKey2.PubKey().SerializeCompressed()

	// Create the 2-of-2 multisig redeem script
	builder := txscript.NewScriptBuilder()
	builder.AddOp(txscript.OP_2)
	builder.AddData(pk1).AddData(pk2)
	builder.AddOp(txscript.OP_2)
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	redeemScript, err := builder.Script()
	if err != nil {
		return "", err
	}

	// Create a new transaction
	redeemTx := wire.NewMsgTx(wire.TxVersion)

	// Use the UTXO hash (transaction ID of the UTXO you're spending)
	utxoHash, err := chainhash.NewHashFromStr("b14e1fda5f3f74cfceb43bce9b35b0bbe3d5666d21cc802b30dc2fb738a475fe") // Replace with actual UTXO hash
	if err != nil {
		return "", err
	}
	outPoint := wire.NewOutPoint(utxoHash, 1) // Replace 1 with the correct output index of the UTXO
	txIn := wire.NewTxIn(outPoint, nil, nil)
	redeemTx.AddTxIn(txIn)

	// Add the output (destination address and amount)
	decodedAddr, err := btcutil.DecodeAddress("mmRNAidg2GSkgCzhZEKwWtrLyazTGgaUWk", &chaincfg.TestNet3Params) // Replace with actual destination address
	if err != nil {
		return "", err
	}
	destinationAddrByte, err := txscript.PayToAddrScript(decodedAddr)
	if err != nil {
		return "", err
	}
	redeemTxOut := wire.NewTxOut(43000, destinationAddrByte) // 43000 is the amount in satoshis; adjust as needed
	redeemTx.AddTxOut(redeemTxOut)

	// Sign the transaction with both private keys
	sig1, err := txscript.RawTxInSignature(redeemTx, 0, redeemScript, txscript.SigHashAll, privKey1)
	if err != nil {
		return "", err
	}

	sig2, err := txscript.RawTxInSignature(redeemTx, 0, redeemScript, txscript.SigHashAll, privKey2)
	if err != nil {
		return "", err
	}

	// Build the SignatureScript (unlocking script)
	signature := txscript.NewScriptBuilder()
	signature.AddOp(txscript.OP_FALSE).AddData(sig1).AddData(sig2).AddData(redeemScript)
	signatureScript, err := signature.Script()
	if err != nil {
		return "", err
	}
	redeemTx.TxIn[0].SignatureScript = signatureScript

	// Serialize the signed transaction
	var signedTx bytes.Buffer
	redeemTx.Serialize(&signedTx)
	hexSignedTx := hex.EncodeToString(signedTx.Bytes())

	return hexSignedTx, nil
}
