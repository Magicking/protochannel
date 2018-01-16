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
)

func getState(ctx context.Context) (string, []string, error) {
	st, err := LastState(ctx)
	if err != nil {
		return "", nil, err
	}
	if st == nil {
		err = CreateAndSign(ctx)
		if err != nil {
			return "", nil, err
		}
		st, err = LastState(ctx)
		if err != nil {
			return "", nil, err
		}
		if st == nil {
			log.Fatalf("Could not create state")
		}
	}
	return st.State, []string(st.Signatures), nil
}

func Status(ctx context.Context, params op.StatusParams) middleware.Responder {
	channel_index := 0
	st, sigs, err := getState(ctx)
	if err != nil {
		err_str := fmt.Sprintf("Could not get state: %v", err)
		log.Printf(err_str)
		return op.NewStatusDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	chn0 := &models.Channel{
		Channel:    int64(channel_index),
		State:      st,
		Signatures: sigs,
	}

	return op.NewStatusOK().WithPayload(&models.Information{
		Abi:             TicTacToeABI,
		Channels:        []*models.Channel{chn0},
		ContractAddress: contractAddr.String(),
	})
}

func SignOffCommit(ctx context.Context, params op.SignOffCommitParams) middleware.Responder {
	msg := params.Message
	log.Printf("%s", *msg.Data)
	log.Printf("%s", msg.Signatures)
	if len(msg.Signatures) == 0 {
		err_str := fmt.Sprintf("No signature provided")
		log.Printf(err_str)
		return op.NewSignOffCommitDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	state := common.FromHex(*msg.Data)
	v, r, s := extractSignature(msg.Signatures)
	ctct := getContextValue(ctx, contractKey).(*TicTacToe)
	//TODO Timeout ?
	opts := &bind.CallOpts{Context: context.TODO()}
	ret, err := ctct.Verify(opts, state, v[0], r[0], s[0])
	log.Printf("CanChallenge: %v", ret)
	if !ret {
		err_str := fmt.Sprintf("State doesn't validate")
		log.Printf(err_str)
		return op.NewSignOffCommitDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	err = InsertState(ctx, &State{State: common.ToHex(state)}, msg.Signatures)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "InsertState", err)
		log.Printf(err_str)
		return op.NewSignOffCommitDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	return op.NewSignOffCommitOK()
}

func CommitToChannel(ctx context.Context, params op.CommitToChannelParams) middleware.Responder {
	msg := params.Message
	state := common.FromHex(*msg.Data)
	if len(msg.Signatures) == 0 {
		err_str := fmt.Sprintf("Signature missing")
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	sig := common.FromHex(msg.Signatures[0])
	addr, err := ecRecover(ctx, state, sig)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "ecRecover", err)
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	//TODO Check operation counter here
	//TODO Check equivalence of old_state against new_state
	ctct := getContextValue(ctx, contractKey).(*TicTacToe)
	var r, s [32]byte
	v := uint8(sig[64])
	copy(r[:], sig[0:32])
	copy(s[:], sig[32:64])
	//TODO set context
	opts := &bind.CallOpts{From: addr, Context: context.TODO()}
	log.Printf("addr=%s, state=%+v, v=%+v, r=%+v, s=%+v\n", addr.String(), state, v, r, s)
	ret, err := ctct.CheckAndApply(opts, state, v, r, s)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "CheckAndApply", err)
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	err = InsertState(ctx, &State{State: common.ToHex(state)}, msg.Signatures)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "InsertState", err)
		log.Printf(err_str)
		return op.NewCommitToChannelDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	state_hex := common.ToHex(ret)
	return op.NewCommitToChannelOK().WithPayload(&models.Message{Data: &state_hex})
}
