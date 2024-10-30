// unlocking.go
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"github.com/btcsuite/btcd/txscript"
)

// GenerateUnlockingScript builds the unlocking script for spending the multisig UTXO
func GenerateUnlockingScript(sigA, sigB, redeemScript []byte) ([]byte, error) {
	builder := txscript.NewScriptBuilder()

	// Add OP_0 for CHECKMULTISIG bug workaround
	builder.AddOp(txscript.OP_0)

	// Add both signatures
	builder.AddData(sigA)
	builder.AddData(sigB)

	// Add the redeem script (the original multisig script)
	builder.AddData(redeemScript)

	// Generate the final unlocking script
	unlockingScript, err := builder.Script()
	if err != nil {
		return nil, err
	}

	return unlockingScript, nil
}

func mainUnlocking() {
	// Placeholder signatures (replace with real signatures in practice)
	sigA, _ := hex.DecodeString("3045022100c5e8f1...") // Replace with actual signature A
	sigB, _ := hex.DecodeString("3045022100b1f9a4...") // Replace with actual signature B

	// Example redeem script (generated from the multisig)
	redeemScript, err := BuildMultiSigRedeemScript(
		[]byte{0x03, 0x9a, 0xa0, 0x52, 0xf9, 0x44, 0xa1, 0x86, 0xbf, 0xca, 0x84, 0xfc, 0x7d, 0x90, 0x20, 0x41, 0xf7, 0x52, 0x6d, 0x96, 0xbc, 0x70, 0xee, 0x49, 0x40, 0xad, 0x74, 0x87, 0x69, 0xd5, 0x38, 0x4f, 0xf4},
		[]byte{0x03, 0x1d, 0x1b, 0xe6, 0x13, 0xb6, 0x28, 0x33, 0x0d, 0x42, 0x2f, 0x93, 0x88, 0x85, 0x4c, 0x2d, 0xc7, 0x54, 0x86, 0x1a, 0x3d, 0x19, 0x77, 0x3a, 0xa6, 0x85, 0xe4, 0x53, 0x6d, 0xd0, 0xd9, 0x54, 0x64},
	)
	if err != nil {
		log.Fatalf("Error building redeem script: %v", err)
	}

	// Generate the unlocking script
	unlockingScript, err := GenerateUnlockingScript(sigA, sigB, redeemScript)
	if err != nil {
		log.Fatalf("Error building unlocking script: %v", err)
	}

	// Disassemble the unlocking script for readability
	scriptStr, _ := txscript.DisasmString(unlockingScript)
	fmt.Println("Unlocking Script:", scriptStr)
}
