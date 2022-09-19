/*
 * Copyright © 2022 ZkBNB Protocol
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

package types

type CreatePairTx struct {
	PairIndex            int64
	AssetAId             int64
	AssetBId             int64
	FeeRate              int64
	TreasuryAccountIndex int64
	TreasuryRate         int64
}

type CreatePairTxConstraints struct {
	PairIndex            Variable
	AssetAId             Variable
	AssetBId             Variable
	FeeRate              Variable
	TreasuryAccountIndex Variable
	TreasuryRate         Variable
}

func EmptyCreatePairTxWitness() (witness CreatePairTxConstraints) {
	return CreatePairTxConstraints{
		PairIndex:            ZeroInt,
		AssetAId:             ZeroInt,
		AssetBId:             ZeroInt,
		FeeRate:              ZeroInt,
		TreasuryAccountIndex: ZeroInt,
		TreasuryRate:         ZeroInt,
	}
}

func SetCreatePairTxWitness(tx *CreatePairTx) (witness CreatePairTxConstraints) {
	witness = CreatePairTxConstraints{
		PairIndex:            tx.PairIndex,
		AssetAId:             tx.AssetAId,
		AssetBId:             tx.AssetBId,
		FeeRate:              tx.FeeRate,
		TreasuryAccountIndex: tx.TreasuryAccountIndex,
		TreasuryRate:         tx.TreasuryRate,
	}
	return witness
}

func VerifyCreatePairTx(
	api API, flag Variable,
	tx CreatePairTxConstraints,
	liquidityBefore LiquidityConstraints,
) (pubData [PubDataSizePerTx]Variable) {
	pubData = CollectPubDataFromCreatePair(api, tx)
	// verify params
	IsVariableEqual(api, flag, tx.PairIndex, liquidityBefore.PairIndex)
	CheckEmptyLiquidityNode(api, flag, liquidityBefore)
	return pubData
}
