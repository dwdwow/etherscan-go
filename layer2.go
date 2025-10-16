package etherscan

import (
	"context"
)

// ============================================================================
// Layer 2 Module
// ============================================================================

// GetPlasmaDepositsOpts contains optional parameters for GetPlasmaDeposits
type GetPlasmaDepositsOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of deposits per page
	// Default: 100
	// Use this to control how many deposits are returned per page
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID
	// Note: Only applicable to Polygon (chainid=137)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetPlasmaDeposits returns a list of Plasma Deposits received by an address
//
// This endpoint returns a list of Plasma deposits received by an address on Polygon.
// Plasma deposits are part of Polygon's Layer 2 scaling solution, where users can
// deposit funds from Ethereum mainnet to Polygon.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get Plasma deposits for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespPlasmaDeposit: List of Plasma deposits with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get Plasma deposits for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	deposits, err := client.GetPlasmaDeposits(ctx, addr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, deposit := range deposits {
//	    fmt.Printf("Deposit Hash: %s\n", deposit.Hash)
//	    fmt.Printf("Amount: %s\n", deposit.Amount)
//	    fmt.Printf("Block Number: %s\n", deposit.BlockNumber)
//	}
//
// Note:
//   - Only applicable to Polygon (chainid=137)
//   - Returns empty slice if no deposits found
//   - Useful for tracking Layer 2 deposits
func (c *HTTPClient) GetPlasmaDeposits(ctx context.Context, address string, opts *GetPlasmaDepositsOpts) ([]RespPlasmaDeposit, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["address"] = address
	params["blocktype"] = "blocks"

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
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

// GetDepositTxsOpts contains optional parameters for GetDepositTxs
type GetDepositTxsOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of deposits per page
	// Default: 1000
	Offset int64 `default:"1000" json:"offset"`

	// Sort order for the results
	// Options: "asc" or "desc" (default: "desc")
	Sort string `default:"desc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Note: Only applicable to Arbitrum Stack (42161, 42170, 33139, 660279) and
	// Optimism Stack (10, 8453, 130, 252, 480, 5000, 81457)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetDepositTxs returns a list of deposits in ETH or ERC20 tokens from Ethereum to L2
//
// This endpoint returns a list of deposits from Ethereum mainnet to Layer 2 networks.
// This is useful for tracking cross-chain deposits and Layer 2 activity.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get deposit transactions for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDepositTx: List of deposit transactions
//   - error: Error if the request fails
//
// Example:
//
//	// Get deposit transactions for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	deposits, err := client.GetDepositTxs(ctx, addr, nil)
//
// Note:
//   - Only applicable to Arbitrum Stack and Optimism Stack networks
//   - Returns empty slice if no deposits found
func (c *HTTPClient) GetDepositTxs(ctx context.Context, address string, opts *GetDepositTxsOpts) ([]RespDepositTx, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["address"] = address

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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

// GetWithdrawalTxsOpts contains optional parameters for GetWithdrawalTxs
type GetWithdrawalTxsOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of withdrawals per page
	// Default: 1000
	Offset int64 `default:"1000" json:"offset"`

	// Sort order for the results
	// Options: "asc" or "desc" (default: "desc")
	Sort string `default:"desc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Note: Only applicable to Arbitrum Stack (42161, 42170, 33139, 660279) and
	// Optimism Stack (10, 8453, 130, 252, 480, 5000, 81457)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetWithdrawalTxs returns a list of withdrawals in ETH or ERC20 tokens from L2 to Ethereum
//
// This endpoint returns a list of withdrawals from Layer 2 networks back to Ethereum mainnet.
// This is useful for tracking cross-chain withdrawals and Layer 2 activity.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get withdrawal transactions for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespWithdrawalTx: List of withdrawal transactions
//   - error: Error if the request fails
//
// Example:
//
//	// Get withdrawal transactions for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	withdrawals, err := client.GetWithdrawalTxs(ctx, addr, nil)
//
// Note:
//   - Only applicable to Arbitrum Stack and Optimism Stack networks
//   - Returns empty slice if no withdrawals found
func (c *HTTPClient) GetWithdrawalTxs(ctx context.Context, address string, opts *GetWithdrawalTxsOpts) ([]RespWithdrawalTx, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["address"] = address

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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
