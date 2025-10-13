package etherscan

import (
	"context"
	"strconv"
)

// ============================================================================
// Logs Module
// ============================================================================

// GetEventLogsByAddressOpts contains optional parameters for GetEventLogsByAddress
type GetEventLogsByAddressOpts struct {
	// FromBlock is the starting block number to search for logs
	// If nil, searches from genesis block
	// Use this to limit the search range and improve performance
	FromBlock *int64

	// ToBlock is the ending block number to search for logs
	// If nil, searches to latest block
	// Use this to limit the search range and improve performance
	ToBlock *int64

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int64

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	// Higher values return more results per page but may be slower
	Offset *int64

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

// GetEventLogsByAddress returns event logs from an address, with optional filtering by block range
//
// This endpoint returns event logs emitted by a specific contract address. Event logs are
// generated when contracts emit events, which are useful for tracking contract state changes
// and interactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The contract address to get logs from
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespEventLogByAddress: List of event logs with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get all event logs from a contract
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	logs, err := client.GetEventLogsByAddress(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, log := range logs {
//	    fmt.Printf("Block: %s, Tx: %s, Topics: %v\n",
//	        log.BlockNumber, log.TransactionHash, log.Topics)
//	}
//
//	// Get logs from a specific block range
//	fromBlock := 18000000
//	toBlock := 18000100
//	logs, err := client.GetEventLogsByAddress(ctx, contractAddr, &GetEventLogsByAddressOpts{
//	    FromBlock: &fromBlock,
//	    ToBlock:   &toBlock,
//	})
//
//	// With pagination
//	page := 1
//	offset := 500
//	logs, err := client.GetEventLogsByAddress(ctx, contractAddr, &GetEventLogsByAddressOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	})
//
// Note:
//   - Returns empty slice if no logs found
//   - Maximum 1000 records per call
//   - Use Page and Offset for pagination
//   - Topics field contains event signature and indexed parameters
//   - Data field contains non-indexed event parameters
func (c *HTTPClient) GetEventLogsByAddress(ctx context.Context, address string, opts *GetEventLogsByAddressOpts) ([]RespEventLogByAddress, error) {
	params := map[string]string{
		"address": address,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.FromBlock != nil {
			params["fromBlock"] = strconv.FormatInt(*opts.FromBlock, 10)
		}
		if opts.ToBlock != nil {
			params["toBlock"] = strconv.FormatInt(*opts.ToBlock, 10)
		}
		if opts.Page != nil {
			params["page"] = strconv.FormatInt(*opts.Page, 10)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.FormatInt(*opts.Offset, 10)
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "logs",
		action:          "getLogs",
		params:          params,
		noFoundReturn:   []RespEventLogByAddress{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEventLogByAddress
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetEventLogsByTopicsOpts contains optional parameters for GetEventLogsByTopics
type GetEventLogsByTopicsOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int64

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	// Higher values return more results per page but may be slower
	Offset *int64

	// FromBlock is the starting block number to search for logs
	// If nil, searches from genesis block
	// Use this to limit the search range and improve performance
	FromBlock *int64

	// ToBlock is the ending block number to search for logs
	// If nil, searches to latest block
	// Use this to limit the search range and improve performance
	ToBlock *int64

	// Topic0 is the first topic to filter by (event signature)
	// This is typically the keccak256 hash of the event signature
	// Example: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" (Transfer event)
	Topic0 *string

	// Topic1 is the second topic to filter by (first indexed parameter)
	// This is typically an address or other indexed parameter
	Topic1 *string

	// Topic2 is the third topic to filter by (second indexed parameter)
	// This is typically an address or other indexed parameter
	Topic2 *string

	// Topic3 is the fourth topic to filter by (third indexed parameter)
	// This is typically an address or other indexed parameter
	Topic3 *string

	// Topic0_1_Opr is the operator between topic0 and topic1
	// Options: "and" or "or"
	// Default: "and"
	Topic0_1_Opr *string

	// Topic0_2_Opr is the operator between topic0 and topic2
	// Options: "and" or "or"
	// Default: "and"
	Topic0_2_Opr *string

	// Topic0_3_Opr is the operator between topic0 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic0_3_Opr *string

	// Topic1_2_Opr is the operator between topic1 and topic2
	// Options: "and" or "or"
	// Default: "and"
	Topic1_2_Opr *string

	// Topic1_3_Opr is the operator between topic1 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic1_3_Opr *string

	// Topic2_3_Opr is the operator between topic2 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic2_3_Opr *string

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

// GetEventLogsByTopics returns event logs filtered by topics in a block range
//
// This endpoint returns event logs filtered by event topics (signatures and indexed parameters).
// Topics are used to filter events based on their signature and indexed parameters. This is
// useful for finding specific types of events across the entire blockchain or within a block range.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Parameters for filtering logs by topics
//
// Returns:
//   - []RespEventLogByTopics: List of event logs matching the topic filters
//   - error: Error if the request fails
//
// Example:
//
//	// Get all Transfer events
//	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
//	logs, err := client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
//	    Topic0: &transferTopic,
//	})
//
//	// Get Transfer events from a specific address (from parameter)
//	fromAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	fromTopic := "0x000000000000000000000000" + fromAddr[2:] // Remove 0x and pad with zeros
//	logs, err := client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
//	    Topic0: &transferTopic,
//	    Topic1: &fromTopic,
//	})
//
//	// Get Transfer events within a block range
//	fromBlock := 18000000
//	toBlock := 18000100
//	logs, err := client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
//	    Topic0:    &transferTopic,
//	    FromBlock: &fromBlock,
//	    ToBlock:   &toBlock,
//	})
//
//	// Complex filtering with multiple topics and operators
//	toAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	toTopic := "0x000000000000000000000000" + toAddr[2:]
//	opr := "or"
//	logs, err := client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
//	    Topic0:      &transferTopic,
//	    Topic1:      &fromTopic,
//	    Topic2:      &toTopic,
//	    Topic0_1_Opr: &opr,
//	})
//
// Note:
//   - Topics must be 32-byte hex strings (64 characters + 0x prefix)
//   - Topic0 is typically the event signature hash
//   - Topic1-3 are indexed parameters, padded with zeros
//   - Use "and" or "or" operators to combine topic filters
//   - Maximum 1000 records per call
func (c *HTTPClient) GetEventLogsByTopics(ctx context.Context, opts *GetEventLogsByTopicsOpts) ([]RespEventLogByTopics, error) {
	params := map[string]string{
		"page":   "1",
		"offset": "1000",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Page != nil {
			params["page"] = strconv.FormatInt(*opts.Page, 10)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.FormatInt(*opts.Offset, 10)
		}
		if opts.FromBlock != nil {
			params["fromBlock"] = strconv.FormatInt(*opts.FromBlock, 10)
		}
		if opts.ToBlock != nil {
			params["toBlock"] = strconv.FormatInt(*opts.ToBlock, 10)
		}
		if opts.Topic0 != nil {
			params["topic0"] = *opts.Topic0
		}
		if opts.Topic1 != nil {
			params["topic1"] = *opts.Topic1
		}
		if opts.Topic2 != nil {
			params["topic2"] = *opts.Topic2
		}
		if opts.Topic3 != nil {
			params["topic3"] = *opts.Topic3
		}
		if opts.Topic0_1_Opr != nil {
			params["topic0_1_opr"] = *opts.Topic0_1_Opr
		}
		if opts.Topic0_2_Opr != nil {
			params["topic0_2_opr"] = *opts.Topic0_2_Opr
		}
		if opts.Topic0_3_Opr != nil {
			params["topic0_3_opr"] = *opts.Topic0_3_Opr
		}
		if opts.Topic1_2_Opr != nil {
			params["topic1_2_opr"] = *opts.Topic1_2_Opr
		}
		if opts.Topic1_3_Opr != nil {
			params["topic1_3_opr"] = *opts.Topic1_3_Opr
		}
		if opts.Topic2_3_Opr != nil {
			params["topic2_3_opr"] = *opts.Topic2_3_Opr
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "logs",
		action:          "getLogs",
		params:          params,
		noFoundReturn:   []RespEventLogByTopics{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEventLogByTopics
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetEventLogsByAddressFilteredByTopicsOpts contains optional parameters for GetEventLogsByAddressFilteredByTopics
type GetEventLogsByAddressFilteredByTopicsOpts struct {
	// Page number for pagination
	// Default: 1
	Page *int64

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	Offset *int64

	// FromBlock is the starting block number to search for logs
	// If nil, searches from genesis block
	FromBlock *int64

	// ToBlock is the ending block number to search for logs
	// If nil, searches to latest block
	ToBlock *int64

	// Topic0 is the first topic to filter by (event signature)
	Topic0 *string

	// Topic1 is the second topic to filter by (first indexed parameter)
	Topic1 *string

	// Topic2 is the third topic to filter by (second indexed parameter)
	Topic2 *string

	// Topic3 is the fourth topic to filter by (third indexed parameter)
	Topic3 *string

	// Topic0_1_Opr is the operator between topic0 and topic1
	// Options: "and" or "or"
	Topic0_1_Opr *string

	// Topic0_2_Opr is the operator between topic0 and topic2
	// Options: "and" or "or"
	Topic0_2_Opr *string

	// Topic0_3_Opr is the operator between topic0 and topic3
	// Options: "and" or "or"
	Topic0_3_Opr *string

	// Topic1_2_Opr is the operator between topic1 and topic2
	// Options: "and" or "or"
	Topic1_2_Opr *string

	// Topic1_3_Opr is the operator between topic1 and topic3
	// Options: "and" or "or"
	Topic1_3_Opr *string

	// Topic2_3_Opr is the operator between topic2 and topic3
	// Options: "and" or "or"
	Topic2_3_Opr *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID
	ChainID *int64

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior
	OnLimitExceeded *RateLimitBehavior
}

// GetEventLogsByAddressFilteredByTopics returns event logs from a specific address filtered by topics and block range
//
// This endpoint combines address filtering with topic filtering to get event logs from a specific
// contract address that match certain event signatures and parameters.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The contract address to get logs from
//   - opts: Optional parameters for filtering
//
// Returns:
//   - []RespEventLogByAddressFilteredByTopics: List of filtered event logs
//   - error: Error if the request fails
//
// Example:
//
//	// Get Transfer events from a specific contract
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
//	logs, err := client.GetEventLogsByAddressFilteredByTopics(ctx, contractAddr, &GetEventLogsByAddressFilteredByTopicsOpts{
//	    Topic0: &transferTopic,
//	})
//
// Note:
//   - Combines address and topic filtering
//   - Maximum 1000 records per call
//   - Use Page and Offset for pagination
func (c *HTTPClient) GetEventLogsByAddressFilteredByTopics(ctx context.Context, address string, opts *GetEventLogsByAddressFilteredByTopicsOpts) ([]RespEventLogByAddressFilteredByTopics, error) {
	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "1000",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Page != nil {
			params["page"] = strconv.FormatInt(*opts.Page, 10)
		}
		if opts.Offset != nil {
			params["offset"] = strconv.FormatInt(*opts.Offset, 10)
		}
		if opts.FromBlock != nil {
			params["fromBlock"] = strconv.FormatInt(*opts.FromBlock, 10)
		}
		if opts.ToBlock != nil {
			params["toBlock"] = strconv.FormatInt(*opts.ToBlock, 10)
		}
		if opts.Topic0 != nil {
			params["topic0"] = *opts.Topic0
		}
		if opts.Topic1 != nil {
			params["topic1"] = *opts.Topic1
		}
		if opts.Topic2 != nil {
			params["topic2"] = *opts.Topic2
		}
		if opts.Topic3 != nil {
			params["topic3"] = *opts.Topic3
		}
		if opts.Topic0_1_Opr != nil {
			params["topic0_1_opr"] = *opts.Topic0_1_Opr
		}
		if opts.Topic0_2_Opr != nil {
			params["topic0_2_opr"] = *opts.Topic0_2_Opr
		}
		if opts.Topic0_3_Opr != nil {
			params["topic0_3_opr"] = *opts.Topic0_3_Opr
		}
		if opts.Topic1_2_Opr != nil {
			params["topic1_2_opr"] = *opts.Topic1_2_Opr
		}
		if opts.Topic1_3_Opr != nil {
			params["topic1_3_opr"] = *opts.Topic1_3_Opr
		}
		if opts.Topic2_3_Opr != nil {
			params["topic2_3_opr"] = *opts.Topic2_3_Opr
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.FormatInt(*opts.ChainID, 10)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "logs",
		action:          "getLogs",
		params:          params,
		noFoundReturn:   []RespEventLogByAddressFilteredByTopics{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEventLogByAddressFilteredByTopics
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
