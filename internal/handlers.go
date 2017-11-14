package internal

import (
	"context"
	"fmt"
	"log"

	models "github.com/Magicking/protochannel/models"
	op "github.com/Magicking/protochannel/restapi/operations"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CommitToChannel(ctx context.Context, params op.CommitToChannelParams) middleware.Responder {
	msg := params.Message
	log.Printf("%s", *msg.Data)
	log.Printf("%s", *msg.Signature)
	data := common.FromHex(*msg.Data)
	sig := common.FromHex(*msg.Signature)
	addr, err := ecRecover(ctx, data, sig)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "MakeTemplate", err)
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	ctct := getContextValue(ctx, contractKey).(*WhitelistCaller)
	//TODO set proper From & context
	opts := &bind.CallOpts{From: addr, Context: context.TODO()}
	ret, err := ctct.IsListed(opts, addr)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "IsListed", err)
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	log.Printf("Addr: %s [%v]", addr.String(), ret)
	return op.NewCommitToChannelOK()
}

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
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.Ecrecover(signHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	pubKey := crypto.ToECDSAPub(rpk)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr, nil
}
