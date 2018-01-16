package internal

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var _key *ecdsa.PrivateKey

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be be 27 or 28 for legacy reasons.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecover
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
func ecRecover(ctx context.Context, data, sig []byte) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	var _sig [65]byte
	copy(_sig[:], sig)
	_sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.Ecrecover(signHash(data), _sig[:])
	if err != nil {
		return common.Address{}, err
	}
	pubKey := crypto.ToECDSAPub(rpk)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr, nil
}

func CreateAndSign(ctx context.Context) error {
	ctct := getContextValue(ctx, contractKey).(*TicTacToe)
	state, err := ctct.CreateBoard(nil)
	if err != nil {
		return fmt.Errorf("CreateAndSign: %v", err)
	}
	sig, err := crypto.Sign(signHash(state), _key)
	if err != nil {
		return fmt.Errorf("CreateAndSign: %v", err)
	}
	sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	sigs := []string{hexutil.Encode(sig)}
	DeleteAllStates(ctx)
	err = InsertState(ctx, &State{
		State: hexutil.Encode(state),
	}, sigs)
	if err != nil {
		return fmt.Errorf("CreateAndSign: %v", err)
	}
	return nil
}

func extractSignature(signatures []string) (v []uint8, r, s [][32]byte) {
	v = make([]uint8, len(signatures))
	r = make([][32]byte, len(signatures))
	s = make([][32]byte, len(signatures))
	for i, e := range signatures {
		sig := common.FromHex(e)
		v[i] = uint8(sig[64])
		copy(r[i][:], sig[0:32])
		copy(s[i][:], sig[32:64])
	}
	return v, r, s
}
