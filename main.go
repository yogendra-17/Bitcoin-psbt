package main

import (
	"fmt"
)

func main() {

	// fmt.Println("====multisig redeem script ===")
	// mainMultisig()

	// fmt.Println("==== MultiSig P2SH Address ====")
	// mainP2SH()

	// // Step 1: Generate and display the disassembled locking script
	// fmt.Println("==== Disassembled Locking Script ====")
	// mainDisassemble()

	// // Step 2: Generate and display the unlocking script for spending the multisig UTXO
	// fmt.Println("\n==== Unlocking Script ====")
	// mainUnlocking()

	// fmt.Println("==== Signed MultiSig Transaction ====")
	// signedTx, err := SpendMultiSig()
	// if err != nil {
	// 	fmt.Printf("Error signing multisig transaction: %v", err)
	// }
	// fmt.Println("Signed Transaction Hex:", signedTx)

	fmt.Println("=== Step 1: Creating Multisig Locking PSBT ===")
	multisigPsbt := CreateMultisigPSBT()
	fmt.Printf("Multisig Locking PSBT: %v\n", multisigPsbt)

	fmt.Println("\n=== Step 2: Signing the PSBT ===")
	SignMultisigPSBT(multisigPsbt)

	fmt.Println("\n=== Step 3: Creating Unlocking PSBT ===")
	unlockPsbt := UnlockMultisigPSBT(multisigPsbt.TxHash().String())
	fmt.Printf("Unlocking PSBT: %v\n", unlockPsbt)

}
