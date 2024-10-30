// unlock_multisig_psbt.go
package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// UnlockMultisigPSBT creates a PSBT for spending from the multisig UTXO
func UnlockMultisigPSBT(multisigUTXOHash string) *wire.MsgTx {
	// Address Z for fee (example testnet address)
	addressZ, _ := btcutil.DecodeAddress("bcrt1qzefpec3w9f5jr743pxt3x8q0p5m9p9r85gryes", &chaincfg.RegressionNetParams)
	aliceAddress, _ := btcutil.DecodeAddress("bcrt1qlx24qm04uskmhl80g0du8kd93dfcnp70fw23aa", &chaincfg.RegressionNetParams)

	// PSBT with multisig UTXO as input
	psbt := wire.NewMsgTx(wire.TxVersion)
	utxoHash, _ := chainhash.NewHashFromStr(multisigUTXOHash)
	psbt.AddTxIn(wire.NewTxIn(wire.NewOutPoint(utxoHash, 0), nil, nil)) // Replace with actual output index

	// 1. Calculate output values for Address Z and Alice
	utxoValue := int64(99000) // Assume the multisig UTXO value
	feeOutput := int64(utxoValue * 1 / 100) // 1% to Address Z
	returnOutput := utxoValue - feeOutput - 1000 // Subtract 1% fee and estimated transaction fee

	// 2. Add output 1: 1% to Address Z
	scriptAddrZ, _ := txscript.PayToAddrScript(addressZ)
	psbt.AddTxOut(wire.NewTxOut(feeOutput, scriptAddrZ))

	// 3. Add output 2: Remaining to Alice
	scriptAlice, _ := txscript.PayToAddrScript(aliceAddress)
	psbt.AddTxOut(wire.NewTxOut(returnOutput, scriptAlice))

	fmt.Println("Generated PSBT for unlocking multisig funds.")
	return psbt
}
