package etherscan

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// ============================================================================
// Account Module - ERC-20/721/1155 Token Transfers
// ============================================================================

// GetERC20TokenTransfersOpts contains optional parameters for GetERC20TokenTransfers
type GetERC20TokenTransfersOpts struct {
	// Address is the address to get token transfers for (optional)
	// If provided, returns all ERC-20 token transfers involving this address
	// Must be a valid Ethereum address format
	Address string `json:"address"`

	// ContractAddress is the token contract address to filter by (optional)
	// If provided, returns transfers only for this specific token
	// Must be a valid contract address format
	ContractAddress string `json:"contractaddress"`

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transfers per page
	// Default: 100
	// Maximum: 10000
	// Higher values return more results per page but may be slower
	Offset int64 `default:"100" json:"offset"`

	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	// Use this to limit the search range and improve performance
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	// Use this to limit the search range and improve performance
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by block number in ascending order (oldest first)
	//   - "desc": Sort by block number in descending order (newest first)
	Sort string `default:"asc" json:"sort"`

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

// GetERC20TokenTransfers returns list of ERC-20 token transfers by address and/or contract address
//
// This endpoint returns a list of ERC-20 token transfers. You can filter by:
// - A specific address (all token transfers involving that address)
// - A specific token contract (all transfers of that token)
// - Both (transfers of a specific token involving a specific address)
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Parameters (at least one of Address or ContractAddress must be provided)
//
// Returns:
//   - []RespERC20TokenTransfer: List of ERC-20 token transfers with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get all token transfers for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	txs, err := client.GetERC20TokenTransfers(ctx, &GetERC20TokenTransfersOpts{
//	    Address: &addr,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tx := range txs {
//	    fmt.Printf("Token: %s, From: %s, To: %s, Value: %s\n",
//	        tx.TokenSymbol, tx.From, tx.To, tx.Value)
//	}
//
//	// Get transfers for a specific token
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	txs, err := client.GetERC20TokenTransfers(ctx, &GetERC20TokenTransfersOpts{
//	    ContractAddress: &contractAddr,
//	})
//
//	// Get transfers for a specific token and address
//	txs, err := client.GetERC20TokenTransfers(ctx, &GetERC20TokenTransfersOpts{
//	    Address:         &addr,
//	    ContractAddress: &contractAddr,
//	})
//
//	// With pagination and custom options
//	page := 1
//	offset := 50
//	sort := "desc"
//	startBlock := 18000000
//	endBlock := 18000100
//	txs, err := client.GetERC20TokenTransfers(ctx, &GetERC20TokenTransfersOpts{
//	    Address:    &addr,
//	    Page:       &page,
//	    Offset:     &offset,
//	    Sort:       &sort,
//	    StartBlock: &startBlock,
//	    EndBlock:   &endBlock,
//	})
//
// Note:
//   - At least one of Address or ContractAddress must be provided
//   - This endpoint returns max 10000 records per call
//   - Use Page and Offset parameters to paginate through results
//   - All values are returned as strings in the token's smallest unit
//   - TokenDecimal field indicates the number of decimal places for the token
func (c *HTTPClient) GetERC20TokenTransfers(ctx context.Context, opts *GetERC20TokenTransfersOpts) ([]RespERC20TokenTransfer, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	if opts.Address == "" && opts.ContractAddress == "" {
		return nil, fmt.Errorf("at least one of Address or ContractAddress must be provided")
	}

	params := map[string]string{}

	if opts.Address != "" {
		params["address"] = opts.Address
	}
	if opts.ContractAddress != "" {
		params["contractaddress"] = opts.ContractAddress
	}

	params["page"] = strconv.FormatInt(opts.Page, 10)
	params["offset"] = strconv.FormatInt(opts.Offset, 10)
	params["startblock"] = strconv.FormatInt(opts.StartBlock, 10)
	params["endblock"] = strconv.FormatInt(opts.EndBlock, 10)
	params["sort"] = opts.Sort

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "tokentx",
		params:          params,
		noFoundReturn:   []RespERC20TokenTransfer{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespERC20TokenTransfer
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetERC721TokenTransfersOpts contains optional parameters for GetERC721TokenTransfers
type GetERC721TokenTransfersOpts struct {
	// Address is the address to get NFT transfers for (optional)
	// If provided, returns all ERC-721 token transfers involving this address
	Address string `json:"address"`

	// ContractAddress is the NFT contract address to filter by (optional)
	// If provided, returns transfers only for this specific NFT collection
	ContractAddress string `json:"contractaddress"`

	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transfers per page
	// Default: 100
	// Maximum: 10000
	Offset int64 `default:"100" json:"offset"`

	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC721TokenTransfers returns list of ERC-721 (NFT) token transfers by address and/or contract address
//
// This endpoint returns a list of ERC-721 (NFT) token transfers. You can filter by:
// - A specific address (all NFT transfers involving that address)
// - A specific NFT contract (all transfers of that NFT collection)
// - Both (transfers of a specific NFT collection involving a specific address)
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Parameters (at least one of Address or ContractAddress must be provided)
//
// Returns:
//   - []RespERC721TokenTransfer: List of ERC-721 token transfers with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get all NFT transfers for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	txs, err := client.GetERC721TokenTransfers(ctx, &GetERC721TokenTransfersOpts{
//	    Address: &addr,
//	})
//
//	// Get transfers for a specific NFT collection
//	contractAddr := "0x60e4d786628fea6478f785a6d7e704777c86a7c6" // MAYC
//	txs, err := client.GetERC721TokenTransfers(ctx, &GetERC721TokenTransfersOpts{
//	    ContractAddress: &contractAddr,
//	})
//
// Note:
//   - At least one of Address or ContractAddress must be provided
//   - This endpoint returns max 10000 records per call
//   - Use Page and Offset parameters to paginate through results
//   - TokenID field contains the specific NFT token ID
func (c *HTTPClient) GetERC721TokenTransfers(ctx context.Context, opts *GetERC721TokenTransfersOpts) ([]RespERC721TokenTransfer, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	if opts.Address == "" && opts.ContractAddress == "" {
		return nil, fmt.Errorf("at least one of Address or ContractAddress must be provided")
	}

	params := map[string]string{}

	if opts.Address != "" {
		params["address"] = opts.Address
	}
	if opts.ContractAddress != "" {
		params["contractaddress"] = opts.ContractAddress
	}

	params["page"] = strconv.FormatInt(opts.Page, 10)
	params["offset"] = strconv.FormatInt(opts.Offset, 10)
	params["startblock"] = strconv.FormatInt(opts.StartBlock, 10)
	params["endblock"] = strconv.FormatInt(opts.EndBlock, 10)
	params["sort"] = opts.Sort

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "tokennfttx",
		params:          params,
		noFoundReturn:   []RespERC721TokenTransfer{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespERC721TokenTransfer
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetERC1155TokenTransfersOpts contains optional parameters for GetERC1155TokenTransfers
type GetERC1155TokenTransfersOpts struct {
	// Address is the address to get token transfers for (optional)
	// If provided, returns all ERC-1155 token transfers involving this address
	Address string `json:"address"`

	// ContractAddress is the token contract address to filter by (optional)
	// If provided, returns transfers only for this specific token
	ContractAddress string `json:"contractaddress"`

	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of transfers per page
	// Default: 100
	// Maximum: 10000
	Offset int64 `default:"100" json:"offset"`

	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Sort order for the results
	// Options: "asc" (default) or "desc"
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetERC1155TokenTransfers returns list of ERC-1155 (Multi Token Standard) token transfers by address and/or contract address
//
// This endpoint returns a list of ERC-1155 token transfers. ERC-1155 is a multi-token standard
// that allows for both fungible and non-fungible tokens in a single contract. You can filter by:
// - A specific address (all token transfers involving that address)
// - A specific token contract (all transfers of that token)
// - Both (transfers of a specific token involving a specific address)
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Parameters (at least one of Address or ContractAddress must be provided)
//
// Returns:
//   - []RespERC1155TokenTransfer: List of ERC-1155 token transfers with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get all ERC-1155 token transfers for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	txs, err := client.GetERC1155TokenTransfers(ctx, &GetERC1155TokenTransfersOpts{
//	    Address: &addr,
//	})
//
//	// Get transfers for a specific ERC-1155 contract
//	contractAddr := "0x76be3b62873462d2142405439777e971754e8e77" // OpenSea Shared Storefront
//	txs, err := client.GetERC1155TokenTransfers(ctx, &GetERC1155TokenTransfersOpts{
//	    ContractAddress: &contractAddr,
//	})
//
// Note:
//   - At least one of Address or ContractAddress must be provided
//   - This endpoint returns max 10000 records per call
//   - Use Page and Offset parameters to paginate through results
//   - TokenID field contains the specific token ID
//   - TokenValue field contains the amount transferred
func (c *HTTPClient) GetERC1155TokenTransfers(ctx context.Context, opts *GetERC1155TokenTransfersOpts) ([]RespERC1155TokenTransfer, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	if opts.Address == "" && opts.ContractAddress == "" {
		return nil, fmt.Errorf("at least one of Address or ContractAddress must be provided")
	}

	params := map[string]string{}

	if opts.Address != "" {
		params["address"] = opts.Address
	}
	if opts.ContractAddress != "" {
		params["contractaddress"] = opts.ContractAddress
	}

	params["page"] = strconv.FormatInt(opts.Page, 10)
	params["offset"] = strconv.FormatInt(opts.Offset, 10)
	params["startblock"] = strconv.FormatInt(opts.StartBlock, 10)
	params["endblock"] = strconv.FormatInt(opts.EndBlock, 10)
	params["sort"] = opts.Sort

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "token1155tx",
		params:          params,
		noFoundReturn:   []RespERC1155TokenTransfer{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespERC1155TokenTransfer
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Account Module - Additional Methods
// ============================================================================

// GetAddressFundedByOpts contains optional parameters for GetAddressFundedBy
type GetAddressFundedByOpts struct {
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

// GetAddressFundedBy returns the address that funded the specified address and its relative age
//
// This endpoint returns information about the address that initially funded the specified address,
// including the block number, timestamp, transaction hash, and funding amount. This is useful
// for tracking the origin of funds and understanding address relationships.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The address to check funding source for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespAddressFundedBy: Funding information including source address and transaction details
//   - error: Error if the request fails
//
// Example:
//
//	// Get funding information for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	info, err := client.GetAddressFundedBy(ctx, addr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if info != nil {
//	    fmt.Printf("Funding Address: %s\n", info.FundingAddress)
//	    fmt.Printf("Funding Transaction: %s\n", info.FundingTxn)
//	    fmt.Printf("Block Number: %s\n", info.Block)
//	    fmt.Printf("Time Stamp: %s\n", info.TimeStamp)
//	    fmt.Printf("Funding Amount: %s wei\n", info.Value)
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	info, err := client.GetAddressFundedBy(ctx, addr, &GetAddressFundedByOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Returns nil if no funding information found
//   - Value is returned in wei (smallest unit)
//   - TimeStamp is Unix timestamp
//   - Useful for tracking fund flows and address relationships
func (c *HTTPClient) GetAddressFundedBy(ctx context.Context, address string, opts *GetAddressFundedByOpts) (*RespAddressFundedBy, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"address": address,
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "fundedby",
		params:          params,
		noFoundReturn:   RespAddressFundedBy{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespAddressFundedBy
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBlocksValidatedByAddressOpts contains optional parameters for GetBlocksValidatedByAddress
type GetBlocksValidatedByAddressOpts struct {
	// BlockType specifies the type of blocks to retrieve
	// Options:
	//   - "blocks" (default): Returns canonical blocks validated by the address
	//   - "uncles": Returns uncle blocks validated by the address
	BlockType string `default:"blocks" json:"blocktype"`

	// Page number for pagination
	// Default: 1
	// Use this to navigate through multiple pages of results
	Page int64 `default:"1" json:"page"`

	// Offset is the number of blocks per page
	// Default: 10
	// Use this to control how many blocks are returned per page
	Offset int64 `default:"10" json:"offset"`

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

// GetBlocksValidatedByAddress returns list of blocks validated by an address
//
// This endpoint returns a list of blocks that were validated (mined) by a specific address.
// This is useful for analyzing miner activity, block rewards, and network participation.
// You can choose to get either canonical blocks or uncle blocks.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get validated blocks for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespBlockValidated: List of validated blocks with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get blocks validated by a miner address
//	minerAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	blocks, err := client.GetBlocksValidatedByAddress(ctx, minerAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, block := range blocks {
//	    fmt.Printf("Block Number: %s\n", block.BlockNumber)
//	    fmt.Printf("Time Stamp: %s\n", block.TimeStamp)
//	    fmt.Printf("Block Reward: %s wei\n", block.BlockReward)
//	}
//
//	// Get uncle blocks only
//	blockType := "uncles"
//	uncles, err := client.GetBlocksValidatedByAddress(ctx, minerAddr, &GetBlocksValidatedByAddressOpts{
//	    BlockType: &blockType,
//	})
//
//	// With pagination
//	page := 1
//	offset := 20
//	blocks, err := client.GetBlocksValidatedByAddress(ctx, minerAddr, &GetBlocksValidatedByAddressOpts{
//	    Page:   &page,
//	    Offset: &offset,
//	})
//
// Note:
//   - Returns empty slice if no blocks found
//   - BlockReward is in wei (smallest unit)
//   - TimeStamp is Unix timestamp
//   - Useful for analyzing miner performance and rewards
func (c *HTTPClient) GetBlocksValidatedByAddress(ctx context.Context, address string, opts *GetBlocksValidatedByAddressOpts) ([]RespBlockValidated, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"address":   address,
		"blocktype": opts.BlockType,
		"page":      strconv.FormatInt(opts.Page, 10),
		"offset":    strconv.FormatInt(opts.Offset, 10),
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "getminedblocks",
		params:          params,
		noFoundReturn:   []RespBlockValidated{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespBlockValidated
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBeaconChainWithdrawalsOpts contains optional parameters
type GetBeaconChainWithdrawalsOpts struct {
	// StartBlock is the starting block number to search from
	// Default: 0 (genesis block)
	StartBlock int64 `default:"0" json:"startblock"`

	// EndBlock is the ending block number to search to
	// Default: 999999999999 (latest block)
	EndBlock int64 `default:"999999999999" json:"endblock"`

	// Page number for pagination
	// Default: 1
	Page int64 `default:"1" json:"page"`

	// Offset is the number of records per page
	// Default: 100
	Offset int64 `default:"100" json:"offset"`

	// Sort order for the results
	// Default: "asc" (ascending)
	Sort string `default:"asc" json:"sort"`

	// ChainID specifies which blockchain network to query
	// Default: 1 (Ethereum mainnet)
	ChainID int64 `default:"1" json:"chainid"`

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock (wait until a token is available)
	OnLimitExceeded RateLimitBehavior `default:"" json:"on_limit_exceeded"`
}

// GetBeaconChainWithdrawals returns list of beacon chain withdrawals made to an address
//
// This endpoint returns a list of beacon chain withdrawals made to a specific address.
// Beacon chain withdrawals are part of Ethereum's proof-of-stake system, where validators
// can withdraw their staked ETH and rewards.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get withdrawals for
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespBeaconChainWithdrawal: List of beacon chain withdrawals with detailed information
//   - error: Error if the request fails
//
// Example:
//
//	// Get beacon chain withdrawals for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	withdrawals, err := client.GetBeaconChainWithdrawals(ctx, addr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, withdrawal := range withdrawals {
//	    fmt.Printf("Withdrawal Index: %s\n", withdrawal.WithdrawalIndex)
//	    fmt.Printf("Validator Index: %s\n", withdrawal.ValidatorIndex)
//	    fmt.Printf("Address: %s\n", withdrawal.Address)
//	    fmt.Printf("Amount: %s Gwei\n", withdrawal.Amount)
//	    fmt.Printf("Block Number: %s\n", withdrawal.BlockNumber)
//	    fmt.Printf("Timestamp: %s\n", withdrawal.Timestamp)
//	}
//
//	// With custom block range
//	startBlock := 18000000
//	endBlock := 18000100
//	withdrawals, err := client.GetBeaconChainWithdrawals(ctx, addr, &GetBeaconChainWithdrawalsOpts{
//	    StartBlock: &startBlock,
//	    EndBlock:   &endBlock,
//	})
//
// Note:
//   - Returns empty slice if no withdrawals found
//   - Amount is in Gwei (1 Gwei = 10^9 wei)
//   - Only available for Ethereum mainnet
//   - Useful for tracking validator rewards and withdrawals
func (c *HTTPClient) GetBeaconChainWithdrawals(ctx context.Context, address string, opts *GetBeaconChainWithdrawalsOpts) ([]RespBeaconChainWithdrawal, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"address":    address,
		"startblock": strconv.FormatInt(opts.StartBlock, 10),
		"endblock":   strconv.FormatInt(opts.EndBlock, 10),
		"page":       strconv.FormatInt(opts.Page, 10),
		"offset":     strconv.FormatInt(opts.Offset, 10),
		"sort":       opts.Sort,
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "txsBeaconWithdrawal",
		params:          params,
		noFoundReturn:   []RespBeaconChainWithdrawal{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespBeaconChainWithdrawal
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetEthBalanceByBlockNumberOpts contains optional parameters for GetEthBalanceByBlockNumber
type GetEthBalanceByBlockNumberOpts struct {
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

// GetEthBalanceByBlockNumber returns historical Eth balance for a single address at a specific block number
//
// This endpoint returns the ETH balance of an address at a specific block number in the past.
// This is useful for historical analysis, auditing, and understanding balance changes over time.
// The balance is returned in wei (the smallest unit of ETH).
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: Address to get balance for
//   - blockNo: Block number to get balance at
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Balance in wei as a string
//   - error: Error if the request fails
//
// Example:
//
//	// Get historical balance for an address
//	addr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	blockNumber := 18000000
//	balance, err := client.GetEthBalanceByBlockNumber(ctx, addr, blockNumber, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Balance at block %d: %s wei\n", blockNumber, balance)
//
//	// Convert wei to ETH (1 ETH = 10^18 wei)
//	balanceWei, _ := new(big.Int).SetString(balance, 10)
//	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balanceWei), big.NewFloat(1e18))
//	fmt.Printf("Balance in ETH: %s\n", balanceEth.String())
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	balance, err := client.GetEthBalanceByBlockNumber(ctx, addr, blockNumber, &GetEthBalanceByBlockNumberOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Balance is returned in wei (smallest unit)
//   - Returns "0" if address not found or has no balance
//   - Useful for historical balance analysis and auditing
func (c *HTTPClient) GetEthBalanceByBlockNumber(ctx context.Context, address string, blockNo int64, opts *GetEthBalanceByBlockNumberOpts) (string, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return "", err
	}

	params := map[string]string{
		"address": address,
		"blockno": strconv.FormatInt(blockNo, 10),
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "account",
		action:          "balancehistory",
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

// GetContractCreatorAndCreationOpts contains optional parameters for GetContractCreatorAndCreation
type GetContractCreatorAndCreationOpts struct {
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

// GetContractCreatorAndCreation returns the creator address and creation transaction hash for contracts
//
// This endpoint returns information about who created specific contracts and when they were created.
// This is useful for tracking contract deployment history, analyzing developer activity, and
// understanding contract relationships.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - contractAddresses: List of contract addresses to check (maximum 5 addresses)
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespContractCreationAndCreation: List of contract creation details
//   - error: Error if the request fails or more than 5 addresses provided
//
// Example:
//
//	// Get creation info for multiple contracts
//	contracts := []string{
//	    "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0",
//	    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
//	}
//	creations, err := client.GetContractCreatorAndCreation(ctx, contracts, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, creation := range creations {
//	    fmt.Printf("Contract: %s\n", creation.ContractAddress)
//	    fmt.Printf("Creator: %s\n", creation.ContractCreator)
//	    fmt.Printf("Creation Tx: %s\n", creation.TxHash)
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	creations, err := client.GetContractCreatorAndCreation(ctx, contracts, &GetContractCreatorAndCreationOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Maximum 5 contract addresses per call
//   - Returns empty slice if no creation info found
//   - Useful for tracking contract deployment history
//   - Helps identify contract relationships and developer activity
func (c *HTTPClient) GetContractCreatorAndCreation(ctx context.Context, contractAddresses []string, opts *GetContractCreatorAndCreationOpts) ([]RespContractCreationAndCreation, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	if len(contractAddresses) > 5 {
		return nil, fmt.Errorf("maximum 5 contract addresses allowed")
	}

	params := map[string]string{
		"contractaddresses": strings.Join(contractAddresses, ","),
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "getcontractcreation",
		params:          params,
		noFoundReturn:   []RespContractCreationAndCreation{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespContractCreationAndCreation
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
