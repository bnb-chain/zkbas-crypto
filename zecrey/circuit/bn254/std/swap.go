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
//
//import (
//	"errors"
//	"github.com/consensys/gnark-crypto/ecc"
//	"github.com/consensys/gnark/std/algebra/twistededwards"
//	"github.com/consensys/gnark/std/hash/mimc"
//	"log"
//	"zecrey-crypto/hash/bn254/zmimc"
//	"zecrey-crypto/zecrey/twistededwards/tebn254/zecrey"
//)
//
//// SwapProof in circuit
//type SwapProofConstraints struct {
//	// commitments
//	// valid Enc
//	A_C_ufeeL_Delta, A_CufeeR_DeltaHExpb_fee_DeltaInv Point
//	Z_r_Deltafee                                      Variable
//	// Ownership
//	A_pk_u, A_T_uAC_uARPrimeInv, A_T_ufeeC_ufeeRPrimeInv Point
//	Z_sk_u, Z_bar_r_A, Z_bar_r_fee, Z_sk_uInv            Variable
//	// range proofs
//	//ARangeProof   CtRangeProofConstraints
//	//FeeRangeProof CtRangeProofConstraints
//	// common inputs
//	// user asset A balance Enc
//	C_uA ElGamalEncConstraints
//	// user asset fee balance Enc
//	C_ufee ElGamalEncConstraints
//	// user asset fee Delta Enc
//	C_ufee_Delta ElGamalEncConstraints
//	// user asset A,B Delta Enc
//	C_uA_Delta, C_uB_Delta ElGamalEncConstraints
//	// liquidity pool asset A,B Delta Enc
//	LC_DaoA_Delta, LC_DaoB_Delta ElGamalEncConstraints
//	// public keys
//	Pk_Dao, Pk_u Point
//	// random value for Delta A & B
//	R_DeltaA, R_DeltaB Variable
//	// commitment for user asset A & fee
//	T_uA, T_ufee Point
//	// liquidity pool asset B balance
//	LC_DaoB ElGamalEncConstraints
//	// random value for dao liquidity asset B
//	R_DaoB Variable
//	// asset A,B,fee Delta & dao liquidity asset B balance
//	B_A_Delta, B_B_Delta, B_fee_Delta Variable
//	B_DaoA, B_DaoB                    Variable
//	// alpha = \delta{x} / x
//	// beta = \delta{y} / y
//	// gamma = 1 - fee %
//	Alpha, Beta                    Variable
//	Gamma                          Variable
//	IsEnabled                      Variable
//	AssetAId, AssetBId, AssetFeeId Variable
//}
//
//// define tests for verifying the swap proof
//func (circuit SwapProofConstraints) Define(curveID ecc.ID, api API) error {
//	// first check if C = c_1 \oplus c_2
//	// get edwards curve params
//	params, err := twistededwards.NewEdCurve(curveID)
//	if err != nil {
//		return err
//	}
//	// verify H
//	H := Point{
//		X: api.Constant(HX),
//		Y: api.Constant(HY),
//	}
//	// mimc
//	hFunc, err := mimc.NewMiMC(zmimc.SEED, curveID, api)
//	if err != nil {
//		return err
//	}
//	VerifySwapProof(api, circuit, params, hFunc, H)
//
//	return nil
//}
//
///*
//	VerifyWithdrawProof verify the withdraw proof in circuit
//	@api: the constraint system
//	@proof: withdraw proof circuit
//	@params: params for the curve tebn254
//*/
//func VerifySwapProof(
//	api API,
//	proof SwapProofConstraints,
//	params twistededwards.EdCurve,
//	hFunc MiMC,
//	h Point,
//) {
//	//IsPointEqual(api, proof.IsEnabled, proof.ARangeProof.A, proof.T_uA)
//	//IsPointEqual(api, proof.IsEnabled, proof.FeeRangeProof.A, proof.T_ufee)
//	var (
//		C_uAPrime, C_ufeePrime       ElGamalEncConstraints
//		C_uAPrimeNeg, C_ufeePrimeNeg ElGamalEncConstraints
//		c                            Variable
//	)
//	// mimc
//	//ARangeFunc, err := mimc.NewMiMC(zmimc.SEED, params.ID, api)
//	//if err != nil {
//	//	return
//	//}
//	//feeRangeFunc, err := mimc.NewMiMC(zmimc.SEED, params.ID, api)
//	//if err != nil {
//	//	return
//	//}
//	//VerifyCtRangeProof(api, proof.ARangeProof, params, ARangeFunc)
//	//VerifyCtRangeProof(api, proof.FeeRangeProof, params, feeRangeFunc)
//	// challenge buf
//	hFunc.Write(FixedCurveParam(api))
//	writePointIntoBuf(&hFunc, proof.Pk_u)
//	writePointIntoBuf(&hFunc, proof.Pk_Dao)
//	writeEncIntoBuf(&hFunc, proof.C_uA)
//	writeEncIntoBuf(&hFunc, proof.C_ufee)
//	writeEncIntoBuf(&hFunc, proof.C_uA_Delta)
//	writeEncIntoBuf(&hFunc, proof.C_ufee_Delta)
//	writePointIntoBuf(&hFunc, proof.T_uA)
//	writePointIntoBuf(&hFunc, proof.T_ufee)
//	// write into buf
//	writePointIntoBuf(&hFunc, proof.A_C_ufeeL_Delta)
//	writePointIntoBuf(&hFunc, proof.A_CufeeR_DeltaHExpb_fee_DeltaInv)
//	// write into buf
//	writePointIntoBuf(&hFunc, proof.A_pk_u)
//	writePointIntoBuf(&hFunc, proof.A_T_uAC_uARPrimeInv)
//	writePointIntoBuf(&hFunc, proof.A_T_ufeeC_ufeeRPrimeInv)
//	// compute challenge
//	c = hFunc.Sum()
//	// TODO verify params
//	verifySwapParams(api, proof, proof.IsEnabled, params, h)
//	// verify Enc
//	var l1, r1 Point
//	l1.ScalarMulNonFixedBase(api, &proof.Pk_u, proof.Z_r_Deltafee, params)
//	r1.ScalarMulNonFixedBase(api, &proof.C_ufee_Delta.CL, c, params)
//	r1.AddGeneric(api, &r1, &proof.A_C_ufeeL_Delta, params)
//	IsPointEqual(api, proof.IsEnabled, l1, r1)
//	// verify ownership
//	// l2,r2
//	var l2, r2 Point
//	l2.ScalarMulFixedBase(api, params.BaseX, params.BaseY, proof.Z_sk_u, params)
//	r2.ScalarMulNonFixedBase(api, &proof.Pk_u, c, params)
//	r2.AddGeneric(api, &r2, &proof.A_pk_u, params)
//	IsPointEqual(api, proof.IsEnabled, l2, r2)
//	C_uAPrime = EncSub(api, proof.C_uA, proof.C_uA_Delta, params)
//	assetDelta := api.Sub(proof.AssetAId, proof.AssetFeeId)
//	isSameAsset := api.IsZero(assetDelta)
//	C_uAPrime2 := EncSub(api, C_uAPrime, proof.C_ufee_Delta, params)
//	C_uAPrime = SelectElgamal(api, isSameAsset, C_uAPrime2, C_uAPrime)
//	C_ufeePrime = EncSub(api, proof.C_ufee, proof.C_ufee_Delta, params)
//	C_ufeePrime = SelectElgamal(api, isSameAsset, C_uAPrime, C_ufeePrime)
//	C_uAPrimeNeg = NegElgamal(api, C_uAPrime)
//	C_ufeePrimeNeg = NegElgamal(api, C_ufeePrime)
//	// l3,r3
//	var g_z_bar_r_A, l3, r3 Point
//	g_z_bar_r_A.ScalarMulFixedBase(api, params.BaseX, params.BaseY, proof.Z_bar_r_A, params)
//	l3.ScalarMulNonFixedBase(api, &C_uAPrimeNeg.CL, proof.Z_sk_uInv, params)
//	l3.AddGeneric(api, &l3, &g_z_bar_r_A, params)
//	r3.AddGeneric(api, &proof.T_uA, &C_uAPrimeNeg.CR, params)
//	r3.ScalarMulNonFixedBase(api, &r3, c, params)
//	r3.AddGeneric(api, &r3, &proof.A_T_uAC_uARPrimeInv, params)
//	IsPointEqual(api, proof.IsEnabled, l3, r3)
//
//	// l4,r4
//	var g_z_bar_r_fee, l4, r4 Point
//	g_z_bar_r_fee.ScalarMulFixedBase(api, params.BaseX, params.BaseY, proof.Z_bar_r_fee, params)
//	l4.ScalarMulNonFixedBase(api, &C_ufeePrimeNeg.CL, proof.Z_sk_uInv, params)
//	l4.AddGeneric(api, &l4, &g_z_bar_r_fee, params)
//	r4.AddGeneric(api, &proof.T_ufee, &C_ufeePrimeNeg.CR, params)
//	r4.ScalarMulNonFixedBase(api, &r4, c, params)
//	r4.AddGeneric(api, &r4, &proof.A_T_ufeeC_ufeeRPrimeInv, params)
//	IsPointEqual(api, proof.IsEnabled, l4, r4)
//}
//
//func SetEmptySwapProofWitness() (witness SwapProofConstraints) {
//	witness.A_C_ufeeL_Delta, _ = SetPointWitness(BasePoint)
//
//	witness.A_CufeeR_DeltaHExpb_fee_DeltaInv, _ = SetPointWitness(BasePoint)
//
//	// response
//	witness.Z_r_Deltafee.Assign(ZeroInt)
//	witness.A_pk_u, _ = SetPointWitness(BasePoint)
//
//	witness.A_T_uAC_uARPrimeInv, _ = SetPointWitness(BasePoint)
//
//	witness.A_T_ufeeC_ufeeRPrimeInv, _ = SetPointWitness(BasePoint)
//
//	witness.Z_sk_u.Assign(ZeroInt)
//	witness.Z_bar_r_A.Assign(ZeroInt)
//	witness.Z_bar_r_fee.Assign(ZeroInt)
//	witness.Z_sk_uInv.Assign(ZeroInt)
//	//witness.ARangeProof, _ = SetCtRangeProofWitness(ARangeProof, isEnabled)
//	//if err != nil {
//	//	return witness, err
//	//}
//	//witness.FeeRangeProof, _ = SetCtRangeProofWitness(FeeRangeProof, isEnabled)
//	//if err != nil {
//	//	return witness, err
//	//}
//	// common inputs
//	witness.C_uA, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.C_ufee, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.C_ufee_Delta, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.C_uA_Delta, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.C_uB_Delta, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.LC_DaoA_Delta, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.LC_DaoB_Delta, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.Pk_Dao, _ = SetPointWitness(BasePoint)
//
//	witness.Pk_u, _ = SetPointWitness(BasePoint)
//
//	witness.R_DeltaA.Assign(ZeroInt)
//	witness.R_DeltaB.Assign(ZeroInt)
//	witness.T_uA, _ = SetPointWitness(BasePoint)
//
//	witness.T_ufee, _ = SetPointWitness(BasePoint)
//
//	witness.LC_DaoB, _ = SetElGamalEncWitness(ZeroElgamalEnc)
//
//	witness.R_DaoB.Assign(ZeroInt)
//	witness.B_A_Delta.Assign(ZeroInt)
//	witness.B_B_Delta.Assign(ZeroInt)
//	witness.B_fee_Delta.Assign(ZeroInt)
//	witness.B_DaoA.Assign(ZeroInt)
//	witness.B_DaoB.Assign(ZeroInt)
//	witness.Alpha.Assign(ZeroInt)
//	witness.Beta.Assign(ZeroInt)
//	witness.Gamma.Assign(ZeroInt)
//
//	witness.AssetAId.Assign(ZeroInt)
//	witness.AssetBId.Assign(ZeroInt)
//	witness.AssetFeeId.Assign(ZeroInt)
//	witness.IsEnabled = SetBoolWitness(false)
//	return witness
//}
//
//// set the witness for swap proof
//func SetSwapProofWitness(proof *zecrey.SwapProof, isEnabled bool) (witness SwapProofConstraints, err error) {
//	if proof == nil {
//		log.Println("[SetSwapProofWitness] invalid params")
//		return witness, err
//	}
//
//	// proof must be correct
//	verifyRes, err := proof.Verify()
//	if err != nil {
//		log.Println("[SetSwapProofWitness] invalid proof:", err)
//		return witness, err
//	}
//	if !verifyRes {
//		log.Println("[SetSwapProofWitness] invalid proof")
//		return witness, errors.New("[SetSwapProofWitness] invalid proof")
//	}
//
//	witness.A_C_ufeeL_Delta, err = SetPointWitness(proof.A_C_ufeeL_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.A_CufeeR_DeltaHExpb_fee_DeltaInv, err = SetPointWitness(proof.A_CufeeR_DeltaHExpb_fee_DeltaInv)
//	if err != nil {
//		return witness, err
//	}
//	// response
//	witness.Z_r_Deltafee.Assign(proof.Z_r_Deltafee)
//	witness.A_pk_u, err = SetPointWitness(proof.A_pk_u)
//	if err != nil {
//		return witness, err
//	}
//	witness.A_T_uAC_uARPrimeInv, err = SetPointWitness(proof.A_T_uAC_uARPrimeInv)
//	if err != nil {
//		return witness, err
//	}
//	witness.A_T_ufeeC_ufeeRPrimeInv, err = SetPointWitness(proof.A_T_ufeeC_ufeeRPrimeInv)
//	if err != nil {
//		return witness, err
//	}
//	witness.Z_sk_u.Assign(proof.Z_sk_u)
//	witness.Z_bar_r_A.Assign(proof.Z_bar_r_A)
//	witness.Z_bar_r_fee.Assign(proof.Z_bar_r_fee)
//	witness.Z_sk_uInv.Assign(proof.Z_sk_uInv)
//	//witness.ARangeProof, err = SetCtRangeProofWitness(proof.ARangeProof, isEnabled)
//	//if err != nil {
//	//	return witness, err
//	//}
//	//witness.FeeRangeProof, err = SetCtRangeProofWitness(proof.FeeRangeProof, isEnabled)
//	//if err != nil {
//	//	return witness, err
//	//}
//	// common inputs
//	witness.C_uA, err = SetElGamalEncWitness(proof.C_uA)
//	if err != nil {
//		return witness, err
//	}
//	witness.C_ufee, err = SetElGamalEncWitness(proof.C_ufee)
//	if err != nil {
//		return witness, err
//	}
//	witness.C_ufee_Delta, err = SetElGamalEncWitness(proof.C_ufee_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.C_uA_Delta, err = SetElGamalEncWitness(proof.C_uA_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.C_uB_Delta, err = SetElGamalEncWitness(proof.C_uB_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.LC_DaoA_Delta, err = SetElGamalEncWitness(proof.LC_poolA_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.LC_DaoB_Delta, err = SetElGamalEncWitness(proof.LC_poolB_Delta)
//	if err != nil {
//		return witness, err
//	}
//	witness.Pk_Dao, err = SetPointWitness(proof.Pk_pool)
//	if err != nil {
//		return witness, err
//	}
//	witness.Pk_u, err = SetPointWitness(proof.Pk_u)
//	if err != nil {
//		return witness, err
//	}
//	witness.R_DeltaA.Assign(proof.R_DeltaA)
//	witness.R_DeltaB.Assign(proof.R_DeltaB)
//	witness.T_uA, err = SetPointWitness(proof.T_uA)
//	if err != nil {
//		return witness, err
//	}
//	witness.T_ufee, err = SetPointWitness(proof.T_ufee)
//	if err != nil {
//		return witness, err
//	}
//	witness.LC_DaoB, err = SetElGamalEncWitness(proof.LC_poolB)
//	if err != nil {
//		return witness, err
//	}
//	witness.R_DaoB.Assign(proof.R_poolB)
//	witness.B_A_Delta.Assign(proof.B_A_Delta)
//	witness.B_B_Delta.Assign(proof.B_B_Delta)
//	witness.B_fee_Delta.Assign(proof.B_fee_Delta)
//	witness.B_DaoA.Assign(proof.B_DaoA)
//	witness.B_DaoB.Assign(proof.B_DaoB)
//	witness.Alpha.Assign(proof.Alpha)
//	witness.Beta.Assign(proof.Beta)
//	witness.Gamma.Assign(uint64(proof.Gamma))
//	witness.AssetAId.Assign(uint64(proof.AssetAId))
//	witness.AssetBId.Assign(uint64(proof.AssetBId))
//	witness.AssetFeeId.Assign(uint64(proof.AssetFeeId))
//	witness.IsEnabled = SetBoolWitness(isEnabled)
//	return witness, nil
//}
//
//func verifySwapParams(
//	api API,
//	proof SwapProofConstraints,
//	isEnabled Variable,
//	params twistededwards.EdCurve,
//	h Point,
//) {
//	var C_uA_Delta, C_uB_Delta, LC_DaoA_Delta, LC_DaoB_Delta ElGamalEncConstraints
//	C_uA_Delta = Enc(api, h, proof.B_A_Delta, proof.R_DeltaA, proof.Pk_u, params)
//	C_uB_Delta = Enc(api, h, proof.B_B_Delta, proof.R_DeltaB, proof.Pk_u, params)
//	LC_DaoA_Delta.CL.ScalarMulNonFixedBase(api, &proof.Pk_Dao, proof.R_DeltaA, params)
//	LC_DaoA_Delta.CR = C_uA_Delta.CR
//	LC_DaoB_Delta.CL.ScalarMulNonFixedBase(api, &proof.Pk_Dao, proof.R_DeltaB, params)
//	LC_DaoB_Delta.CR = C_uB_Delta.CR
//	IsElGamalEncEqual(api, isEnabled, C_uA_Delta, proof.C_uA_Delta)
//	IsElGamalEncEqual(api, isEnabled, C_uB_Delta, proof.C_uB_Delta)
//	IsElGamalEncEqual(api, isEnabled, LC_DaoA_Delta, proof.LC_DaoA_Delta)
//	IsElGamalEncEqual(api, isEnabled, LC_DaoB_Delta, proof.LC_DaoB_Delta)
//	// TODO verify AMM info & DAO balance info
//	api.AssertIsLessOrEqual(proof.B_B_Delta, proof.B_DaoB)
//	k := api.Mul(proof.B_DaoA, proof.B_DaoB)
//	kPrime := api.Mul(api.Add(proof.B_DaoA, proof.B_A_Delta), api.Sub(proof.B_DaoB, proof.B_B_Delta))
//	api.AssertIsLessOrEqual(kPrime, k)
//}
