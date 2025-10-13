package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Block Module
// ============================================================================

// GetBlockAndUncleRewardsOpts contains optional parameters for GetBlockAndUncleRewards
type GetBlockAndUncleRewardsOpts struct {
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

// GetBlockAndUncleRewards returns the block reward and uncle block rewards for a given block number
//
// This endpoint returns detailed information about block rewards and uncle block rewards
// for a specific block. This is useful for analyzing miner rewards and network economics.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - blockno: The block number to check block rewards for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespBlockReward: Block reward information including block reward and uncle rewards
//   - error: Error if the request fails
//
// Example:
//
//	// Get block rewards for a specific block
//	blockNumber := 18000000
//	reward, err := client.GetBlockAndUncleRewards(ctx, blockNumber, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Block Number: %s\n", reward.BlockNumber)
//	fmt.Printf("Time Stamp: %s\n", reward.TimeStamp)
//	fmt.Printf("Block Miner: %s\n", reward.BlockMiner)
//	fmt.Printf("Block Reward: %s\n", reward.BlockReward)
//	fmt.Printf("Uncle Inclusion Reward: %s\n", reward.UncleInclusionReward)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	reward, err := client.GetBlockAndUncleRewards(ctx, blockNumber, &GetBlockAndUncleRewardsOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Returns empty RespBlockReward if block not found
//   - BlockReward is in wei (smallest unit)
//   - UncleInclusionReward is the reward for including uncle blocks
//   - Useful for analyzing miner economics and network health
func (c *HTTPClient) GetBlockAndUncleRewards(ctx context.Context, blockno int, opts *GetBlockAndUncleRewardsOpts) (*RespBlockReward, error) {
	params := map[string]string{
		"blockno": strconv.Itoa(blockno),
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
		module:          "block",
		action:          "getblockreward",
		params:          params,
		noFoundReturn:   RespBlockReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespBlockReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBlockTransactionsCountOpts contains optional parameters
type GetBlockTransactionsCountOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetBlockTransactionsCount returns the number of transactions in a specified block
//
// Note: Only supported on Ethereum mainnet (chainid=1)
//
// Example:
//
//	count, err := client.GetBlockTransactionsCount(ctx, 12345678, nil)
func (c *HTTPClient) GetBlockTransactionsCount(ctx context.Context, blockno int, opts *GetBlockTransactionsCountOpts) (*RespBlockTxsCountByBlockNo, error) {
	params := map[string]string{
		"blockno": strconv.Itoa(blockno),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			if *opts.ChainID != 1 {
				return nil, fmt.Errorf("this endpoint is only supported on Ethereum mainnet (chainid=1)")
			}
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "block",
		action:          "getblocktxnscount",
		params:          params,
		noFoundReturn:   RespBlockTxsCountByBlockNo{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespBlockTxsCountByBlockNo
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBlockCountdownTimeOpts contains optional parameters
type GetBlockCountdownTimeOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetBlockCountdownTime returns estimated time remaining until a future block is mined
//
// Example:
//
//	countdown, err := client.GetBlockCountdownTime(ctx, 20000000, nil)
func (c *HTTPClient) GetBlockCountdownTime(ctx context.Context, blockno int, opts *GetBlockCountdownTimeOpts) (*RespEstimateBlockCountdownTimeByBlockNo, error) {
	params := map[string]string{
		"blockno": strconv.Itoa(blockno),
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
		module:          "block",
		action:          "getblockcountdown",
		params:          params,
		noFoundReturn:   RespEstimateBlockCountdownTimeByBlockNo{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEstimateBlockCountdownTimeByBlockNo
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBlockNumberByTimestampOpts contains optional parameters
type GetBlockNumberByTimestampOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetBlockNumberByTimestamp returns the block number that was mined at a certain timestamp
//
// Args:
//   - timestamp: Unix timestamp in seconds
//   - closest: "before" or "after"
//
// Example:
//
//	blockNo, err := client.GetBlockNumberByTimestamp(ctx, 1620000000, "before", nil)
func (c *HTTPClient) GetBlockNumberByTimestamp(ctx context.Context, timestamp int64, closest string, opts *GetBlockNumberByTimestampOpts) (int, error) {
	params := map[string]string{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"closest":   closest,
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
		module:          "block",
		action:          "getblocknobytime",
		params:          params,
		noFoundReturn:   -1,
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return 0, err
	}

	// Handle different return types
	switch v := data.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("unexpected type for block number: %T", data)
	}
}

// GetDailyAvgBlockSizesOpts contains optional parameters
type GetDailyAvgBlockSizesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgBlockSizes returns daily average block size within a date range
//
// Args:
//   - startdate: Starting date in yyyy-MM-dd format (e.g. "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g. "2019-02-28")
//
// Example:
//
//	sizes, err := client.GetDailyAvgBlockSizes(ctx, "2019-02-01", "2019-02-28", nil)
func (c *HTTPClient) GetDailyAvgBlockSizes(ctx context.Context, startdate, enddate string, opts *GetDailyAvgBlockSizesOpts) ([]RespDailyAvgBlockSize, error) {
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
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavgblocksize",
		params:          params,
		noFoundReturn:   []RespDailyAvgBlockSize{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgBlockSize
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Additional daily stats methods follow the same pattern...
// Continuing with remaining block module methods in similar fashion
