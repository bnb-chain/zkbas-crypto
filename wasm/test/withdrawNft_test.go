package test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdrawNftSegmentFormat(t *testing.T) {

	var segmentFormat *txtypes.WithdrawNftSegmentFormat
	segmentFormat = &txtypes.WithdrawNftSegmentFormat{
		AccountIndex:      1,
		NftIndex:          15,
		ToAddress:         "0x507Bd54B4232561BC0Ca106F7b029d064fD6bE4c",
		GasAccountIndex:   1,
		GasFeeAssetId:     3,
		GasFeeAssetAmount: "3",
		ExpiredAt:         1654656781000, // milli seconds
		Nonce:             1,
	}

	res, err := json.Marshal(segmentFormat)
	assert.Nil(t, err)

	log.Println(string(res))
}
