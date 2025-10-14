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
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	FromBlock int64 `default:"0" json:"fromblock"`

	// ToBlock is the ending block number to search for logs
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	ToBlock int64 `default:"999999999999" json:"toblock"`

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"1000" json:"offset"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
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
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"address": address,
	}

	if opts.FromBlock != 0 {
		params["fromBlock"] = strconv.FormatInt(opts.FromBlock, 10)
	}
	if opts.ToBlock != 999999999999 {
		params["toBlock"] = strconv.FormatInt(opts.ToBlock, 10)
	}
	if opts.Page != 1 {
		params["page"] = strconv.FormatInt(opts.Page, 10)
	}
	if opts.Offset != 1000 {
		params["offset"] = strconv.FormatInt(opts.Offset, 10)
	}
	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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
	Page int64 `default:"1" json:"page"`

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"1000" json:"offset"`

	// FromBlock is the starting block number to search for logs
	// If 0, searches from genesis block
	// Use this to limit the search range and improve performance
	FromBlock int64 `default:"0" json:"fromblock"`

	// ToBlock is the ending block number to search for logs
	// If 0, searches to latest block
	// Use this to limit the search range and improve performance
	ToBlock int64 `default:"999999999999" json:"toblock"`

	// Topic0 is the first topic to filter by (event signature)
	// This is typically the keccak256 hash of the event signature
	// Example: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" (Transfer event)
	Topic0 string `default:"" json:"topic0"`

	// Topic1 is the second topic to filter by (first indexed parameter)
	// This is typically an address or other indexed parameter
	Topic1 string `default:"" json:"topic1"`

	// Topic2 is the third topic to filter by (second indexed parameter)
	// This is typically an address or other indexed parameter
	Topic2 string `default:"" json:"topic2"`

	// Topic3 is the fourth topic to filter by (third indexed parameter)
	// This is typically an address or other indexed parameter
	Topic3 string `default:"" json:"topic3"`

	// Topic0_1_Opr is the operator between topic0 and topic1
	// Options: "and" or "or"
	// Default: "and"
	Topic0_1_Opr string `default:"and" json:"topic0_1_opr"`

	// Topic0_2_Opr is the operator between topic0 and topic2
	// Options: "and" or "or"
	// Default: "and"
	Topic0_2_Opr string `default:"and" json:"topic0_2_opr"`

	// Topic0_3_Opr is the operator between topic0 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic0_3_Opr string `default:"and" json:"topic0_3_opr"`

	// Topic1_2_Opr is the operator between topic1 and topic2
	// Options: "and" or "or"
	// Default: "and"
	Topic1_2_Opr string `default:"and" json:"topic1_2_opr"`

	// Topic1_3_Opr is the operator between topic1 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic1_3_Opr string `default:"and" json:"topic1_3_opr"`

	// Topic2_3_Opr is the operator between topic2 and topic3
	// Options: "and" or "or"
	// Default: "and"
	Topic2_3_Opr string `default:"and" json:"topic2_3_opr"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
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
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"page":   "1",
		"offset": "1000",
	}

	if opts != nil {
		if opts.Page != 0 {
			params["page"] = strconv.FormatInt(opts.Page, 10)
		}
		if opts.Offset != 0 {
			params["offset"] = strconv.FormatInt(opts.Offset, 10)
		}
		if opts.FromBlock != 0 {
			params["fromBlock"] = strconv.FormatInt(opts.FromBlock, 10)
		}
		if opts.ToBlock != 999999999999 {
			params["toBlock"] = strconv.FormatInt(opts.ToBlock, 10)
		}
		if opts.Topic0 != "" {
			params["topic0"] = opts.Topic0
		}
		if opts.Topic1 != "" {
			params["topic1"] = opts.Topic1
		}
		if opts.Topic2 != "" {
			params["topic2"] = opts.Topic2
		}
		if opts.Topic3 != "" {
			params["topic3"] = opts.Topic3
		}
		if opts.Topic0_1_Opr != "" {
			params["topic0_1_opr"] = opts.Topic0_1_Opr
		}
		if opts.Topic0_2_Opr != "" {
			params["topic0_2_opr"] = opts.Topic0_2_Opr
		}
		if opts.Topic0_3_Opr != "" {
			params["topic0_3_opr"] = opts.Topic0_3_Opr
		}
		if opts.Topic1_2_Opr != "" {
			params["topic1_2_opr"] = opts.Topic1_2_Opr
		}
		if opts.Topic1_3_Opr != "" {
			params["topic1_3_opr"] = opts.Topic1_3_Opr
		}
		if opts.Topic2_3_Opr != "" {
			params["topic2_3_opr"] = opts.Topic2_3_Opr
		}
		if opts.ChainID != 0 {
			params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
		}
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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
	Page int64 `default:"1" json:"page"`

	// Offset is the number of records per page
	// Default: 1000
	// Maximum: 1000
	Offset int64 `default:"1000" json:"offset"`

	// FromBlock is the starting block number to search for logs
	// If 0, searches from genesis block
	FromBlock int64 `default:"0" json:"fromblock"`

	// ToBlock is the ending block number to search for logs
	// If 0, searches to latest block
	ToBlock int64 `default:"999999999999" json:"toblock"`

	// Topic0 is the first topic to filter by (event signature)
	Topic0 string `default:"" json:"topic0"`

	// Topic1 is the second topic to filter by (first indexed parameter)
	Topic1 string `default:"" json:"topic1"`

	// Topic2 is the third topic to filter by (second indexed parameter)
	Topic2 string `default:"" json:"topic2"`

	// Topic3 is the fourth topic to filter by (third indexed parameter)
	Topic3 string `default:"" json:"topic3"`

	// Topic0_1_Opr is the operator between topic0 and topic1
	// Options: "and" or "or"
	Topic0_1_Opr string `default:"and" json:"topic0_1_opr"`

	// Topic0_2_Opr is the operator between topic0 and topic2
	// Options: "and" or "or"
	Topic0_2_Opr string `default:"and" json:"topic0_2_opr"`

	// Topic0_3_Opr is the operator between topic0 and topic3
	// Options: "and" or "or"
	Topic0_3_Opr string `default:"and" json:"topic0_3_opr"`

	// Topic1_2_Opr is the operator between topic1 and topic2
	// Options: "and" or "or"
	Topic1_2_Opr string `default:"and" json:"topic1_2_opr"`

	// Topic1_3_Opr is the operator between topic1 and topic3
	// Options: "and" or "or"
	Topic1_3_Opr string `default:"and" json:"topic1_3_opr"`

	// Topic2_3_Opr is the operator between topic2 and topic3
	// Options: "and" or "or"
	Topic2_3_Opr string `default:"and" json:"topic2_3_opr"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
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
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"address": address,
		"page":    "1",
		"offset":  "1000",
	}

	if opts != nil {
		if opts.Page != 0 {
			params["page"] = strconv.FormatInt(opts.Page, 10)
		}
		if opts.Offset != 0 {
			params["offset"] = strconv.FormatInt(opts.Offset, 10)
		}
		if opts.FromBlock != 0 {
			params["fromBlock"] = strconv.FormatInt(opts.FromBlock, 10)
		}
		if opts.ToBlock != 999999999999 {
			params["toBlock"] = strconv.FormatInt(opts.ToBlock, 10)
		}
		if opts.Topic0 != "" {
			params["topic0"] = opts.Topic0
		}
		if opts.Topic1 != "" {
			params["topic1"] = opts.Topic1
		}
		if opts.Topic2 != "" {
			params["topic2"] = opts.Topic2
		}
		if opts.Topic3 != "" {
			params["topic3"] = opts.Topic3
		}
		if opts.Topic0_1_Opr != "" {
			params["topic0_1_opr"] = opts.Topic0_1_Opr
		}
		if opts.Topic0_2_Opr != "" {
			params["topic0_2_opr"] = opts.Topic0_2_Opr
		}
		if opts.Topic0_3_Opr != "" {
			params["topic0_3_opr"] = opts.Topic0_3_Opr
		}
		if opts.Topic1_2_Opr != "" {
			params["topic1_2_opr"] = opts.Topic1_2_Opr
		}
		if opts.Topic1_3_Opr != "" {
			params["topic1_3_opr"] = opts.Topic1_3_Opr
		}
		if opts.Topic2_3_Opr != "" {
			params["topic2_3_opr"] = opts.Topic2_3_Opr
		}
		if opts.ChainID != 0 {
			params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
		}
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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
