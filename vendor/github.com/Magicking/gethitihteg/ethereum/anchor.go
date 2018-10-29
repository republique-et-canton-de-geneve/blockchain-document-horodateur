// Copyright 2017 Sylvain 6120 Laurent
// This file is part of the gethitihteg library.
//
// The gethitihteg library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gethitihteg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gethitihteg library. If not, see <http://www.gnu.org/licenses/>.

package ethereum

import (
	"fmt"
	"math/big"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Anchor is the base wrapper object that reflects a transactor on the
// Ethereum network.
type Anchor struct {
	address *common.Address
	backend bind.ContractTransactor
}

// NewAnchor creates an anchor through which data may be sent through.
func NewAnchor(address *common.Address, backend bind.ContractTransactor) *Anchor {
	return &Anchor{
		address: address,
		backend: backend,
	}
}

// transact executes an actual transaction invocation, first deriving any missing
// authorization fields without signing and sending the transaction.
func (c *Anchor) transactWithoutSigning(opts *bind.TransactOpts, to *common.Address, input []byte) (*types.Transaction, error) {
	var err error

	if to == nil {
		return nil, fmt.Errorf("'to' must not be nil")
	}
	// Ensure a valid value field and resolve the account nonce
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.backend.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}
	// Figure out the gas allowance and gas price values
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.backend.SuggestGasPrice(ensureContext(opts.Context))
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		// Gas estimation cannot succeed without code for method invocations
		//if to != nil {
		//	if code, err := c.backend.PendingCodeAt(ensureContext(opts.Context), *c.address); err != nil {
		//		return nil, err
		//	} else if len(code) == 0 {
		//		return nil, bind.ErrNoCode
		//	}
		//}
		// If the to surely has code (or code is not needed), estimate the transaction
		msg := ethereum.CallMsg{From: opts.From, To: to, Value: value, Data: input}
		gasLimit, err = c.backend.EstimateGas(ensureContext(opts.Context), msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}
	// Create the transaction
	rawTx := types.NewTransaction(nonce, *c.address, value, gasLimit, gasPrice, input)

	return rawTx, nil
}

func (c *Anchor) PrepareData(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return c.transactWithoutSigning(opts, c.address, input)
}

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
