package etherscan

import (
	"context"
	"strconv"
)

// GetNormalTxsOpts contains optional parameters for GetNormalTxs
type GetNormalTxsOpts struct {
	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"100" json:"offset"`

	// Sort order for the results
	// Default: "asc"
	// Options:
	//   - "asc": Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetNormalTxs returns list of 'Normal' transactions by address
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
//	txs, err := client.GetNormalTxs(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
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
//	txs, err := client.GetNormalTxs(ctx, address, &GetNormalTxsOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	    Sort:   &sort,
//	})
//
//	// Get transactions from a specific block range
//	startBlock := 18000000
//	endBlock := 18000100
//	txs, err := client.GetNormalTxs(ctx, address, &GetNormalTxsOpts{
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
func (c *HTTPClient) GetNormalTxs(ctx context.Context, address string, opts *GetNormalTxsOpts) ([]RespNormalTx, error) {
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

// GetBridgeTxsOpts contains optional parameters for GetBridgeTxs
type GetBridgeTxsOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"100" json:"offset"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Note: This endpoint is only applicable to specific chains
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetBridgeTxs returns bridge transactions for an address
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
//	txs, err := client.GetBridgeTxs(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
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
//	txs, err := client.GetBridgeTxs(ctx, address, &GetBridgeTxsOpts{
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
func (c *HTTPClient) GetBridgeTxs(ctx context.Context, address string, opts *GetBridgeTxsOpts) ([]RespBridgeTx, error) {
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
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetContractExecutionStatus returns the status code of a contract execution
//
// This endpoint returns the execution status of a contract transaction.
// It provides information about whether the contract execution was successful
// or failed, along with error details if applicable.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txHash: The transaction hash to check the execution status for
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
//	status, err := client.GetContractExecutionStatus(ctx, txHash, &GetContractExecutionStatusOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - IsError field: "0" = success, "1" = failed
//   - ErrDescription field: Contains error message if execution failed
//   - Only applicable to contract transactions (transactions that interact with smart contracts)
//   - For regular ETH transfers, this endpoint may not provide meaningful results
func (c *HTTPClient) GetContractExecutionStatus(ctx context.Context, txHash string, opts *GetContractExecutionStatusOpts) (*RespContractExecutionStatus, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["txhash"] = txHash

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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

// GetTxReceiptStatusOpts contains optional parameters for GetTxReceiptStatus
type GetTxReceiptStatusOpts struct {
	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetTxReceiptStatus returns the status code of a transaction execution
//
// This endpoint returns the receipt status of a transaction, indicating whether
// the transaction was successful or failed. This is different from contract execution
// status as it applies to all types of transactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txHash: The transaction hash to check the receipt status for
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
//	status, err := client.GetTxReceiptStatus(ctx, "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44", nil)
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
//	status, err := client.GetTxReceiptStatus(ctx, txHash, &GetTxReceiptStatusOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Status field: "1" = success, "0" = failed
//   - Only applicable for post Byzantium Fork transactions (block 4,370,000 on Ethereum)
//   - This endpoint works for all transaction types (ETH transfers, contract calls, etc.)
//   - Pre-Byzantium transactions will not have receipt status information
func (c *HTTPClient) GetTxReceiptStatus(ctx context.Context, txHash string, opts *GetTxReceiptStatusOpts) (*RespCheckTxReceiptStatus, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["txhash"] = txHash

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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

// GetInternalTxsByAddressOpts contains optional parameters for GetInternalTxsByAddress
type GetInternalTxsByAddressOpts struct {
	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"100" json:"offset"`

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetInternalTxsByAddress returns list of 'Internal' transactions by address
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
//	txs, err := client.GetInternalTxsByAddress(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
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
//	txs, err := client.GetInternalTxsByAddress(ctx, address, &GetInternalTxsByAddressOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	    Sort:   &sort,
//	})
//
//	// Get internal transactions from a specific block range
//	startBlock := 18000000
//	endBlock := 18000100
//	txs, err := client.GetInternalTxsByAddress(ctx, address, &GetInternalTxsByAddressOpts{
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
func (c *HTTPClient) GetInternalTxsByAddress(ctx context.Context, address string, opts *GetInternalTxsByAddressOpts) ([]RespInternalTxByAddress, error) {
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

// GetInternalTxsByHashOpts contains optional parameters for GetInternalTxsByHash
type GetInternalTxsByHashOpts struct {
	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetInternalTxsByHash returns list of internal transactions by transaction hash
//
// This endpoint returns all internal transactions that occurred within a specific
// transaction execution. When a transaction calls a smart contract, that contract
// may in turn call other contracts, creating a chain of internal transactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txHash: The transaction hash to get internal transactions for
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
//	txs, err := client.GetInternalTxsByHash(ctx, "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44", nil)
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
//	txs, err := client.GetInternalTxsByHash(ctx, txHash, &GetInternalTxsByHashOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - This endpoint returns max 10000 records per call
//   - Only returns internal transactions if the specified transaction involved contract calls
//   - For simple ETH transfers, this will return an empty list
//   - Internal transactions show the complete execution trace of a transaction
//   - All values are returned as strings in Wei
func (c *HTTPClient) GetInternalTxsByHash(ctx context.Context, txHash string, opts *GetInternalTxsByHashOpts) ([]RespInternalTxByHash, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["txhash"] = txHash

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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

// GetInternalTxsByBlockRangeOpts contains optional parameters for GetInternalTxsByBlockRange
type GetInternalTxsByBlockRangeOpts struct {
	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transactions per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"100" json:"offset"`

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// If 0, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If empty, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetInternalTxsByBlockRange returns list of internal transactions within a block range
//
// This endpoint returns all internal transactions that occurred within a specified
// range of blocks. This is useful for analyzing contract activity across multiple blocks
// or for getting a comprehensive view of internal transactions in a time period.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startBlock: The starting block number to search from (inclusive)
//   - endBlock: The ending block number to search to (inclusive)
//   - opts: Optional parameters (can be nil for defaults)
//
// Returns:
//   - []RespInternalTxByBlockRange: List of internal transactions within the block range
//   - error: Error if the request fails
//
// Example:
//
//	// Get internal transactions in a block range
//	txs, err := client.GetInternalTxsByBlockRange(ctx, 18000000, 18000100, nil)
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
//	txs, err := client.GetInternalTxsByBlockRange(ctx, startBlock, endBlock, &GetInternalTxsByBlockRangeOpts{
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
func (c *HTTPClient) GetInternalTxsByBlockRange(ctx context.Context, startBlock, endBlock int, opts *GetInternalTxsByBlockRangeOpts) ([]RespInternalTxByBlockRange, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["startblock"] = strconv.Itoa(startBlock)
	params["endblock"] = strconv.Itoa(endBlock)

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
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
