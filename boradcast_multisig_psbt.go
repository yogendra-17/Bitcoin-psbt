package main

//boradcast_multisig_psbt.go

// I'm running the regtest network in my terminal,To broadcast the transaction, I will use the `sendrawtransaction` RPC command to submit the signed transaction to the network.
//bitcoin-cli -testnet sendrawtransaction "<multisig_tx_hex>"
//bitcoin-cli -testnet sendrawtransaction "<unlock_tx_hex>"

// verify transaction
// bitcoin-cli -testnet gettransaction "<multisig_tx_id>"

// To verify that the funds were sent correctly by checking the balances of Alice’s address and Address Z:
// bitcoin-cli -testnet -rpcwallet="" listunspent
