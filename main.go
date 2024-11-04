package main

import (
	"fmt"
)

func main() {
    // step 0 : Generate Alice and Bob's keypairs
	// Generate Alice's keypair
	// privateKey, publicKey, address := keypair()
	// fmt.Println("Private Key:", privateKey)
	// fmt.Println("Public Key:", publicKey)
	// fmt.Println("Testnet Address:", address)

	// // Generate Bob's keypair
	// privateKey2, publicKey2, address2 := keypair()
	// fmt.Println("\nPrivate Key (Bob):", privateKey2)
	// fmt.Println("Public Key (Bob):", publicKey2)
	// fmt.Println("Testnet Address (Bob):", address2)

	// // generate address for fee recipient
    // privateKey3, publicKey3, address3 := keypair()
	// fmt.Println("\nprivate key fee recipient:", privateKey3)
	// fmt.Println("Fee Recipient Public Key:", publicKey3)
	// fmt.Println("Fee Recipient Testnet Address:", address3)



	// Step 1: Create the multisig locking transaction
	fmt.Println("=== Step 1: Creating Multisig Locking Transaction ===")
	multisigTx := createMultisigTx()

	// Step 2: Sign the multisig transaction by Alice and Bob
	fmt.Println("\n=== Step 2: Signing the Multisig Transaction ===")

	// Sign the multisig transaction and get the raw signed transactionste
	rawTxHex := SignMultisigPSBT(multisigTx)
	fmt.Println("Raw Signed Transaction:", rawTxHex)
	// // Send the raw signed transaction to the Bitcoin Testnet
	// SendRawTransaction(rawTxHex)


	// Step 3: Create the unlocking transaction to spend the multisig UTXO
	fmt.Println("\n=== Step 3: Creating Unlocking Transaction ===")
	multisigUTXO := multisigTx.TxHash().String() // Get the transaction ID of the multisig tx
	unlockTx := createUnlockTx(multisigUTXO)

	// Print a summary of the generated transactions
	fmt.Println("\n=== Transaction Summary ===")
	fmt.Printf("Multisig Locking Transaction: %v\n", multisigTx)
	fmt.Printf("Unlocking Transaction: %v\n", unlockTx)


}
