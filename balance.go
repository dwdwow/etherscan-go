package etherscan

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// GetEthBalanceOpts contains optional parameters for GetEthBalance
type GetEthBalanceOpts struct {
	// Tag specifies the block parameter to get balance at
	// Options:
	//   - "latest" (default): Get balance at the most recent block
	//   - "earliest": Get balance at the earliest block
	//   - "pending": Get balance at the pending block
	Tag *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetEthBalance returns the Ether balance of a given address
//
// This endpoint returns the ETH balance of a specified address.
// The balance is returned in Wei (the smallest unit of ETH).
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to check ETH balance for (must be a valid Ethereum address)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - string: The balance in Wei as a string
//   - error: Error if the request fails
//
// Example:
//
//	// Basic usage with defaults
//	balance, err := client.GetEthBalance(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Balance: %s wei\n", balance)
//
//	// With custom options
//	tag := "latest"
//	chainID := etherscan.EthereumMainnet
//	balance, err := client.GetEthBalance(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", &GetEthBalanceOpts{
//	    Tag:     &tag,
//	    ChainID: &chainID,
//	})
//
//	// Query different chain
//	polygonChainID := etherscan.PolygonMainnet
//	balance, err := client.GetEthBalance(ctx, address, &GetEthBalanceOpts{
//	    ChainID: &polygonChainID,
//	})
//
// Note:
//   - The balance is returned in Wei (1 ETH = 10^18 Wei)
//   - For addresses with no balance, returns "0"
//   - This endpoint is not rate limited for free tier users
func (c *HTTPClient) GetEthBalance(ctx context.Context, address string, opts *GetEthBalanceOpts) (string, error) {
	params := map[string]string{"address": address}

	if opts != nil {
		if opts.Tag != nil {
			params["tag"] = *opts.Tag
		} else {
			params["tag"] = "latest"
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
	} else {
		params["tag"] = "latest"
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil && opts.OnLimitExceeded != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "balance",
		params:          params,
		noFoundReturn:   "0",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}

// GetEthBalancesOpts contains optional parameters for GetEthBalances
type GetEthBalancesOpts struct {
	// Tag specifies the block parameter to get balances at
	// Options:
	//   - "latest" (default): Get balances at the most recent block
	//   - "earliest": Get balances at the earliest block
	//   - "pending": Get balances at the pending block
	Tag *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetEthBalances returns Ether balances for multiple addresses in a single call
//
// This endpoint allows you to get ETH balances for up to 20 addresses in a single API call,
// which is more efficient than making individual calls for each address.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - addresses: List of addresses to get balances for (maximum 20 addresses per call)
//     Each address must be a valid Ethereum address format
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespEthBalanceEntry: List of balance entries containing account address and balance
//   - error: Error if the request fails
//
// Example:
//
//	// Get balances for multiple addresses
//	addresses := []string{
//	    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
//	    "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
//	    "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
//	}
//
//	balances, err := client.GetEthBalances(ctx, addresses, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, bal := range balances {
//	    fmt.Printf("Address: %s, Balance: %s wei\n", bal.Account, bal.Balance)
//	}
//
//	// With custom options
//	tag := "latest"
//	chainID := etherscan.EthereumMainnet
//	balances, err := client.GetEthBalances(ctx, addresses, &GetEthBalancesOpts{
//	    Tag:     &tag,
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Maximum 20 addresses per call
//   - All balances are returned in Wei (1 ETH = 10^18 Wei)
//   - For addresses with no balance, the balance will be "0"
//   - This endpoint is not rate limited for free tier users
//   - All addresses must be valid Ethereum address format
func (c *HTTPClient) GetEthBalances(ctx context.Context, addresses []string, opts *GetEthBalancesOpts) ([]RespEthBalanceEntry, error) {
	params := map[string]string{
		"address": strings.Join(addresses, ","),
		"tag":     "latest",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Tag != nil {
			params["tag"] = *opts.Tag
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "balancemulti",
		params:          params,
		noFoundReturn:   []RespEthBalanceEntry{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEthBalanceEntry
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
