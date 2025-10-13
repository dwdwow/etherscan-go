package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Stats Module - Daily Statistics
// ============================================================================

// GetDailyBlockCountRewardsOpts contains optional parameters for GetDailyBlockCountRewards
type GetDailyBlockCountRewardsOpts struct {
	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by date in ascending order (oldest first)
	//   - "desc": Sort by date in descending order (newest first)
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyBlockCountRewards returns daily block count and rewards within a date range
//
// This endpoint returns daily block count and rewards data within a specified date range.
// This is useful for analyzing network activity, miner rewards, and blockchain statistics
// over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyBlockCountReward: List of daily block count and rewards data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily block count and rewards for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	rewards, err := client.GetDailyBlockCountRewards(ctx, startDate, endDate, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, reward := range rewards {
//	    fmt.Printf("Date: %s\n", reward.UTCDate)
//	    fmt.Printf("Block Count: %s\n", reward.BlockCount)
//	    fmt.Printf("Block Rewards: %s ETH\n", reward.BlockRewards_Eth)
//	}
//
//	// With custom sort order
//	sort := "desc"
//	rewards, err := client.GetDailyBlockCountRewards(ctx, startDate, endDate, &GetDailyBlockCountRewardsOpts{
//	    Sort: &sort,
//	})
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
//   - Useful for analyzing network activity over time
func (c *HTTPClient) GetDailyBlockCountRewards(ctx context.Context, startdate, enddate string, opts *GetDailyBlockCountRewardsOpts) ([]RespDailyBlockCountReward, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyblkcount",
		params:          params,
		noFoundReturn:   []RespDailyBlockCountReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyBlockCountReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyBlockRewardsOpts contains optional parameters for GetDailyBlockRewards
type GetDailyBlockRewardsOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyBlockRewards returns daily block rewards distributed to miners within a date range
//
// This endpoint returns daily block rewards distributed to miners within a specified date range.
// This is useful for analyzing miner rewards and network economics over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyBlockReward: List of daily block rewards data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily block rewards for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	rewards, err := client.GetDailyBlockRewards(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyBlockRewards(ctx context.Context, startdate, enddate string, opts *GetDailyBlockRewardsOpts) ([]RespDailyBlockReward, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyblockrewards",
		params:          params,
		noFoundReturn:   []RespDailyBlockReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyBlockReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgBlockTimeOpts contains optional parameters for GetDailyAvgBlockTime
type GetDailyAvgBlockTimeOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgBlockTime returns daily average time for a block to be included in the blockchain
//
// This endpoint returns the daily average time for a block to be included in the blockchain
// within a specified date range. This is useful for analyzing network performance and
// block production efficiency over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgTimeBlockMined: List of daily average block times
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average block time for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	blockTimes, err := client.GetDailyAvgBlockTime(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyAvgBlockTime(ctx context.Context, startdate, enddate string, opts *GetDailyAvgBlockTimeOpts) ([]RespDailyAvgTimeBlockMined, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavgblocktime",
		params:          params,
		noFoundReturn:   []RespDailyAvgTimeBlockMined{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgTimeBlockMined
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyUncleBlockCountAndRewardsOpts contains optional parameters for GetDailyUncleBlockCountAndRewards
type GetDailyUncleBlockCountAndRewardsOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyUncleBlockCountAndRewards returns daily uncle block count and rewards
//
// This endpoint returns daily uncle block count and rewards within a specified date range.
// Uncle blocks are blocks that were mined but not included in the main chain.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyUncleBlockCountAndReward: List of daily uncle block stats
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily uncle block stats for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	uncles, err := client.GetDailyUncleBlockCountAndRewards(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyUncleBlockCountAndRewards(ctx context.Context, startdate, enddate string, opts *GetDailyUncleBlockCountAndRewardsOpts) ([]RespDailyUncleBlockCountAndReward, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyuncleblkcount",
		params:          params,
		noFoundReturn:   []RespDailyUncleBlockCountAndReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyUncleBlockCountAndReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Stats Module - Supply and Price
// ============================================================================

// GetTotalEthSupplyOpts contains optional parameters
type GetTotalEthSupplyOpts struct {
	ChainID         *int64
	OnLimitExceeded *RateLimitBehavior
}

// GetTotalEthSupply returns the current amount of Eth in circulation excluding ETH2 Staking rewards and EIP1559 burnt fees
//
// This endpoint returns the current amount of ETH in circulation, excluding ETH2 staking
// rewards and EIP-1559 burnt fees. This represents the traditional ETH supply.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Current ETH supply in wei
//   - error: Error if the request fails
//
// Example:
//
//	// Get total ETH supply
//	supply, err := client.GetTotalEthSupply(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total ETH supply: %s wei\n", supply)
//
// Note:
//   - Returns supply in wei (smallest unit)
//   - Excludes ETH2 staking rewards and EIP-1559 burnt fees
func (c *HTTPClient) GetTotalEthSupply(ctx context.Context, opts *GetTotalEthSupplyOpts) (string, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "ethsupply",
		params:          params,
		noFoundReturn:   "",
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

// GetTotalEth2SupplyOpts contains optional parameters
type GetTotalEth2SupplyOpts struct {
	ChainID         *int64
	OnLimitExceeded *RateLimitBehavior
}

// GetTotalEth2Supply returns the current amount of Eth in circulation, ETH2 Staking rewards, EIP1559 burnt fees
//
// This endpoint returns the current amount of ETH in circulation, including ETH2 staking
// rewards and EIP-1559 burnt fees. This represents the total ETH supply including all sources.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Total ETH supply in wei
//   - error: Error if the request fails
//
// Example:
//
//	// Get total ETH2 supply
//	supply, err := client.GetTotalEth2Supply(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total ETH2 supply: %s wei\n", supply)
//
// Note:
//   - Returns supply in wei (smallest unit)
//   - Includes ETH2 staking rewards and EIP-1559 burnt fees
func (c *HTTPClient) GetTotalEth2Supply(ctx context.Context, opts *GetTotalEth2SupplyOpts) (string, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "ethsupply2",
		params:          params,
		noFoundReturn:   "",
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

// GetEthPriceOpts contains optional parameters for GetEthPrice
type GetEthPriceOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetEthPrice returns the latest price of the native/gas token
//
// This endpoint returns the latest price of the native/gas token (ETH on Ethereum,
// MATIC on Polygon, etc.) in USD. This is useful for getting current token price
// information for financial calculations.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthPrice: Token price information including USD price and market cap
//   - error: Error if the request fails
//
// Example:
//
//	// Get current ETH price
//	price, err := client.GetEthPrice(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("ETH Price: $%s\n", price.EthUsd)
//	fmt.Printf("Market Cap: $%s\n", price.EthUsdMarketCap)
//
// Note:
//   - Returns current token price in USD
//   - Includes market cap information
func (c *HTTPClient) GetEthPrice(ctx context.Context, opts *GetEthPriceOpts) (*RespEthPrice, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "ethprice",
		params:          params,
		noFoundReturn:   RespEthPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEthHistoricalPricesOpts contains optional parameters for GetEthHistoricalPrices
type GetEthHistoricalPricesOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetEthHistoricalPrices returns the historical price of 1 ETH
//
// This endpoint returns the historical price of 1 ETH within a specified date range.
// This is useful for analyzing price trends and historical market data.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespEthHistoricalPrice: List of historical ETH prices
//   - error: Error if the request fails
//
// Example:
//
//	// Get historical ETH prices for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	prices, err := client.GetEthHistoricalPrices(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetEthHistoricalPrices(ctx context.Context, startdate, enddate string, opts *GetEthHistoricalPricesOpts) ([]RespEthHistoricalPrice, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "ethdailyprice",
		params:          params,
		noFoundReturn:   []RespEthHistoricalPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEthHistoricalPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Stats Module - Network Statistics
// ============================================================================

// GetEthereumNodesSizeOpts contains optional parameters
type GetEthereumNodesSizeOpts struct {
	ChainID         *int64
	OnLimitExceeded *RateLimitBehavior
}

// GetEthereumNodesSize returns the size of the Ethereum blockchain, in bytes, over a date range
//
// This endpoint returns the size of the Ethereum blockchain in bytes over a specified date range.
// This is useful for analyzing network growth and storage requirements over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - clienttype: Client type (e.g., "geth", "parity")
//   - syncmode: Sync mode (e.g., "default", "archive")
//   - sort: Sort order ("asc" or "desc")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespEtheumNodeSize: List of network size data
//   - error: Error if the request fails
//
// Example:
//
//	// Get Ethereum network size for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	clientType := "geth"
//	syncMode := "default"
//	sort := "asc"
//	sizes, err := client.GetEthereumNodesSize(ctx, startDate, endDate, clientType, syncMode, sort, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetEthereumNodesSize(ctx context.Context, startdate, enddate, clienttype, syncmode, sort string, opts *GetEthereumNodesSizeOpts) ([]RespEtheumNodeSize, error) {
	params := map[string]string{
		"startdate":  startdate,
		"enddate":    enddate,
		"clienttype": clienttype,
		"syncmode":   syncmode,
		"sort":       sort,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "chainsize",
		params:          params,
		noFoundReturn:   []RespEtheumNodeSize{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEtheumNodeSize
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeCountOpts contains optional parameters for GetNodeCount
type GetNodeCountOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetNodeCount returns the total number of discoverable Ethereum nodes
//
// This endpoint returns the total number of discoverable Ethereum nodes. This is useful
// for analyzing network health and decentralization metrics.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespNodeCount: Node count information
//   - error: Error if the request fails
//
// Example:
//
//	// Get total node count
//	nodeCount, err := client.GetNodeCount(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total nodes: %s\n", nodeCount.TotalNodeCount)
//
// Note:
//   - Returns total number of discoverable nodes
//   - Useful for network health analysis
func (c *HTTPClient) GetNodeCount(ctx context.Context, opts *GetNodeCountOpts) (*RespNodeCount, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "nodecount",
		params:          params,
		noFoundReturn:   RespNodeCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespNodeCount
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDailyTxFeesOpts contains optional parameters for GetDailyTxFees
type GetDailyTxFeesOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyTxFees returns the amount of transaction fees paid to miners per day
//
// This endpoint returns the amount of transaction fees paid to miners per day within
// a specified date range. This is useful for analyzing network economics and miner rewards.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyTxFee: List of daily transaction fees
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily transaction fees for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	fees, err := client.GetDailyTxFees(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyTxFees(ctx context.Context, startdate, enddate string, opts *GetDailyTxFeesOpts) ([]RespDailyTxFee, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailytxnfee",
		params:          params,
		noFoundReturn:   []RespDailyTxFee{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyTxFee
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyNewAddressesOpts contains optional parameters for GetDailyNewAddresses
type GetDailyNewAddressesOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyNewAddresses returns the number of new Ethereum addresses created per day
//
// This endpoint returns the number of new Ethereum addresses created per day within
// a specified date range. This is useful for analyzing network adoption and growth.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyNewAddress: List of daily new address counts
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily new addresses for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	addresses, err := client.GetDailyNewAddresses(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyNewAddresses(ctx context.Context, startdate, enddate string, opts *GetDailyNewAddressesOpts) ([]RespDailyNewAddress, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailynewaddress",
		params:          params,
		noFoundReturn:   []RespDailyNewAddress{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyNewAddress
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyNetworkUtilizationsOpts contains optional parameters for GetDailyNetworkUtilizations
type GetDailyNetworkUtilizationsOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyNetworkUtilizations returns the daily average gas used over gas limit, in percentage
//
// This endpoint returns the daily average gas used over gas limit in percentage within
// a specified date range. This is useful for analyzing network congestion and utilization.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyNetworkUtilization: List of daily network utilization data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily network utilization for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	utilizations, err := client.GetDailyNetworkUtilizations(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyNetworkUtilizations(ctx context.Context, startdate, enddate string, opts *GetDailyNetworkUtilizationsOpts) ([]RespDailyNetworkUtilization, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailynetutilization",
		params:          params,
		noFoundReturn:   []RespDailyNetworkUtilization{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyNetworkUtilization
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgHashratesOpts contains optional parameters for GetDailyAvgHashrates
type GetDailyAvgHashratesOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgHashrates returns the historical measure of processing power of the Ethereum network
//
// This endpoint returns the historical measure of processing power of the Ethereum network
// within a specified date range. This is useful for analyzing network security and mining activity.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgHashrate: List of daily average hashrate data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average hashrates for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	hashrates, err := client.GetDailyAvgHashrates(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyAvgHashrates(ctx context.Context, startdate, enddate string, opts *GetDailyAvgHashratesOpts) ([]RespDailyAvgHashrate, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavghashrate",
		params:          params,
		noFoundReturn:   []RespDailyAvgHashrate{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgHashrate
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyTxCountsOpts contains optional parameters for GetDailyTxCounts
type GetDailyTxCountsOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyTxCounts returns the number of transactions performed on the Ethereum blockchain per day
//
// This endpoint returns the number of transactions performed on the Ethereum blockchain
// per day within a specified date range. This is useful for analyzing network activity and usage.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyTxCount: List of daily transaction counts
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily transaction counts for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	txCounts, err := client.GetDailyTxCounts(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyTxCounts(ctx context.Context, startdate, enddate string, opts *GetDailyTxCountsOpts) ([]RespDailyTxCount, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailytx",
		params:          params,
		noFoundReturn:   []RespDailyTxCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyTxCount
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgDifficultiesOpts contains optional parameters for GetDailyAvgDifficulties
type GetDailyAvgDifficultiesOpts struct {
	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgDifficulties returns the historical mining difficulty of the Ethereum network
//
// This endpoint returns the historical mining difficulty of the Ethereum network within
// a specified date range. This is useful for analyzing network security and mining complexity.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgDifficulty: List of daily average difficulty data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average difficulties for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	difficulties, err := client.GetDailyAvgDifficulties(ctx, startDate, endDate, nil)
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
func (c *HTTPClient) GetDailyAvgDifficulties(ctx context.Context, startdate, enddate string, opts *GetDailyAvgDifficultiesOpts) ([]RespDailyAvgDifficulty, error) {
	params := map[string]string{
		"startdate": startdate,
		"enddate":   enddate,
		"sort":      "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Sort != nil {
			params["sort"] = *opts.Sort
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavgnetdifficulty",
		params:          params,
		noFoundReturn:   []RespDailyAvgDifficulty{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgDifficulty
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
