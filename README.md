# Bitcoin-psbt

This project demonstrates how to create, sign, and unlock a 2-of-2 multisig Bitcoin transaction using Partially Signed Bitcoin Transactions (PSBT) on Bitcoin’s regtest. 

Note: The implementation uses hardcoded placeholders values that need to be replaces before running the repo.

Steps:
1. Creates a 2-of-2 multisig PSBT to lock funds.
2. Signs the PSBT with Alice’s and Bob’s private keys.
3. Unlocks the multisig PSBT, transferring funds to a recipient address (Address Z) and returning the balance to Alice.

## Getting Started
Prerequisites:
- Go 
- btcsuite libraries (btcd, btcutil, and txscript from btcsuite)
- bitcoind regtest server

Install dependencies using:
```
go get github.com/btcsuite/btcd
go get github.com/btcsuite/btcutil

```
Installation
- git clone https://github.com/yogendra-17/Bitcoin-psbt
- cd Bitcoin-psbt
- go run main.go create_multisig_psbt.go sign_multisig_psbt.go unlock_multisig_psbt.go


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
