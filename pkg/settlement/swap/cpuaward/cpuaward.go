// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpuaward

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/pkg/settlement/swap/transaction"
	"github.com/klauspost/cpuid"
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
	ticker := time.NewTicker(time.Second * 60)
	go func() {
		for _ = range ticker.C {
			tip1 := fmt.Sprintf("根据以下CPU信息计算奖励发送给节点地址:%x", s.ownerAddress)
			println(tip1)
			score, _ := CPUScore()
			tip2 := fmt.Sprintf("获取CPU分数为:%x", score)
			println(tip2)

		}
	}()

	// time.Sleep(time.Minute)
}

// CPUScore returns the score of current device's CPU
func CPUScore() (score int, err error) {
	// Print basic CPU information:
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", cpuid.CPU.Features)
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")

	// Test if we have a specific feature:
	if cpuid.CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
	}

	score = cpuid.CPU.PhysicalCores * cpuid.CPU.ThreadsPerCore * (cpuid.CPU.CacheLine*100000 + cpuid.CPU.Cache.L1D*100 + cpuid.CPU.Cache.L2*10 + cpuid.CPU.Cache.L3)
	return score, nil
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
