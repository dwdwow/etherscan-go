package etherscan

import (
	"context"
	"strconv"
)

// ============================================================================
// Layer 2 Module
// ============================================================================

// GetPlasmaDepositsOpts contains optional parameters
type GetPlasmaDepositsOpts struct {
	Page            *int
	Offset          *int
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetPlasmaDeposits returns a list of Plasma Deposits received by an address
//
// Note: Only applicable to Polygon (chainid=137)
func (c *HTTPClient) GetPlasmaDeposits(ctx context.Context, address string, opts *GetPlasmaDepositsOpts) ([]RespPlasmaDeposit, error) {
	params := map[string]string{
		"address":   address,
		"blocktype": "blocks",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Page != nil {
			params["page"] = strconv.Itoa(*opts.Page)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.Itoa(*opts.Offset)
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		} else {
			params["chainid"] = "137" // Default to Polygon
		}
		onLimitExceeded = opts.OnLimitExceeded
	} else {
		params["chainid"] = "137"
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "txnbridge",
		params:          params,
		noFoundReturn:   []RespPlasmaDeposit{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespPlasmaDeposit
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDepositTxsOpts contains optional parameters
type GetDepositTxsOpts struct {
	Page            *int
	Offset          *int
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDepositTxs returns a list of deposits in ETH or ERC20 tokens from Ethereum to L2
//
// Note: Only applicable to Arbitrum Stack (42161, 42170, 33139, 660279) and
// Optimism Stack (10, 8453, 130, 252, 480, 5000, 81457)
func (c *HTTPClient) GetDepositTxs(ctx context.Context, address string, opts *GetDepositTxsOpts) ([]RespDepositTx, error) {
	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "1000",
		"sort":    "desc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Page != nil {
			params["page"] = strconv.Itoa(*opts.Page)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.Itoa(*opts.Offset)
		}
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "getdeposittxs",
		params:          params,
		noFoundReturn:   []RespDepositTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDepositTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetWithdrawalTxsOpts contains optional parameters
type GetWithdrawalTxsOpts struct {
	Page            *int
	Offset          *int
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetWithdrawalTxs returns a list of withdrawals in ETH or ERC20 tokens from L2 to Ethereum
//
// Note: Only applicable to Arbitrum Stack (42161, 42170, 33139, 660279) and
// Optimism Stack (10, 8453, 130, 252, 480, 5000, 81457)
func (c *HTTPClient) GetWithdrawalTxs(ctx context.Context, address string, opts *GetWithdrawalTxsOpts) ([]RespWithdrawalTx, error) {
	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "1000",
		"sort":    "desc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Page != nil {
			params["page"] = strconv.Itoa(*opts.Page)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.Itoa(*opts.Offset)
		}
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "getwithdrawaltxs",
		params:          params,
		noFoundReturn:   []RespWithdrawalTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespWithdrawalTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
