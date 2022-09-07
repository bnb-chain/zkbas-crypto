/*
 * Copyright © 2021 Zecrey Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package std

import (
	"math/big"
)

type FullExitNftTx struct {
	AccountIndex           int64
	AccountNameHash        []byte
	CreatorAccountIndex    int64
	CreatorAccountNameHash []byte
	CreatorTreasuryRate    int64
	NftIndex               int64
	CollectionId           int64
	NftContentHash         []byte
	NftL1Address           string
	NftL1TokenId           *big.Int
}

type FullExitNftTxConstraints struct {
	AccountIndex           Variable
	AccountNameHash        Variable
	CreatorAccountIndex    Variable
	CreatorAccountNameHash Variable
	CreatorTreasuryRate    Variable
	NftIndex               Variable
	CollectionId           Variable
	NftContentHash         Variable
	NftL1Address           Variable
	NftL1TokenId           Variable
}

func EmptyFullExitNftTxWitness() (witness FullExitNftTxConstraints) {
	return FullExitNftTxConstraints{
		AccountIndex:           ZeroInt,
		AccountNameHash:        ZeroInt,
		CreatorAccountIndex:    ZeroInt,
		CreatorAccountNameHash: ZeroInt,
		CreatorTreasuryRate:    ZeroInt,
		NftIndex:               ZeroInt,
		CollectionId:           ZeroInt,
		NftContentHash:         ZeroInt,
		NftL1Address:           ZeroInt,
		NftL1TokenId:           ZeroInt,
	}
}

func SetFullExitNftTxWitness(tx *FullExitNftTx) (witness FullExitNftTxConstraints) {
	witness = FullExitNftTxConstraints{
		AccountIndex:           tx.AccountIndex,
		AccountNameHash:        tx.AccountNameHash,
		CreatorAccountIndex:    tx.CreatorAccountIndex,
		CreatorAccountNameHash: tx.CreatorAccountNameHash,
		CreatorTreasuryRate:    tx.CreatorTreasuryRate,
		NftIndex:               tx.NftIndex,
		CollectionId:           tx.CollectionId,
		NftContentHash:         tx.NftContentHash,
		NftL1Address:           tx.NftL1Address,
		NftL1TokenId:           tx.NftL1TokenId,
	}
	return witness
}

func VerifyFullExitNftTx(
	api API, flag Variable,
	tx FullExitNftTxConstraints,
	accountsBefore [NbAccountsPerTx]AccountConstraints, nftBefore NftConstraints,
) (pubData [PubDataSizePerTx]Variable) {
	pubData = CollectPubDataFromFullExitNft(api, tx)
	// verify params
	IsVariableEqual(api, flag, tx.AccountNameHash, accountsBefore[0].AccountNameHash)
	IsVariableEqual(api, flag, tx.AccountIndex, accountsBefore[0].AccountIndex)
	IsVariableEqual(api, flag, tx.NftIndex, nftBefore.NftIndex)
	isCheck := api.IsZero(api.IsZero(tx.CreatorAccountNameHash))
	isCheck = api.And(flag, isCheck)
	IsVariableEqual(api, flag, tx.CreatorAccountIndex, nftBefore.CreatorAccountIndex)
	IsVariableEqual(api, flag, tx.CreatorTreasuryRate, nftBefore.CreatorTreasuryRate)
	isOwner := api.And(api.IsZero(api.Sub(tx.AccountIndex, nftBefore.OwnerAccountIndex)), flag)
	IsVariableEqual(api, isOwner, tx.NftContentHash, nftBefore.NftContentHash)
	IsVariableEqual(api, isOwner, tx.NftL1Address, nftBefore.NftL1Address)
	IsVariableEqual(api, isOwner, tx.NftL1TokenId, nftBefore.NftL1TokenId)
	tx.NftContentHash = api.Select(isOwner, tx.NftContentHash, 0)
	tx.NftL1Address = api.Select(isOwner, tx.NftL1Address, 0)
	tx.NftL1TokenId = api.Select(isOwner, tx.NftL1TokenId, 0)
	return pubData
}
