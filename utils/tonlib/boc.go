package tonlib

import (
	"math/big"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type JettonTransferOption struct {
	Jetton              string
	Destination         string
	ResponseDestination string
	Amount              uint64
	Message             string
}

type Transaction struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Payload string `json:"payload"`
}

func CreateTransaction(opt JettonTransferOption) (*tlb.InternalMessage, error) {

	destinationAddress := address.MustParseAddr(opt.Destination)
	responseAddress := address.MustParseAddr(opt.ResponseDestination)

	jettonAddress, err := GetJettonWallet(opt.ResponseDestination, opt.Jetton)
	if err != nil || jettonAddress == nil {
		return nil, err
	}

	forwardPayload := cell.BeginCell().
		MustStoreUInt(0, 32).
		MustStoreStringSnake(opt.Message).
		EndCell()

	boc := cell.BeginCell().
		MustStoreUInt(0xf8a7ea5, 32).
		MustStoreUInt(0, 64).
		MustStoreBigCoins(new(big.Int).SetUint64(opt.Amount)).
		MustStoreAddr(destinationAddress).
		MustStoreAddr(responseAddress).
		MustStoreBoolBit(false).
		MustStoreCoins(1).
		MustStoreBoolBit(true).
		MustStoreRef(forwardPayload).
		EndCell()

	return &tlb.InternalMessage{
		Bounce:  true,
		DstAddr: address.MustParseAddr(*jettonAddress),
		Amount:  tlb.MustFromTON("0.1"),
		Body:    boc,
	}, nil

}
