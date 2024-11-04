# Bitcoin-psbt

This project demonstrates how to create, sign, and unlock a 2-of-2 multisig Bitcoin transaction using Partially Signed Bitcoin Transactions (PSBT) on Bitcoin’s testnet. 

Note: The implementation uses hardcoded placeholders values that need to be replaces before running the repo.

Steps:
1. Creates a 2-of-2 multisig PSBT to lock funds.
2. Signs the PSBT with Alice’s and Bob’s private keys.
3. Unlocks the multisig PSBT, transferring funds to a recipient address (Address Z) and returning the balance to Alice.

## Getting Started
Prerequisites:
- Go 
- btcsuite libraries (btcd, btcutil, and txscript from btcsuite)
- bitcoind testnet server

setup bitcoind testnet server using:

```bash
# bitcoin.conf
testnet=1
server=1
rpcuser=yourrpcusername
rpcpassword=yourrpcpassword
rpcallowip=127.0.0.1
rpcport=18332

```
start testnet server using:

```bash
bitcoind -daemon -testnet

```

check the testnet server using:

```bash
bitcoin-cli -testnet getblockchaininfo

```



Install dependencies using:
```bash
go get github.com/btcsuite/btcd

```
Installation

- Clone the repository using:
```bash
- git clone https://github.com/yogendra-17/Bitcoin-psbt
- cd Bitcoin-psbt
- go mod tidy

```

## Project Structure

```
PSBT
├── main.go                   
├── create_multisig_psbt.go   
├── sign_multisig_psbt.go
├── unlock_multisig_psbt.go  
└── README.md    
```
Description of Each File
- main.go: Coordinates the overall process by calling the functions in sequence.
- create_multisig_psbt.go: Generates the initial PSBT to lock funds in a multisig address.
- sign_multisig_psbt.go: Signs the PSBT using Alice's and Bob's private keys.
- unlock_multisig_psbt.go: Generates a PSBT to unlock and spend the funds in the multisig address.

## Running the Project


Step 0: Generate Keypairs for Alice, Bob, and the Fee Recipient
Before proceeding with the multisig transaction, you need to generate keypairs for Alice, Bob, and the fee recipient.

To generate keypairs, follow these instructions:

Open main.go.

Comment Out Steps 1, 2, and 3 by adding // at the beginning of each line related to these steps. The main.go file should look like this:

```go
func main() {
    // Step 1: Create the multisig locking transaction
    // fmt.Println("=== Step 1: Creating Multisig Locking Transaction ===")
    // multisigTx := createMultisigTx()

    // Step 2: Sign the multisig transaction by Alice and Bob
    // fmt.Println("\n=== Step 2: Signing the Multisig Transaction ===")
    // rawTxHex := SignMultisigPSBT(multisigTx)
    // fmt.Println("Raw Signed Transaction:", rawTxHex)

    // Step 3: Create the unlocking transaction to spend the multisig UTXO
    // fmt.Println("\n=== Step 3: Creating Unlocking Transaction ===")
    // multisigUTXO := multisigTx.TxHash().String() // Get the transaction ID of the multisig tx
    // unlockTx := createUnlockTx(multisigUTXO)

    // Step 0: Generate Alice and Bob's keypairs
    privateKey, publicKey, address := keypair()
    fmt.Println("Private Key (Alice):", privateKey)
    fmt.Println("Public Key (Alice):", publicKey)
    fmt.Println("Testnet Address (Alice):", address)

    privateKey2, publicKey2, address2 := keypair()
    fmt.Println("\nPrivate Key (Bob):", privateKey2)
    fmt.Println("Public Key (Bob):", publicKey2)
    fmt.Println("Testnet Address (Bob):", address2)

    privateKey3, publicKey3, address3 := keypair()
    fmt.Println("\nPrivate Key (Fee Recipient):", privateKey3)
    fmt.Println("Fee Recipient Public Key:", publicKey3)
    fmt.Println("Fee Recipient Testnet Address:", address3)
}

```

Run the project to generate keypairs:

```bash
go run main.go keypair.go

```
use the bitcoin testnet faucet to get some testnet coins for the alice address,once you have the testnet coins you can get txhash of the utxo on any block explorer.

Replace the txhash of the utxo in the create_multisig_psbt.go file.

Hint:Note down the generated private keys, public keys, and addresses and replace them in the placeholders in the subsequent steps.


Step 1-3: Proceed with the Multisig Transaction Steps

Comment Out Step 0 by adding // at the beginning of each line related to keypair generation.
Uncomment Steps 1, 2, and 3 by removing the // from those lines.

Note: Remember to replace the placeholder values with the actual values you noted down in Step 0.


The main.go file should now look like this:

```go
func main() {
    // Step 0: Generate Alice and Bob's keypairs
    // privateKey, publicKey, address := keypair()
    // fmt.Println("Private Key (Alice):", privateKey)
    // fmt.Println("Public Key (Alice):", publicKey)
    // fmt.Println("Testnet Address (Alice):", address)

    // privateKey2, publicKey2, address2 := keypair()
    // fmt.Println("\nPrivate Key (Bob):", privateKey2)
    // fmt.Println("Public Key (Bob):", publicKey2)
    // fmt.Println("Testnet Address (Bob):", address2)

    // privateKey3, publicKey3, address3 := keypair()
    // fmt.Println("\nPrivate Key (Fee Recipient):", privateKey3)
    // fmt.Println("Fee Recipient Public Key:", publicKey3)
    // fmt.Println("Fee Recipient Testnet Address:", address3)

    // Step 1: Create the multisig locking transaction
    fmt.Println("=== Step 1: Creating Multisig Locking Transaction ===")
    multisigTx := createMultisigTx()

    // Step 2: Sign the multisig transaction by Alice and Bob
    fmt.Println("\n=== Step 2: Signing the Multisig Transaction ===")
    rawTxHex := SignMultisigPSBT(multisigTx)
    fmt.Println("Raw Signed Transaction:", rawTxHex)

    // Step 3: Create the unlocking transaction to spend the multisig UTXO
    fmt.Println("\n=== Step 3: Creating Unlocking Transaction ===")
    multisigUTXO := multisigTx.TxHash().String() // Get the transaction ID of the multisig tx
    unlockTx := createUnlockTx(multisigUTXO)

    fmt.Println("\n=== Transaction Summary ===")
    fmt.Printf("Multisig Locking Transaction: %v\n", multisigTx)
    fmt.Printf("Unlocking Transaction: %v\n", unlockTx)
}
``` 

Running the Full Process
Now you can proceed to run the entire process:

```bash
go run main.go create_multisig_psbt.go sign_multisig_psbt.go unlock_multisig_psbt.go
```

## Placeholder values to replace

### Summary of Placeholders

| Placeholder               | File                     | Variable                | Suggested Replacement                        |
|---------------------------|--------------------------|-------------------------|----------------------------------------------|
| **Alice’s Private Key**       | `sign_multisig_psbt.go`  | `alicePrivKeyHex`       | Alice’s actual private key (hex-encoded)    |
| **Bob’s Private Key**         | `sign_multisig_psbt.go`  | `bobPrivKeyHex`         | Bob’s actual private key (hex-encoded)      |
| **Alice’s Public Key**        | `create_multisig_psbt.go`| `alicePubKey`           | Alice’s actual public key (compressed hex)  |
| **Bob’s Public Key**          | `create_multisig_psbt.go`| `bobPubKey`             | Bob’s actual public key (compressed hex)    |
| **UTXO Transaction ID**       | `create_multisig_psbt.go`| `utxoTxID`              | Transaction ID of actual UTXO                            |
| **Address Z (Fee Recipient)** | `create_multisig_psbt.go`, `unlock_multisig_psbt.go` | `feeAddressStr` | Replace with actual fee recipient address                      |
| **Alice’s Address**           | `unlock_multisig_psbt.go`| `aliceAddress`          | Replace with Alice’s receiving address                   |
| **UTXO Value**                | `create_multisig_psbt.go`, `unlock_multisig_psbt.go` | `utxoValue` | Actual UTXO value in satoshis              |
