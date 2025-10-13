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
	ChainID *int64

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
func (c *HTTPClient) GetBlockAndUncleRewards(ctx context.Context, blockno int64, opts *GetBlockAndUncleRewardsOpts) (*RespBlockReward, error) {
	params := map[string]string{
		"blockno": strconv.FormatInt(blockno, 10),
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

// GetBlockTransactionsCountOpts contains optional parameters for GetBlockTransactionsCount
type GetBlockTransactionsCountOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Note: Only supported on Ethereum mainnet (chainid=1)
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetBlockTransactionsCount returns the number of transactions in a specified block
//
// This endpoint returns the number of transactions in a specified block, including
// normal transactions, internal transactions, and various token transfers. This is
// useful for analyzing block activity and transaction volume.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - blockno: The block number to get transaction count for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespBlockTxsCountByBlockNo: Transaction count details including different transaction types
//   - error: Error if the request fails or if chainid is not 1
//
// Example:
//
//	// Get transaction count for a specific block
//	blockNumber := 18000000
//	count, err := client.GetBlockTransactionsCount(ctx, blockNumber, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Block Number: %s\n", count.Block)
//	fmt.Printf("Normal Transactions: %s\n", count.TxsCount)
//	fmt.Printf("Internal Transactions: %s\n", count.InternalTxsCount)
//	fmt.Printf("ERC20 Transfers: %s\n", count.ERC20TxsCount)
//	fmt.Printf("ERC721 Transfers: %s\n", count.ERC721TxsCount)
//	fmt.Printf("ERC1155 Transfers: %s\n", count.ERC1155TxsCount)
//
// Note:
//   - Only supported on Ethereum mainnet (chainid=1)
//   - Returns detailed breakdown of different transaction types
//   - Useful for analyzing block activity and transaction volume
func (c *HTTPClient) GetBlockTransactionsCount(ctx context.Context, blockno int64, opts *GetBlockTransactionsCountOpts) (*RespBlockTxsCountByBlockNo, error) {
	params := map[string]string{
		"blockno": strconv.FormatInt(blockno, 10),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			if *opts.ChainID != 1 {
				return nil, fmt.Errorf("this endpoint is only supported on Ethereum mainnet (chainid=1)")
			}
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
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

// GetBlockCountdownTimeOpts contains optional parameters for GetBlockCountdownTime
type GetBlockCountdownTimeOpts struct {
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

// GetBlockCountdownTime returns estimated time remaining until a future block is mined
//
// This endpoint returns the estimated time remaining until a future block is mined.
// This is useful for predicting when specific blocks will be mined and planning
// transactions or operations around block timing.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - blockno: The future block number to estimate countdown for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEstimateBlockCountdownTimeByBlockNo: Countdown details including estimated time
//   - error: Error if the request fails
//
// Example:
//
//	// Get countdown time for a future block
//	futureBlock := 20000000
//	countdown, err := client.GetBlockCountdownTime(ctx, futureBlock, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Current Block: %s\n", countdown.CurrentBlock)
//	fmt.Printf("Target Block: %s\n", countdown.CountdownBlock)
//	fmt.Printf("Remaining Blocks: %s\n", countdown.RemainingBlock)
//	fmt.Printf("Estimated Time: %s seconds\n", countdown.EstimateTimeInSec)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	countdown, err := client.GetBlockCountdownTime(ctx, futureBlock, &GetBlockCountdownTimeOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Only works for future blocks (block number > current block)
//   - Estimate is based on current network block time
//   - Useful for planning transactions and operations
func (c *HTTPClient) GetBlockCountdownTime(ctx context.Context, blockno int64, opts *GetBlockCountdownTimeOpts) (*RespEstimateBlockCountdownTimeByBlockNo, error) {
	params := map[string]string{
		"blockno": strconv.FormatInt(blockno, 10),
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
	ChainID         *int64
	OnLimitExceeded *RateLimitBehavior
}

// GetBlockNumberByTimestamp returns the block number that was mined at a certain timestamp
//
// This endpoint returns the block number that was mined at a certain timestamp.
// This is useful for finding blocks by time and analyzing historical data.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - timestamp: Unix timestamp in seconds
//   - closest: Return closest block "before" or "after" the timestamp
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - int: Block number that was mined closest to the provided timestamp
//   - error: Error if the request fails
//
// Example:
//
//	// Get block number by timestamp
//	timestamp := int64(1620000000) // Unix timestamp
//	blockNo, err := client.GetBlockNumberByTimestamp(ctx, timestamp, "before", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Block number: %d\n", blockNo)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	blockNo, err := client.GetBlockNumberByTimestamp(ctx, timestamp, "after", &GetBlockNumberByTimestampOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Returns -1 if no block found
//   - Timestamp must be in Unix seconds format
//   - "before" returns the latest block before the timestamp
//   - "after" returns the earliest block after the timestamp
func (c *HTTPClient) GetBlockNumberByTimestamp(ctx context.Context, timestamp int64, closest string, opts *GetBlockNumberByTimestampOpts) (int, error) {
	params := map[string]string{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"closest":   closest,
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

// GetDailyAvgBlockSizesOpts contains optional parameters for GetDailyAvgBlockSizes
type GetDailyAvgBlockSizesOpts struct {
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

// GetDailyAvgBlockSizes returns daily average block size within a date range
//
// This endpoint returns daily average block size within a specified date range.
// This is useful for analyzing network capacity and block size trends over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startdate: Starting date in yyyy-MM-dd format (e.g., "2019-02-01")
//   - enddate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgBlockSize: List of daily average block sizes
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average block sizes for a date range
//	startDate := "2019-02-01"
//	endDate := "2019-02-28"
//	sizes, err := client.GetDailyAvgBlockSizes(ctx, startDate, endDate, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, size := range sizes {
//	    fmt.Printf("Date: %s, Block Size: %s bytes\n", size.UTCDate, size.BlockSizeBytes)
//	}
//
//	// With custom sort order
//	sizes, err := client.GetDailyAvgBlockSizes(ctx, startDate, endDate, &GetDailyAvgBlockSizesOpts{
//	    Sort: &[]string{"desc"}[0],
//	})
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
//   - Block size is returned in bytes
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
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
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
