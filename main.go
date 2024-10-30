package main

import (
	"fmt"
)

func main() {


	// Step 1: Create the multisig locking transaction
	fmt.Println("=== Step 1: Creating Multisig Locking Transaction ===")
	multisigTx := createMultisigTx()

	// Step 2: Sign the multisig transaction by Alice and Bob
	fmt.Println("\n=== Step 2: Signing the Multisig Transaction ===")
	SignMultisigPSBT(multisigTx)

	// Step 3: Create the unlocking transaction to spend the multisig UTXO
	fmt.Println("\n=== Step 3: Creating Unlocking Transaction ===")
	multisigUTXO := multisigTx.TxHash().String() // Get the transaction ID of the multisig tx
	unlockTx := createUnlockTx(multisigUTXO)

	// Print a summary of the generated transactions
	fmt.Println("\n=== Transaction Summary ===")
	fmt.Printf("Multisig Locking Transaction: %v\n", multisigTx)
	fmt.Printf("Unlocking Transaction: %v\n", unlockTx)


}
