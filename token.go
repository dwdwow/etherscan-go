package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Token Module
// ============================================================================

// GetERC20TotalSupplyOpts contains optional parameters for GetERC20TotalSupply
type GetERC20TotalSupplyOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20TotalSupply returns the current amount of an ERC-20 token in circulation
//
// This endpoint returns the total supply of an ERC-20 token in hex format.
// The total supply represents the total number of tokens that have been minted.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The total supply in hex format (e.g., "0x1b1ae4d6e2ef500000")
//   - error: Error if the request fails
//
// Example:
//
//	// Get total supply of USDC
//	usdcAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	supply, err := client.GetERC20TotalSupply(ctx, usdcAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total supply: %s\n", supply)
//
// Note:
//   - Returns supply in hex format with "0x" prefix
//   - Value is in the token's smallest unit
//   - Use token decimals to convert to human-readable format
func (c *HTTPClient) GetERC20TotalSupply(ctx context.Context, contractAddress string, opts *GetERC20TotalSupplyOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "tokensupply",
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

// GetERC20AccountBalanceOpts contains optional parameters for GetERC20AccountBalance
type GetERC20AccountBalanceOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20AccountBalance returns the current balance of an ERC-20 token of an address
//
// This endpoint returns the current balance of an ERC-20 token for a specific address.
// The balance is returned in the token's smallest unit (wei equivalent).
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - address: The address to check for token balance
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The token balance in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get ERC-20 token balance
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	userAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	balance, err := client.GetERC20AccountBalance(ctx, contractAddr, userAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Token balance: %s\n", balance)
//
// Note:
//   - Returns balance in hex format
//   - Value is in the token's smallest unit
//   - Use token decimals to convert to human-readable format
func (c *HTTPClient) GetERC20AccountBalance(ctx context.Context, contractAddress, address string, opts *GetERC20AccountBalanceOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress
	params["address"] = address

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "tokenbalance",
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

// GetERC20HistoricalTotalSupplyOpts contains optional parameters for GetERC20HistoricalTotalSupply
type GetERC20HistoricalTotalSupplyOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20HistoricalTotalSupply returns the amount of an ERC-20 token in circulation at a certain block height
//
// This endpoint returns the amount of an ERC-20 token in circulation at a specific block height.
// This is useful for analyzing token supply changes over time and historical data.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - blockNo: The block number to check total supply for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The historical total supply in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get historical total supply at a specific block
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	blockNo := 1000000
//	supply, err := client.GetERC20HistoricalTotalSupply(ctx, contractAddr, blockNo, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Historical supply: %s\n", supply)
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns supply in hex format
//   - Value is in the token's smallest unit
func (c *HTTPClient) GetERC20HistoricalTotalSupply(ctx context.Context, contractAddress string, blockNo int64, opts *GetERC20HistoricalTotalSupplyOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress
	params["blockno"] = strconv.FormatInt(blockNo, 10)

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "tokensupplyhistory",
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

// GetERC20HistoricalAccountBalanceOpts contains optional parameters for GetERC20HistoricalAccountBalance
type GetERC20HistoricalAccountBalanceOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20HistoricalAccountBalance returns the balance of an ERC-20 token of an address at a certain block height
//
// This endpoint returns the balance of an ERC-20 token for a specific address at a specific block height.
// This is useful for analyzing historical token balances and tracking changes over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - address: The address to check for token balance
//   - blockNo: The block number to check balance for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The historical token balance in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get historical token balance at a specific block
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	userAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	blockNo := 1000000
//	balance, err := client.GetERC20HistoricalAccountBalance(ctx, contractAddr, userAddr, blockNo, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Historical balance: %s\n", balance)
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns balance in hex format
//   - Value is in the token's smallest unit
func (c *HTTPClient) GetERC20HistoricalAccountBalance(ctx context.Context, contractAddress, address string, blockNo int64, opts *GetERC20HistoricalAccountBalanceOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress
	params["address"] = address
	params["blockno"] = strconv.FormatInt(blockNo, 10)

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "tokenbalancehistory",
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

// GetERC20HoldersOpts contains optional parameters for GetERC20Holders
type GetERC20HoldersOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of holders per page
	// Default: 100, max: 10000
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20Holders returns the current ERC20 token holders and number of tokens held
//
// This endpoint returns the current ERC20 token holders and number of tokens held.
// This is useful for analyzing token distribution and holder information.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespERC20HolderInfo: List of token holders with their addresses and token quantities
//   - error: Error if the request fails
//
// Example:
//
//	// Get ERC-20 token holders
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	holders, err := client.GetERC20Holders(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, holder := range holders {
//	    fmt.Printf("Address: %s, Balance: %s\n", holder.TokenHolderAddress, holder.TokenHolderQuantity)
//	}
//
// Note:
//   - Returns empty slice if no holders found
//   - Maximum 10000 records per page
func (c *HTTPClient) GetERC20Holders(ctx context.Context, contractAddress string, opts *GetERC20HoldersOpts) ([]RespERC20HolderInfo, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "token",
		action:          "tokenholderlist",
		params:          params,
		noFoundReturn:   []RespERC20HolderInfo{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespERC20HolderInfo
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetERC20HolderCountOpts contains optional parameters for GetERC20HolderCount
type GetERC20HolderCountOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC20HolderCount returns the total number of holders for an ERC-20 token
//
// This endpoint returns the total number of holders for an ERC-20 token. This is useful
// for analyzing token distribution and adoption metrics.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The total number of token holders in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get total holder count
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	count, err := client.GetERC20HolderCount(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total holders: %s\n", count)
//
// Note:
//   - Returns count in hex format
//   - Useful for token distribution analysis
func (c *HTTPClient) GetERC20HolderCount(ctx context.Context, contractAddress string, opts *GetERC20HolderCountOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "token",
		action:          "tokenholdercount",
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

// GetTopERC20HoldersOpts contains optional parameters for GetTopERC20Holders
type GetTopERC20HoldersOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetTopERC20Holders returns the top token holders of an ERC-20 token
//
// This endpoint returns the top token holders of an ERC-20 token. This is useful
// for analyzing token concentration and identifying major holders.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20 token
//   - offset: Number of top holders to return (max 1000)
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespTopTokenHolder: List of top token holders with their addresses, quantities and address types
//   - error: Error if the request fails
//
// Example:
//
//	// Get top 100 token holders
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	offset := 100
//	holders, err := client.GetTopERC20Holders(ctx, contractAddr, offset, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, holder := range holders {
//	    fmt.Printf("Address: %s, Balance: %s, Type: %s\n", holder.TokenHolderAddress, holder.TokenHolderQuantity, holder.TokenHolderAddressType)
//	}
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - This beta endpoint is only available on Ethereum mainnet
//   - Maximum offset is 1000
func (c *HTTPClient) GetTopERC20Holders(ctx context.Context, contractAddress string, offset int64, opts *GetTopERC20HoldersOpts) ([]RespTopTokenHolder, error) {
	if offset > 1000 {
		return nil, fmt.Errorf("offset cannot exceed 1000")
	}

	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress
	params["offset"] = strconv.FormatInt(offset, 10)

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "token",
		action:          "topholders",
		params:          params,
		noFoundReturn:   []RespTopTokenHolder{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespTopTokenHolder
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTokenInfoOpts contains optional parameters for GetTokenInfo
type GetTokenInfoOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetTokenInfo returns project information and social media links of an ERC20/ERC721/ERC1155 token
//
// This endpoint returns project information and social media links of an ERC20/ERC721/ERC1155 token.
// This is useful for getting comprehensive token metadata and project information.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddress: The contract address of the ERC-20/ERC-721/ERC-1155 token
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespTokenInfo: Token information including name, symbol, supply, social media links etc.
//   - error: Error if the request fails
//
// Example:
//
//	// Get token information
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	info, err := client.GetTokenInfo(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Token Name: %s\n", info.TokenName)
//	fmt.Printf("Token Symbol: %s\n", info.TokenSymbol)
//	fmt.Printf("Total Supply: %s\n", info.TotalSupply)
//	fmt.Printf("Website: %s\n", info.Website)
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns comprehensive token metadata
func (c *HTTPClient) GetTokenInfo(ctx context.Context, contractAddress string, opts *GetTokenInfoOpts) (*RespTokenInfo, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["contractaddress"] = contractAddress

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "token",
		action:          "tokeninfo",
		params:          params,
		noFoundReturn:   RespTokenInfo{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	// The API returns a list, but we only need the first element
	var resultList []RespTokenInfo
	if err := unmarshalResponse(data, &resultList); err != nil {
		// Try unmarshaling as single object
		var result RespTokenInfo
		if err := unmarshalResponse(data, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}

	if len(resultList) > 0 {
		return &resultList[0], nil
	}

	return &RespTokenInfo{}, nil
}

// GetAccountERC20HoldingsOpts contains optional parameters for GetAccountERC20Holdings
type GetAccountERC20HoldingsOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of tokens per page
	// Default: 100
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetAccountERC20Holdings returns the ERC-20 tokens and amount held by an address
//
// This endpoint returns the ERC-20 tokens and amount held by an address. This is useful
// for analyzing token portfolios and holdings of specific addresses.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to check for token balances
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespERC20Holding: List of ERC-20 token balances held by the address
//   - error: Error if the request fails
//
// Example:
//
//	// Get ERC-20 token holdings for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	holdings, err := client.GetAccountERC20Holdings(ctx, addr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, holding := range holdings {
//	    fmt.Printf("Token: %s, Balance: %s\n", holding.TokenName, holding.TokenBalance)
//	}
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns empty slice if no tokens found
func (c *HTTPClient) GetAccountERC20Holdings(ctx context.Context, address string, opts *GetAccountERC20HoldingsOpts) ([]RespERC20Holding, error) {
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
		action:          "addresstokenbalance",
		params:          params,
		noFoundReturn:   []RespERC20Holding{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespERC20Holding
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountNFTHoldingsOpts contains optional parameters for GetAccountNFTHoldings
type GetAccountNFTHoldingsOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of tokens per page
	// Default: 100
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetAccountNFTHoldings returns the ERC-721 tokens and amount held by an address
//
// This endpoint returns the ERC-721 tokens and amount held by an address. This is useful
// for analyzing NFT portfolios and holdings of specific addresses.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to check for token balances
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespNFTHolding: List of ERC-721 token balances held by the address
//   - error: Error if the request fails
//
// Example:
//
//	// Get ERC-721 token holdings for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	holdings, err := client.GetAccountNFTHoldings(ctx, addr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, holding := range holdings {
//	    fmt.Printf("Token: %s, Balance: %s\n", holding.TokenName, holding.TokenBalance)
//	}
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns empty slice if no tokens found
func (c *HTTPClient) GetAccountNFTHoldings(ctx context.Context, address string, opts *GetAccountNFTHoldingsOpts) ([]RespNFTHolding, error) {
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
		action:          "addresstokennftbalance",
		params:          params,
		noFoundReturn:   []RespNFTHolding{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespNFTHolding
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountNFTInventoriesOpts contains optional parameters for GetAccountNFTInventories
type GetAccountNFTInventoriesOpts struct {
	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of tokens per page
	// Default: 100, max: 1000
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetAccountNFTInventories returns the ERC-721 token inventory of an address, filtered by contract address
//
// This endpoint returns the ERC-721 token inventory of an address, filtered by contract address.
// This is useful for analyzing specific NFT collections held by an address.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to check for token inventory
//   - contractAddress: The ERC-721 token contract address to check for inventory
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespNFTTokenInventory: List of ERC-721 token inventory for the address filtered by contract
//   - error: Error if the request fails
//
// Example:
//
//	// Get ERC-721 token inventory for an address and contract
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	inventory, err := client.GetAccountNFTInventories(ctx, addr, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range inventory {
//	    fmt.Printf("Token ID: %s, Balance: %s\n", item.TokenID, item.TokenBalance)
//	}
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Returns empty slice if no tokens found
//   - Maximum 1000 records per page
func (c *HTTPClient) GetAccountNFTInventories(ctx context.Context, address, contractAddress string, opts *GetAccountNFTInventoriesOpts) ([]RespNFTTokenInventory, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["address"] = address
	params["contractaddress"] = contractAddress

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "addresstokennftinventory",
		params:          params,
		noFoundReturn:   []RespNFTTokenInventory{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespNFTTokenInventory
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
