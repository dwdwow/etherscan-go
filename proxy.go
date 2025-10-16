package etherscan

import (
	"context"
)

// ============================================================================
// Geth/Parity Proxy Module
// ============================================================================

// RpcEthBlockNumberOpts contains optional parameters for RpcEthBlockNumber
type RpcEthBlockNumberOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthBlockNumber returns the number of the most recent block
//
// This endpoint returns the number of the most recent block in hex format.
// This is equivalent to the eth_blockNumber JSON-RPC method.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The most recent block number in hex format (e.g., "0x1b4")
//   - error: Error if the request fails
//
// Example:
//
//	// Get the latest block number
//	blockNumber, err := client.RpcEthBlockNumber(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Latest block: %s\n", blockNumber)
//
//	// Convert hex to decimal
//	blockNum, _ := strconv.ParseInt(blockNumber[2:], 16, 64)
//	fmt.Printf("Block number: %d\n", blockNum)
//
// Note:
//   - Returns block number in hex format with "0x" prefix
//   - Equivalent to eth_blockNumber JSON-RPC method
func (c *HTTPClient) RpcEthBlockNumber(ctx context.Context, opts *RpcEthBlockNumberOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_blockNumber",
		params:          params,
		noFoundReturn:   RespEthBlockNumberHex{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthBlockNumberHex
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthBlockByNumberOpts contains optional parameters for RpcEthBlockByNumber
type RpcEthBlockByNumberOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthBlockByNumber returns information about a block by block number
//
// This endpoint returns information about a block by block number. This is equivalent
// to the eth_getBlockByNumber JSON-RPC method. You can choose to get full transaction
// objects or just transaction hashes.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - tag: Block number in hex format (e.g., "0xC36B3C") or "latest", "earliest", "pending"
//   - boolean: If true, returns full transaction objects. If false, returns only transaction hashes
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthBlockInfo: Block information including transactions
//   - error: Error if the request fails
//
// Example:
//
//	// Get latest block with full transaction objects
//	block, err := client.RpcEthBlockByNumber(ctx, "latest", true, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Block Number: %s\n", block.Number)
//	fmt.Printf("Block Hash: %s\n", block.Hash)
//	fmt.Printf("Transaction Count: %d\n", len(block.Transactions))
//
//	// Get specific block with transaction hashes only
//	block, err := client.RpcEthBlockByNumber(ctx, "0x1B4", false, nil)
//
// Note:
//   - Equivalent to eth_getBlockByNumber JSON-RPC method
//   - Tag can be block number in hex or "latest", "earliest", "pending"
//   - Boolean parameter controls transaction detail level
func (c *HTTPClient) RpcEthBlockByNumber(ctx context.Context, tag string, opts *RpcEthBlockByNumberOpts) (*RespEthBlockInfo, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["tag"] = tag
	params["boolean"] = "false"

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getBlockByNumber",
		params:          params,
		noFoundReturn:   RespEthBlock{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthBlock
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthBlockByNumber returns information about a block by block number
//
// This endpoint returns information about a block by block number.
// The block can be specified by number in hex or by tag ("latest", "earliest", "pending").
// This is equivalent to the eth_getBlockByNumber JSON-RPC method.
//
// Example usage:
//
//	// Get latest block with transaction hashes only
//	block, err := client.RpcEthBlockByNumber(ctx, "latest", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Block Number: %s\n", block.Number)
//	fmt.Printf("Block Hash: %s\n", block.Hash)
//	fmt.Printf("Transaction Count: %d\n", len(block.Transactions))
//
//	// Get specific block from Ethereum mainnet
//	block, err := client.RpcEthBlockByNumber(ctx, "0x1B4", &RpcEthBlockByNumberOpts{
//	    ChainID: &[]int64{1}[0],
//	})
//
// Note:
//   - Returns block information with transaction hashes (not full transaction objects)
//   - For full transaction objects, use RpcEthBlockByNumberWithFullTxs instead
//   - Block tag can be block number in hex or "latest", "earliest", "pending"

func (c *HTTPClient) RpcEthBlockByNumberWithFullTxs(ctx context.Context, tag string, opts *RpcEthBlockByNumberOpts) (*RespEthBlockInfoWithFullTxs, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["tag"] = tag
	params["boolean"] = "true"

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getBlockByNumber",
		params:          params,
		noFoundReturn:   RespEthBlockWithFullTxs{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthBlockWithFullTxs
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthUncleByBlockNumberAndIndexOpts contains optional parameters for RpcEthUncleByBlockNumberAndIndex
type RpcEthUncleByBlockNumberAndIndexOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthUncleByBlockNumberAndIndex returns information about an uncle block by block number and index
//
// This endpoint returns information about an uncle block by block number and index.
// Uncle blocks are blocks that were mined but not included in the main chain.
// This is equivalent to the eth_getUncleByBlockNumberAndIndex JSON-RPC method.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - tag: Block number in hex format (e.g., "0xC36B3C") or "latest", "earliest", "pending"
//   - index: Position of the uncle's index in the block, in hex (e.g., "0x0")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthUncleBlockInfo: Uncle block information
//   - error: Error if the request fails
//
// Example:
//
//	// Get uncle block information
//	blockTag := "0x1B4"
//	uncleIndex := "0x0"
//	uncle, err := client.RpcEthUncleByBlockNumberAndIndex(ctx, blockTag, uncleIndex, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Uncle Block Number: %s\n", uncle.Number)
//	fmt.Printf("Uncle Block Hash: %s\n", uncle.Hash)
//
// Note:
//   - Equivalent to eth_getUncleByBlockNumberAndIndex JSON-RPC method
//   - Uncle blocks are blocks that were mined but not included in the main chain
func (c *HTTPClient) RpcEthUncleByBlockNumberAndIndex(ctx context.Context, tag, index string, opts *RpcEthUncleByBlockNumberAndIndexOpts) (*RespEthUncleBlockInfo, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["tag"] = tag
	params["index"] = index

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getUncleByBlockNumberAndIndex",
		params:          params,
		noFoundReturn:   RespEthUncleBlock{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthUncleBlock
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthBlockTxCountByNumberOpts contains optional parameters for RpcEthBlockTxCountByNumber
type RpcEthBlockTxCountByNumberOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthBlockTransactionCountByNumber returns the number of transactions in a block by block number
//
// This endpoint returns the number of transactions in a block by block number in hex format.
// This is equivalent to the eth_getBlockTransactionCountByNumber JSON-RPC method.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - tag: Block number in hex format (e.g., "0x10FB78") or "latest", "earliest", "pending"
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Number of transactions in the block in hex format (e.g., "0x1A")
//   - error: Error if the request fails
//
// Example:
//
//	// Get transaction count for a specific block
//	blockTag := "0x10FB78"
//	count, err := client.RpcEthBlockTransactionCountByNumber(ctx, blockTag, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction count: %s\n", count)
//
//	// Convert hex to decimal
//	txCount, _ := strconv.ParseInt(count[2:], 16, 64)
//	fmt.Printf("Transaction count: %d\n", txCount)
//
// Note:
//   - Equivalent to eth_getBlockTransactionCountByNumber JSON-RPC method
//   - Returns count in hex format with "0x" prefix
func (c *HTTPClient) RpcEthBlockTxCountByNumber(ctx context.Context, tag string, opts *RpcEthBlockTxCountByNumberOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["tag"] = tag

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getBlockTransactionCountByNumber",
		params:          params,
		noFoundReturn:   RespEthBlockTxCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthBlockTxCount
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthTxByHashOpts contains optional parameters for RpcEthTxByHash
type RpcEthTxByHashOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthTransactionByHash returns information about a transaction by its hash
//
// This endpoint returns information about a transaction by its hash. This is equivalent
// to the eth_getTransactionByHash JSON-RPC method. Returns detailed transaction information
// including sender, receiver, value, gas, and input data.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txHash: Transaction hash (e.g., "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthTxInfo: Transaction information including hash, from, to, value, gas, etc.
//   - error: Error if the request fails
//
// Example:
//
//	// Get transaction information by hash
//	txHash := "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44"
//	tx, err := client.RpcEthTransactionByHash(ctx, txHash, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction Hash: %s\n", tx.Hash)
//	fmt.Printf("From: %s\n", tx.From)
//	fmt.Printf("To: %s\n", tx.To)
//	fmt.Printf("Value: %s\n", tx.Value)
//	fmt.Printf("Gas: %s\n", tx.Gas)
//	fmt.Printf("Gas Price: %s\n", tx.GasPrice)
//
// Note:
//   - Equivalent to eth_getTransactionByHash JSON-RPC method
//   - Returns nil if transaction not found
//   - All values are in hex format
func (c *HTTPClient) RpcEthTxByHash(ctx context.Context, txHash string, opts *RpcEthTxByHashOpts) (*RespEthTxInfo, error) {
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
		module:          "proxy",
		action:          "eth_getTransactionByHash",
		params:          params,
		noFoundReturn:   RespEthTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthTxByBlockNumberAndIndexOpts contains optional parameters for RpcEthTxByBlockNumberAndIndex
type RpcEthTxByBlockNumberAndIndexOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position
//
// This endpoint returns information about a transaction by block number and transaction
// index position. This is equivalent to the eth_getTransactionByBlockNumberAndIndex
// JSON-RPC method. Returns detailed transaction information.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - tag: Block number in hex format (e.g., "0x10FB78") or "latest", "earliest", "pending"
//   - index: Transaction index position in hex (e.g., "0x0")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthTxInfo: Transaction information including hash, from, to, value, gas, etc.
//   - error: Error if the request fails
//
// Example:
//
//	// Get transaction by block number and index
//	blockTag := "0x10FB78"
//	txIndex := "0x0"
//	tx, err := client.RpcEthTransactionByBlockNumberAndIndex(ctx, blockTag, txIndex, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction Hash: %s\n", tx.Hash)
//	fmt.Printf("From: %s\n", tx.From)
//	fmt.Printf("To: %s\n", tx.To)
//	fmt.Printf("Value: %s\n", tx.Value)
//
// Note:
//   - Equivalent to eth_getTransactionByBlockNumberAndIndex JSON-RPC method
//   - Returns nil if transaction not found
//   - Index must be within the block's transaction count
func (c *HTTPClient) RpcEthTxByBlockNumberAndIndex(ctx context.Context, tag, index string, opts *RpcEthTxByBlockNumberAndIndexOpts) (*RespEthTxInfo, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return nil, err
	}

	// Add required parameters
	params["tag"] = tag
	params["index"] = index

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getTransactionByBlockNumberAndIndex",
		params:          params,
		noFoundReturn:   RespEthTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthTx
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthTxCountOpts contains optional parameters for RpcEthTxCount
type RpcEthTxCountOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthTransactionCount returns the number of transactions performed by an address
//
// This endpoint returns the number of transactions performed by an address (nonce) in hex format.
// This is equivalent to the eth_getTransactionCount JSON-RPC method. The nonce represents
// the number of transactions sent from this address.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get transaction count for
//   - tag: Block parameter - "latest", "earliest" or "pending"
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Number of transactions in hex format (e.g., "0x1A")
//   - error: Error if the request fails
//
// Example:
//
//	// Get transaction count for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	count, err := client.RpcEthTransactionCount(ctx, addr, "latest", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction count: %s\n", count)
//
//	// Convert hex to decimal
//	nonce, _ := strconv.ParseInt(count[2:], 16, 64)
//	fmt.Printf("Nonce: %d\n", nonce)
//
// Note:
//   - Equivalent to eth_getTransactionCount JSON-RPC method
//   - Returns nonce in hex format with "0x" prefix
//   - Nonce represents the number of transactions sent from this address
func (c *HTTPClient) RpcEthTxCount(ctx context.Context, address, tag string, opts *RpcEthTxCountOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["address"] = address
	params["tag"] = tag

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getTransactionCount",
		params:          params,
		noFoundReturn:   RespEthTxCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthTxCount
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthSendRawTxOpts contains optional parameters for RpcEthSendRawTx
type RpcEthSendRawTxOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthSendRawTransaction submits a pre-signed transaction for broadcast to the Ethereum network
//
// This endpoint submits a pre-signed transaction for broadcast to the Ethereum network.
// This is equivalent to the eth_sendRawTransaction JSON-RPC method. The transaction
// must be properly signed and serialized before submission.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - hex: Signed raw transaction data to broadcast
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Transaction hash of the broadcasted transaction
//   - error: Error if the request fails
//
// Example:
//
//	// Submit a raw transaction
//	rawTx := "0xf86c808502540be400825208943535353535353535353535353535353535353535880de0b6b3a76400008025a028ef61340bd939bc2195fe537567866003e1a15d3c71ff63e1590620aa636276a067cbe9d8997f761aecb703304b3800ccf555c9f3dc64214b297fb1966a3b6d83"
//	txHash, err := client.RpcEthSendRawTransaction(ctx, rawTx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction Hash: %s\n", txHash)
//
// Note:
//   - Equivalent to eth_sendRawTransaction JSON-RPC method
//   - Transaction must be properly signed and serialized
//   - For long hex strings, request is automatically sent as POST
//   - Returns transaction hash if successful
func (c *HTTPClient) RpcEthSendRawTx(ctx context.Context, hex string, opts *RpcEthSendRawTxOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["hex"] = hex

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_sendRawTransaction",
		params:          params,
		method:          "POST",
		noFoundReturn:   RespEthSendRawTx{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthSendRawTx
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthTxReceiptOpts contains optional parameters for RpcEthTxReceipt
type RpcEthTxReceiptOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthTransactionReceipt returns the receipt of a transaction by transaction hash
//
// This endpoint returns the receipt of a transaction by transaction hash. This is equivalent
// to the eth_getTransactionReceipt JSON-RPC method. Returns detailed transaction receipt
// information including gas used, logs, and execution status.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - txHash: Hash of the transaction
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespEthTxReceiptInfo: Transaction receipt information including gas used, logs, status
//   - error: Error if the request fails
//
// Example:
//
//	// Get transaction receipt
//	txHash := "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44"
//	receipt, err := client.RpcEthTransactionReceipt(ctx, txHash, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transaction Hash: %s\n", receipt.TransactionHash)
//	fmt.Printf("Status: %s\n", receipt.Status)
//	fmt.Printf("Gas Used: %s\n", receipt.GasUsed)
//	fmt.Printf("Cumulative Gas Used: %s\n", receipt.CumulativeGasUsed)
//	fmt.Printf("Contract Address: %s\n", receipt.ContractAddress)
//
// Note:
//   - Equivalent to eth_getTransactionReceipt JSON-RPC method
//   - Returns nil if transaction not found
//   - Status is "0x1" for success, "0x0" for failure
//   - Only applicable for post Byzantium Fork transactions
func (c *HTTPClient) RpcEthTxReceipt(ctx context.Context, txHash string, opts *RpcEthTxReceiptOpts) (*RespEthTxReceiptInfo, error) {
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
		module:          "proxy",
		action:          "eth_getTransactionReceipt",
		params:          params,
		noFoundReturn:   RespEthTxReceipt{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthTxReceipt
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result.Result, nil
}

// RpcEthCallOpts contains optional parameters for RpcEthCall
type RpcEthCallOpts struct {
	// Tag specifies the block parameter for the call
	// Options: "latest", "earliest", "pending", or block number in hex
	// Default: "latest"
	Tag string `default:"latest" json:"tag"`

	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthCall executes a new message call immediately without creating a transaction on the block chain
//
// This endpoint executes a new message call immediately without creating a transaction
// on the blockchain. This is equivalent to the eth_call JSON-RPC method. Useful for
// calling view/pure functions on smart contracts.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - to: Address to interact with
//   - data: Hash of the method signature and encoded parameters
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The return value of the executed contract call in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Call a contract function
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	// Method signature: balanceOf(address)
//	methodSig := "0x70a08231"
//	// Encoded parameter: address padded to 32 bytes
//	param := "000000000000000000000000742d35cc6634c0532925a3b844bc9e7595f0beb"
//	callData := methodSig + param
//
//	result, err := client.RpcEthCall(ctx, contractAddr, callData, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Call result: %s\n", result)
//
//	// With custom block tag
//	tag := "0x1B4"
//	result, err := client.RpcEthCall(ctx, contractAddr, callData, &RpcEthCallOpts{
//	    Tag: &tag,
//	})
//
// Note:
//   - Equivalent to eth_call JSON-RPC method
//   - Gas parameter is capped at 2x the current block gas limit
//   - Returns result in hex format
//   - Useful for calling view/pure functions on smart contracts
func (c *HTTPClient) RpcEthCall(ctx context.Context, to, data string, opts *RpcEthCallOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["to"] = to
	params["data"] = data

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	result, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_call",
		params:          params,
		noFoundReturn:   RespEthCall{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var resp RespEthCall
	if err := unmarshalResponse(result, &resp); err != nil {
		return "", err
	}
	return resp.Result, nil
}

// RpcEthGetCodeOpts contains optional parameters for RpcEthGetCode
type RpcEthGetCodeOpts struct {
	// Tag specifies the block parameter for the call
	// Options: "latest", "earliest", "pending", or block number in hex
	// Default: "latest"
	Tag string `default:"latest" json:"tag"`

	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthGetCode returns code at a given address
//
// This endpoint returns code at a given address. This is equivalent to the eth_getCode
// JSON-RPC method. Returns the bytecode of a contract at the specified address.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get code from
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The code from the given address in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get contract bytecode
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	code, err := client.RpcEthGetCode(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Contract bytecode: %s\n", code)
//
//	// With custom block tag
//	tag := "0x1B4"
//	code, err := client.RpcEthGetCode(ctx, contractAddr, &RpcEthGetCodeOpts{
//	    Tag: &tag,
//	})
//
// Note:
//   - Equivalent to eth_getCode JSON-RPC method
//   - Returns "0x" if address has no code (EOA)
//   - Returns contract bytecode if address is a contract
func (c *HTTPClient) RpcEthGetCode(ctx context.Context, address string, opts *RpcEthGetCodeOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
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
		module:          "proxy",
		action:          "eth_getCode",
		params:          params,
		noFoundReturn:   RespEthGetCode{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthGetCode
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthGetStorageAtOpts contains optional parameters for RpcEthGetStorageAt
type RpcEthGetStorageAtOpts struct {
	// Tag specifies the block parameter for the call
	// Options: "latest", "earliest", "pending", or block number in hex
	// Default: "latest"
	Tag string `default:"latest" json:"tag"`

	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthGetStorageAt returns the value from a storage position at a given address
//
// This endpoint returns the value from a storage position at a given address. This is
// equivalent to the eth_getStorageAt JSON-RPC method. Useful for reading contract
// storage variables.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get storage from
//   - position: Position in storage (hex string)
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The value at the given storage position in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Get storage value at position 0
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	position := "0x0"
//	value, err := client.RpcEthGetStorageAt(ctx, contractAddr, position, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Storage value: %s\n", value)
//
// Note:
//   - Equivalent to eth_getStorageAt JSON-RPC method
//   - This endpoint is experimental and may have potential issues
//   - Position must be in hex format
func (c *HTTPClient) RpcEthGetStorageAt(ctx context.Context, address, position string, opts *RpcEthGetStorageAtOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["address"] = address
	params["position"] = position

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_getStorageAt",
		params:          params,
		noFoundReturn:   RespEthGetStorageAt{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthGetStorageAt
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthGetGasPriceOpts contains optional parameters for RpcEthGetGasPrice
type RpcEthGetGasPriceOpts struct {
	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthGetGasPrice returns the current price per gas in wei
//
// This endpoint returns the current price per gas in wei. This is equivalent to the
// eth_gasPrice JSON-RPC method. Useful for determining current gas prices for
// transaction estimation.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The current gas price in wei (hex string)
//   - error: Error if the request fails
//
// Example:
//
//	// Get current gas price
//	gasPrice, err := client.RpcEthGetGasPrice(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Gas price: %s wei\n", gasPrice)
//
//	// Convert to Gwei
//	priceWei, _ := new(big.Int).SetString(gasPrice[2:], 16)
//	priceGwei := new(big.Float).Quo(new(big.Float).SetInt(priceWei), big.NewFloat(1e9))
//	fmt.Printf("Gas price: %s Gwei\n", priceGwei.String())
//
// Note:
//   - Equivalent to eth_gasPrice JSON-RPC method
//   - Returns gas price in wei as hex string
//   - Useful for transaction gas price estimation
func (c *HTTPClient) RpcEthGetGasPrice(ctx context.Context, opts *RpcEthGetGasPriceOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_gasPrice",
		params:          params,
		noFoundReturn:   RespEthGetGasPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var result RespEthGetGasPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}

// RpcEthEstimateGasOpts contains optional parameters for RpcEthEstimateGas
type RpcEthEstimateGasOpts struct {
	// Value is the value sent in transaction (hex string)
	// Default: "" (no value sent)
	Value string `default:"" json:"value"`

	// Gas is the amount of gas provided for transaction (hex string)
	// Default: "" (uses default gas limit)
	Gas string `default:"" json:"gas"`

	// GasPrice is the gas price in wei (hex string)
	// Default: "" (uses current network gas price)
	GasPrice string `default:"" json:"gasprice"`

	// ChainID specifies which blockchain network to query
	// Default: empty (uses client default)
	ChainID int64 `json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// RpcEthEstimateGas makes a call or transaction, which won't be added to the blockchain and returns the used gas
//
// This endpoint makes a call or transaction, which won't be added to the blockchain and
// returns the used gas. This is equivalent to the eth_estimateGas JSON-RPC method.
// Useful for estimating gas costs before sending transactions.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - to: Address to interact with
//   - data: Hash of the method signature and encoded parameters
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The estimated gas usage in hex format
//   - error: Error if the request fails
//
// Example:
//
//	// Estimate gas for a contract call
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	callData := "0x70a08231000000000000000000000000742d35cc6634c0532925a3b844bc9e7595f0beb"
//	gasEstimate, err := client.RpcEthEstimateGas(ctx, contractAddr, callData, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Estimated gas: %s\n", gasEstimate)
//
//	// With custom parameters
//	value := "0x0"
//	gas := "0x5208"
//	gasPrice := "0x4a817c800"
//	gasEstimate, err := client.RpcEthEstimateGas(ctx, contractAddr, callData, &RpcEthEstimateGasOpts{
//	    Value:    &value,
//	    Gas:      &gas,
//	    GasPrice: &gasPrice,
//	})
//
// Note:
//   - Equivalent to eth_estimateGas JSON-RPC method
//   - Gas parameter is capped at 2x the current block gas limit
//   - Post EIP-1559: gasPrice must be higher than block's baseFeePerGas
//   - Returns estimated gas usage in hex format
func (c *HTTPClient) RpcEthEstimateGas(ctx context.Context, to, data string, opts *RpcEthEstimateGasOpts) (string, error) {
	// Apply defaults and extract API parameters
	params, err := ApplyDefaultsAndExtractParams(opts)
	if err != nil {
		return "", err
	}

	// Add required parameters
	params["to"] = to
	params["data"] = data

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	result, err := c.request(requestParams{
		ctx:             ctx,
		module:          "proxy",
		action:          "eth_estimateGas",
		params:          params,
		noFoundReturn:   RespEthEstimateGas{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	var resp RespEthEstimateGas
	if err := unmarshalResponse(result, &resp); err != nil {
		return "", err
	}
	return resp.Result, nil
}
