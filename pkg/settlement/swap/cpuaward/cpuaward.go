// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpuaward

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/pkg/settlement/swap/transaction"
)

const ()

var ()

// Service is the main interface for interacting with the nodes chequebook.
type Service interface {
	Compute()
}

type service struct {
	lock               sync.Mutex
	transactionService transaction.Service

	ownerAddress common.Address

	erc20Service erc20.Service

	initNum *big.Int
}

// New creates a new chequebook service for the provided chequebook contract.
func NewCPUAward(transactionService transaction.Service, ownerAddress common.Address) (Service, error) {

	return &service{
		transactionService: transactionService,
		ownerAddress:       ownerAddress,
		initNum:            big.NewInt(0),
	}, nil
}

// Compute returns the score of current device's CPU
func (s *service) Compute() {
	// tiker := time.NewTicker(time.Second)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-tiker.C)
	// }
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			println("test")
		}
	}()

	// time.Sleep(time.Minute)
}

// // Deposit starts depositing erc20 token into the chequebook. This returns once the transactions has been broadcast.
// func (s *service) Deposit(ctx context.Context, amount *big.Int) (hash common.Hash, err error) {
// 	balance, err := s.erc20Service.BalanceOf(ctx, s.ownerAddress)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	// check we can afford this so we don't waste gas
// 	if balance.Cmp(amount) < 0 {
// 		return common.Hash{}, ErrInsufficientFunds
// 	}

// 	return s.erc20Service.Transfer(ctx, s.address, amount)
// }

// // WaitForDeposit waits for the deposit transaction to confirm and verifies the result.
// func (s *service) WaitForDeposit(ctx context.Context, txHash common.Hash) error {
// 	receipt, err := s.transactionService.WaitForReceipt(ctx, txHash)
// 	if err != nil {
// 		return err
// 	}
// 	if receipt.Status != 1 {
// 		return transaction.ErrTransactionReverted
// 	}
// 	return nil
// }
