package txtypes

import (
	"errors"
	"hash"
	"math/big"
)

type FullExitNftTxInfo struct {
	TxType uint8

	// Get from layer1 events.
	NftIndex        int64
	AccountNameHash []byte

	// Set by layer2.
	AccountIndex           int64
	CreatorAccountIndex    int64
	CreatorTreasuryRate    int64
	CreatorAccountNameHash []byte
	NftL1Address           string
	NftL1TokenId           *big.Int
	NftContentHash         []byte
	CollectionId           int64
}

func (txInfo *FullExitNftTxInfo) GetTxType() int {
	return TxTypeFullExitNft
}

func (txInfo *FullExitNftTxInfo) Validate() error {
	return nil
}

func (txInfo *FullExitNftTxInfo) VerifySignature(pubKey string) error {
	return nil
}

func (txInfo *FullExitNftTxInfo) GetFromAccountIndex() int64 {
	return NilAccountIndex
}

func (txInfo *FullExitNftTxInfo) GetNonce() int64 {
	return NilNonce
}

func (txInfo *FullExitNftTxInfo) GetExpiredAt() int64 {
	return NilExpiredAt
}

func (txInfo *FullExitNftTxInfo) Hash(hFunc hash.Hash) (msgHash []byte, err error) {
	return msgHash, errors.New("not support")
}

func (txInfo *FullExitNftTxInfo) GetGas() (int64, int64, *big.Int) {
	return NilAccountIndex, NilAssetId, nil
}
