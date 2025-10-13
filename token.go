package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Token Module
// ============================================================================

// GetERC20TotalSupplyOpts contains optional parameters
type GetERC20TotalSupplyOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20TotalSupply returns the current amount of an ERC-20 token in circulation
func (c *HTTPClient) GetERC20TotalSupply(ctx context.Context, contractaddress string, opts *GetERC20TotalSupplyOpts) (string, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetERC20AccountBalanceOpts contains optional parameters
type GetERC20AccountBalanceOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20AccountBalance returns the current balance of an ERC-20 token of an address
func (c *HTTPClient) GetERC20AccountBalance(ctx context.Context, contractaddress, address string, opts *GetERC20AccountBalanceOpts) (string, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
		"address":         address,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetERC20HistoricalTotalSupplyOpts contains optional parameters
type GetERC20HistoricalTotalSupplyOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20HistoricalTotalSupply returns the amount of an ERC-20 token in circulation at a certain block height
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetERC20HistoricalTotalSupply(ctx context.Context, contractaddress string, blockno int, opts *GetERC20HistoricalTotalSupplyOpts) (string, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
		"blockno":         strconv.Itoa(blockno),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetERC20HistoricalAccountBalanceOpts contains optional parameters
type GetERC20HistoricalAccountBalanceOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20HistoricalAccountBalance returns the balance of an ERC-20 token of an address at a certain block height
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetERC20HistoricalAccountBalance(ctx context.Context, contractaddress, address string, blockno int, opts *GetERC20HistoricalAccountBalanceOpts) (string, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
		"address":         address,
		"blockno":         strconv.Itoa(blockno),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetERC20HoldersOpts contains optional parameters
type GetERC20HoldersOpts struct {
	Page            *int
	Offset          *int
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20Holders returns the current ERC20 token holders and number of tokens held
func (c *HTTPClient) GetERC20Holders(ctx context.Context, contractaddress string, opts *GetERC20HoldersOpts) ([]RespERC20HolderInfo, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
		"page":            "1",
		"offset":          "100",
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
		}
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

// GetERC20HolderCountOpts contains optional parameters
type GetERC20HolderCountOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetERC20HolderCount returns the total number of holders for an ERC-20 token
func (c *HTTPClient) GetERC20HolderCount(ctx context.Context, contractaddress string, opts *GetERC20HolderCountOpts) (string, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetTopERC20HoldersOpts contains optional parameters
type GetTopERC20HoldersOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetTopERC20Holders returns the top token holders of an ERC-20 token
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier.
// This beta endpoint is only available on Ethereum mainnet.
func (c *HTTPClient) GetTopERC20Holders(ctx context.Context, contractaddress string, offset int, opts *GetTopERC20HoldersOpts) ([]RespTopTokenHolder, error) {
	if offset > 1000 {
		return nil, fmt.Errorf("offset cannot exceed 1000")
	}

	params := map[string]string{
		"contractaddress": contractaddress,
		"offset":          strconv.Itoa(offset),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetTokenInfoOpts contains optional parameters
type GetTokenInfoOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetTokenInfo returns project information and social media links of an ERC20/ERC721/ERC1155 token
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetTokenInfo(ctx context.Context, contractaddress string, opts *GetTokenInfoOpts) (*RespTokenInfo, error) {
	params := map[string]string{
		"contractaddress": contractaddress,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetAccountERC20HoldingsOpts contains optional parameters
type GetAccountERC20HoldingsOpts struct {
	Page            *int
	Offset          *int
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetAccountERC20Holdings returns the ERC-20 tokens and amount held by an address
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetAccountERC20Holdings(ctx context.Context, address string, opts *GetAccountERC20HoldingsOpts) ([]RespERC20Holding, error) {
	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "100",
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
		}
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

// GetAccountNFTHoldingsOpts contains optional parameters
type GetAccountNFTHoldingsOpts struct {
	Page            *int
	Offset          *int
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetAccountNFTHoldings returns the ERC-721 tokens and amount held by an address
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetAccountNFTHoldings(ctx context.Context, address string, opts *GetAccountNFTHoldingsOpts) ([]RespNFTHolding, error) {
	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "100",
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
		}
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

// GetAccountNFTInventoriesOpts contains optional parameters
type GetAccountNFTInventoriesOpts struct {
	Page            *int
	Offset          *int
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetAccountNFTInventories returns the ERC-721 token inventory of an address, filtered by contract address
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetAccountNFTInventories(ctx context.Context, address, contractaddress string, opts *GetAccountNFTInventoriesOpts) ([]RespNFTTokenInventory, error) {
	params := map[string]string{
		"address":         address,
		"contractaddress": contractaddress,
		"page":            "1",
		"offset":          "100",
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
		}
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
