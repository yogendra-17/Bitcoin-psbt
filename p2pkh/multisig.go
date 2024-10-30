package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/txscript"
)

// BuildMultiSigRedeemScript generates a 2-of-2 MultiSig redeem script using two public keys
func BuildMultiSigRedeemScript(pubKey1, pubKey2 []byte) ([]byte, error) {
	builder := txscript.NewScriptBuilder()

	// Add the minimum required signatures for spending
	builder.AddOp(txscript.OP_2)

	// Add both public keys
	builder.AddData(pubKey1)
	builder.AddData(pubKey2)

	// Add the total number of public keys and the CHECKMULTISIG opcode
	builder.AddOp(txscript.OP_2) // Total number of keys (2 in this case)
	builder.AddOp(txscript.OP_CHECKMULTISIG)

	// Generate the redeem script
	redeemScript, err := builder.Script()
	if err != nil {
		return nil, err
	}

	return redeemScript, nil
}

func mainMultisig() {
	// Using the provided public keys
	pubKey1, _ := hex.DecodeString("039aa052f944a186bfca84fc7d902041f7526d96bc70ee4940ad748769d5384ff4")
	pubKey2, _ := hex.DecodeString("031d1be613b628330d422f9388854c2dc754861a3d19773aa685e4536dd0d95464")

	// Build the MultiSig redeem script
	redeemScript, err := BuildMultiSigRedeemScript(pubKey1, pubKey2)
	if err != nil {
		log.Fatalf("Error building multisig redeem script: %v", err)
	}

	// Display the human-readable disassembled redeem script
	redeemScriptStr, _ := txscript.DisasmString(redeemScript)
	fmt.Println("MultiSig Redeem Script:", redeemScriptStr)

	// Store the redeem script as byte slice for further use
}
