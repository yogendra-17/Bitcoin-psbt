package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/txscript"
)

func DisassembleScript() (string, error) {
	// locking script 
	lockingScript := "" // replace with the redeem hash.
	script, err := hex.DecodeString(lockingScript)
	if err != nil {
		return "", err
	}
	scriptStr, err := txscript.DisasmString(script)
	if err != nil {
		return "", err
	}
	return scriptStr, nil
}

func mainDisassemble() {
	scriptStr, err := DisassembleScript()
	if err != nil {
		fmt.Println("Error disassembling script:", err)
		return
	}
	fmt.Println("Disassembled Locking Script:", scriptStr)
}
