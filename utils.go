package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"golang.org/x/crypto/ripemd160"
)

var curve *btcec.KoblitzCurve = btcec.S256()

// hash160 performs a SHA-256 followed by RIPEMD-160 hash on the input data.
// This is commonly used to generate a redeem-hash for P2SH.
func hash160(data []byte) []byte {
	sha := sha256.New()
	sha.Write(data)
	ripe := ripemd160.New()
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}

// dblSha256 applies SHA-256 twice on the input data. It’s commonly used for checksum calculations.
func dblSha256(data []byte) []byte {
	sha1 := sha256.New()
	sha1.Write(data)
	sha2 := sha256.New()
	sha2.Write(sha1.Sum(nil))
	return sha2.Sum(nil)
}

// privToPub converts a private key to a compressed public key.
func privToPub(key []byte) []byte {
	return compress(curve.ScalarBaseMult(key))
}

// onCurve checks if a given point (x, y) is on the elliptic curve.
func onCurve(x, y *big.Int) bool {
	return curve.IsOnCurve(x, y)
}

// compress compresses a public key point (x, y) to a compressed format.
func compress(x, y *big.Int) []byte {
	two := big.NewInt(2)
	rem := two.Mod(y, two).Uint64()
	rem += 2
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(rem))
	rest := x.Bytes()
	pad := 32 - len(rest)
	if pad != 0 {
		zeroes := make([]byte, pad)
		rest = append(zeroes, rest...)
	}
	return append(b[1:], rest...)
}

// expand decompresses a compressed public key to x, y coordinates.
func expand(key []byte) (*big.Int, *big.Int) {
	params := curve.Params()
	exp := big.NewInt(1)
	exp.Add(params.P, exp)
	exp.Div(exp, big.NewInt(4))
	x := big.NewInt(0).SetBytes(key[1:33])
	y := big.NewInt(0).SetBytes(key[:1])
	beta := big.NewInt(0)
	beta.Exp(x, big.NewInt(3), nil)
	beta.Add(beta, big.NewInt(7))
	beta.Exp(beta, exp, params.P)
	if y.Add(beta, y).Mod(y, big.NewInt(2)).Int64() == 0 {
		y = beta
	} else {
		y = beta.Sub(params.P, beta)
	}
	return x, y
}

// addPrivKeys adds two private keys together and returns the result.
func addPrivKeys(k1, k2 []byte) []byte {
	i1 := big.NewInt(0).SetBytes(k1)
	i2 := big.NewInt(0).SetBytes(k2)
	i1.Add(i1, i2)
	i1.Mod(i1, curve.Params().N)
	k := i1.Bytes()
	zero, _ := hex.DecodeString("00")
	return append(zero, k...)
}

// addPubKeys adds two public keys together and returns the compressed result.
func addPubKeys(k1, k2 []byte) []byte {
	x1, y1 := expand(k1)
	x2, y2 := expand(k2)
	return compress(curve.Add(x1, y1, x2, y2))
}

// uint32ToByte converts a uint32 integer to a 4-byte array.
func uint32ToByte(i uint32) []byte {
	a := make([]byte, 4)
	binary.BigEndian.PutUint32(a, i)
	return a
}

// uint16ToByte converts a uint16 integer to a 2-byte array.
func uint16ToByte(i uint16) []byte {
	a := make([]byte, 2)
	binary.BigEndian.PutUint16(a, i)
	return a[1:]
}

// byteToUint16 converts a byte array to a uint16 integer.
func byteToUint16(b []byte) uint16 {
	if len(b) == 1 {
		zero := make([]byte, 1)
		b = append(zero, b...)
	}
	return binary.BigEndian.Uint16(b)
}