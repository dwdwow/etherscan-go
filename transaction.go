package etherscan

import (
	"context"
	"strconv"
)

// GetNormalTransactionsOpts contains optional parameters for GetNormalTransactions
type GetNormalTransactionsOpts struct {
	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	StartBlock *int

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	EndBlock *int

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset *int

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort *string

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

// GetNormalTransactions returns list of 'Normal' transactions by address
//
// This endpoint returns a list of normal (external) transactions performed by an address.
// Normal transactions are transactions that are sent from an externally owned account (EOA)
// to another account or contract.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to get transactions for (must be a valid Ethereum address)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespNormalTx: List of normal transactions with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get recent transactions for an address
//	txs, err := client.GetNormalTransactions(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("Hash: %s, From: %s, To: %s, Value: %s wei\n",
//	        tx.Hash, tx.From, tx.To, tx.Value)
//	}
//
//	// Get transactions with pagination
//	page := 1
//	offset := 50
//	sort := "desc"
//	txs, err := client.GetNormalTransactions(ctx, address, &GetNormalTransactionsOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	    Sort:   &sort,
//	})
//
//	// Get transactions from a specific block range
//	startBlock := 18000000
//	endBlock := 18000100
//	txs, err := client.GetNormalTransactions(ctx, address, &GetNormalTransactionsOpts{
//	    StartBlock: &startBlock,
//	    EndBlock:   &endBlock,
//	})
//
// Note:
//   - This endpoint returns max 10000 records per call
//   - Use StartBlock and EndBlock parameters to paginate through larger ranges
//   - For free tier: max 1000 records, for paid tier: max 5000 records
//   - Transactions include both ETH transfers and contract interactions
//   - Internal transactions (contract-to-contract) are not included (use GetInternalTransactionsByAddress)
//   - All values are returned as strings in Wei
func (c *HTTPClient) GetNormalTransactions(ctx context.Context, address string, opts *GetNormalTransactionsOpts) ([]RespNormalTx, error) {
	params := map[string]string{
		"address":    address,
		"startblock": "0",
		"endblock":   "999999999999",
		"page":       "1",
		"offset":     "100",
		"sort":       "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.StartBlock != nil {
			params["startblock"] = strconv.Itoa(*opts.StartBlock)
		}
		if opts.EndBlock != nil {
			params["endblock"] = strconv.Itoa(*opts.EndBlock)
		}
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
		action:          "txlist",
		params:          params,
		noFoundReturn:   []RespNormalTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespNormalTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBridgeTransactionsOpts contains optional parameters for GetBridgeTransactions
type GetBridgeTransactionsOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset *int

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Note: This endpoint is only applicable to specific chains
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetBridgeTransactions returns bridge transactions for an address
//
// This endpoint returns bridge transactions that involve the specified address.
// Bridge transactions are cross-chain transfers that move assets between different
// blockchain networks.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to get bridge transactions for (must be a valid address)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespBridgeTx: List of bridge transactions with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get bridge transactions for an address
//	txs, err := client.GetBridgeTransactions(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("Hash: %s, Amount: %s, Token: %s\n",
//	        tx.Hash, tx.Amount, tx.TokenName)
//	}
//
//	// With pagination
//	page := 1
//	offset := 50
//	txs, err := client.GetBridgeTransactions(ctx, address, &GetBridgeTransactionsOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	})
//
// Note:
//   - This endpoint is only applicable to specific blockchain networks:
//   - Gnosis (ChainID: 100)
//   - BitTorrent Chain (ChainID: 199)
//   - Polygon (ChainID: 137)
//   - Returns empty list for other networks
//   - Bridge transactions show cross-chain asset movements
//   - All amounts are returned as strings in the smallest unit of the token
func (c *HTTPClient) GetBridgeTransactions(ctx context.Context, address string, opts *GetBridgeTransactionsOpts) ([]RespBridgeTx, error) {
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
		action:          "txnbridge",
		params:          params,
		noFoundReturn:   []RespBridgeTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespBridgeTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Transaction Module
// ============================================================================

// GetContractExecutionStatusOpts contains optional parameters for GetContractExecutionStatus
type GetContractExecutionStatusOpts struct {
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

// GetContractExecutionStatus returns the status code of a contract execution
//
// This endpoint returns the execution status of a contract transaction.
// It provides information about whether the contract execution was successful
// or failed, along with error details if applicable.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txhash: The transaction hash to check the execution status for
//     Must be a valid transaction hash (0x... format)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - *RespContractExecutionStatus: Contract execution status information
//   - error: Error if the request fails
//
// Example:
//
//	// Check contract execution status
//	status, err := client.GetContractExecutionStatus(ctx, "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if status.IsError == "1" {
//	    fmt.Printf("Contract execution failed: %s\n", status.ErrDescription)
//	} else {
//	    fmt.Println("Contract execution successful")
//	}
//
//	// With custom chain
//	chainID := etherscan.PolygonMainnet
//	status, err := client.GetContractExecutionStatus(ctx, txhash, &GetContractExecutionStatusOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - IsError field: "0" = success, "1" = failed
//   - ErrDescription field: Contains error message if execution failed
//   - Only applicable to contract transactions (transactions that interact with smart contracts)
//   - For regular ETH transfers, this endpoint may not provide meaningful results
func (c *HTTPClient) GetContractExecutionStatus(ctx context.Context, txhash string, opts *GetContractExecutionStatusOpts) (*RespContractExecutionStatus, error) {
	params := map[string]string{
		"txhash": txhash,
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
		module:          "transaction",
		action:          "getstatus",
		params:          params,
		noFoundReturn:   RespContractExecutionStatus{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespContractExecutionStatus
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTransactionReceiptStatusOpts contains optional parameters for GetTransactionReceiptStatus
type GetTransactionReceiptStatusOpts struct {
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

// GetTransactionReceiptStatus returns the status code of a transaction execution
//
// This endpoint returns the receipt status of a transaction, indicating whether
// the transaction was successful or failed. This is different from contract execution
// status as it applies to all types of transactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txhash: The transaction hash to check the receipt status for
//     Must be a valid transaction hash (0x... format)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - *RespCheckTxReceiptStatus: Transaction receipt status information
//   - error: Error if the request fails
//
// Example:
//
//	// Check transaction receipt status
//	status, err := client.GetTransactionReceiptStatus(ctx, "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if status.Status == "1" {
//	    fmt.Println("Transaction successful")
//	} else {
//	    fmt.Println("Transaction failed")
//	}
//
//	// With custom chain
//	chainID := etherscan.PolygonMainnet
//	status, err := client.GetTransactionReceiptStatus(ctx, txhash, &GetTransactionReceiptStatusOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Status field: "1" = success, "0" = failed
//   - Only applicable for post Byzantium Fork transactions (block 4,370,000 on Ethereum)
//   - This endpoint works for all transaction types (ETH transfers, contract calls, etc.)
//   - Pre-Byzantium transactions will not have receipt status information
func (c *HTTPClient) GetTransactionReceiptStatus(ctx context.Context, txhash string, opts *GetTransactionReceiptStatusOpts) (*RespCheckTxReceiptStatus, error) {
	params := map[string]string{
		"txhash": txhash,
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
		module:          "transaction",
		action:          "gettxreceiptstatus",
		params:          params,
		noFoundReturn:   RespCheckTxReceiptStatus{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespCheckTxReceiptStatus
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ============================================================================
// Account Module - Internal Transactions
// ============================================================================

// GetInternalTransactionsByAddressOpts contains optional parameters for GetInternalTransactionsByAddress
type GetInternalTransactionsByAddressOpts struct {
	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	StartBlock *int

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	EndBlock *int

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset *int

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort *string

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

// GetInternalTransactionsByAddress returns list of 'Internal' transactions by address
//
// This endpoint returns a list of internal transactions (also called trace transactions)
// that involve the specified address. Internal transactions are transactions that occur
// within smart contract executions, such as contract-to-contract calls, contract creation,
// and self-destruct operations.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to get internal transactions for (must be a valid Ethereum address)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespInternalTxByAddress: List of internal transactions with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get internal transactions for an address
//	txs, err := client.GetInternalTransactionsByAddress(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("Hash: %s, From: %s, To: %s, Value: %s, Type: %s\n",
//	        tx.Hash, tx.From, tx.To, tx.Value, tx.Type)
//	}
//
//	// Get internal transactions with pagination
//	page := 1
//	offset := 50
//	sort := "desc"
//	txs, err := client.GetInternalTransactionsByAddress(ctx, address, &GetInternalTransactionsByAddressOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	    Sort:   &sort,
//	})
//
//	// Get internal transactions from a specific block range
//	startBlock := 18000000
//	endBlock := 18000100
//	txs, err := client.GetInternalTransactionsByAddress(ctx, address, &GetInternalTransactionsByAddressOpts{
//	    StartBlock: &startBlock,
//	    EndBlock:   &endBlock,
//	})
//
// Note:
//   - This endpoint returns max 10000 records per call
//   - Use StartBlock and EndBlock parameters to paginate through larger ranges
//   - Internal transactions include contract calls, contract creation, and self-destruct operations
//   - Different from normal transactions - these are trace-level transactions within contract execution
//   - All values are returned as strings in Wei
//   - Type field indicates the type of internal transaction (call, create, etc.)
func (c *HTTPClient) GetInternalTransactionsByAddress(ctx context.Context, address string, opts *GetInternalTransactionsByAddressOpts) ([]RespInternalTxByAddress, error) {
	params := map[string]string{
		"address":    address,
		"startblock": "0",
		"endblock":   "999999999999",
		"page":       "1",
		"offset":     "100",
		"sort":       "asc",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.StartBlock != nil {
			params["startblock"] = strconv.Itoa(*opts.StartBlock)
		}
		if opts.EndBlock != nil {
			params["endblock"] = strconv.Itoa(*opts.EndBlock)
		}
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
		action:          "txlistinternal",
		params:          params,
		noFoundReturn:   []RespInternalTxByAddress{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespInternalTxByAddress
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInternalTransactionsByHashOpts contains optional parameters for GetInternalTransactionsByHash
type GetInternalTransactionsByHashOpts struct {
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

// GetInternalTransactionsByHash returns list of internal transactions by transaction hash
//
// This endpoint returns all internal transactions that occurred within a specific
// transaction execution. When a transaction calls a smart contract, that contract
// may in turn call other contracts, creating a chain of internal transactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txhash: The transaction hash to get internal transactions for
//     Must be a valid transaction hash (0x... format)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespInternalTxByHash: List of internal transactions within the specified transaction
//   - error: Error if the request fails
//
// Example:
//
//	// Get internal transactions for a specific transaction
//	txs, err := client.GetInternalTransactionsByHash(ctx, "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("From: %s, To: %s, Value: %s, Type: %s\n",
//	        tx.From, tx.To, tx.Value, tx.Type)
//	}
//
//	// With custom chain
//	chainID := etherscan.PolygonMainnet
//	txs, err := client.GetInternalTransactionsByHash(ctx, txhash, &GetInternalTransactionsByHashOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - This endpoint returns max 10000 records per call
//   - Only returns internal transactions if the specified transaction involved contract calls
//   - For simple ETH transfers, this will return an empty list
//   - Internal transactions show the complete execution trace of a transaction
//   - All values are returned as strings in Wei
func (c *HTTPClient) GetInternalTransactionsByHash(ctx context.Context, txhash string, opts *GetInternalTransactionsByHashOpts) ([]RespInternalTxByHash, error) {
	params := map[string]string{
		"txhash": txhash,
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
		action:          "txlistinternal",
		params:          params,
		noFoundReturn:   []RespInternalTxByHash{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespInternalTxByHash
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInternalTransactionsByBlockRangeOpts contains optional parameters for GetInternalTransactionsByBlockRange
type GetInternalTransactionsByBlockRangeOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page *int

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset *int

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort *string

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

// GetInternalTransactionsByBlockRange returns list of internal transactions within a block range
//
// This endpoint returns all internal transactions that occurred within a specified
// range of blocks. This is useful for analyzing contract activity across multiple blocks
// or for getting a comprehensive view of internal transactions in a time period.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startblock: The starting block number to search from (inclusive)
//   - endblock: The ending block number to search to (inclusive)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespInternalTxByBlockRange: List of internal transactions within the block range
//   - error: Error if the request fails
//
// Example:
//
//	// Get internal transactions in a block range
//	txs, err := client.GetInternalTransactionsByBlockRange(ctx, 18000000, 18000100, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("Block: %s, Hash: %s, From: %s, To: %s, Value: %s\n",
//	        tx.BlockNumber, tx.Hash, tx.From, tx.To, tx.Value)
//	}
//
//	// With pagination
//	page := 1
//	offset := 50
//	sort := "desc"
//	txs, err := client.GetInternalTransactionsByBlockRange(ctx, startblock, endblock, &GetInternalTransactionsByBlockRangeOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	    Sort:   &sort,
//	})
//
// Note:
//   - This endpoint returns max 10000 records per call
//   - Use Page and Offset parameters to paginate through results
//   - Block range should be reasonable to avoid timeout (recommended: < 1000 blocks)
//   - Internal transactions include contract calls, contract creation, and self-destruct operations
//   - All values are returned as strings in Wei
//   - TraceID field helps identify which internal transactions belong to the same execution trace
func (c *HTTPClient) GetInternalTransactionsByBlockRange(ctx context.Context, startblock, endblock int, opts *GetInternalTransactionsByBlockRangeOpts) ([]RespInternalTxByBlockRange, error) {
	params := map[string]string{
		"startblock": strconv.Itoa(startblock),
		"endblock":   strconv.Itoa(endblock),
		"page":       "1",
		"offset":     "100",
		"sort":       "asc",
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
		action:          "txlistinternal",
		params:          params,
		noFoundReturn:   []RespInternalTxByBlockRange{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespInternalTxByBlockRange
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
