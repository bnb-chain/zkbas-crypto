package schnorr

import (
	"zecrey-crypto/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

func HashSchnorr(A *P256, R *P256) *big.Int {
	ARBytes := util.ContactBytes(A.Bytes(), R.Bytes())
	var buffer bytes.Buffer
	buffer.Write(ARBytes)
	c, _ := util.HashToInt(buffer, zmimc.Hmimc)
	return c
}
