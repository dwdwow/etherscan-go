package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Gas Tracker Module
// ============================================================================

// GetConfirmationTimeEstimateOpts contains optional parameters for GetConfirmationTimeEstimate
type GetConfirmationTimeEstimateOpts struct {
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

// GetConfirmationTimeEstimate returns the estimated time, in seconds, for a transaction to be confirmed on the blockchain
//
// This endpoint returns the estimated time for a transaction to be confirmed on the blockchain
// based on the provided gas price. This is useful for estimating transaction confirmation times
// and optimizing gas prices for faster confirmation.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - gasPrice: The price paid per unit of gas, in wei
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Estimated confirmation time in seconds
//   - error: Error if the request fails
//
// Example:
//
//	// Get confirmation time estimate for a gas price
//	gasPrice := 20000000000 // 20 Gwei in wei
//	estimate, err := client.GetConfirmationTimeEstimate(ctx, gasPrice, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Estimated confirmation time: %s seconds\n", estimate)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	estimate, err := client.GetConfirmationTimeEstimate(ctx, gasPrice, &GetConfirmationTimeEstimateOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Gas price should be in wei (smallest unit)
//   - Returns estimated time in seconds as string
//   - Useful for optimizing transaction confirmation times
func (c *HTTPClient) GetConfirmationTimeEstimate(ctx context.Context, gasPrice int64, opts *GetConfirmationTimeEstimateOpts) (string, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return "", err
	}

	params := map[string]string{
		"gasprice": strconv.FormatInt(gasPrice, 10),
	}

	if opts != nil && opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "gastracker",
		action:          "gasestimate",
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

// GetGasOracleOpts contains optional parameters for GetGasOracle
type GetGasOracleOpts struct {
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

// GetGasOracle returns the current Safe, Proposed and Fast gas prices
//
// This endpoint returns the current gas price recommendations (Safe, Proposed, Fast) and
// network statistics. Post EIP-1559, gas price recommendations are modeled as Priority Fees,
// and the response includes base fee information and network utilization.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespGasOracle: Current gas prices and network statistics
//   - error: Error if the request fails
//
// Example:
//
//	// Get current gas prices
//	gasOracle, err := client.GetGasOracle(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Safe Gas Price: %s Gwei\n", gasOracle.SafeGasPrice)
//	fmt.Printf("Proposed Gas Price: %s Gwei\n", gasOracle.ProposedGasPrice)
//	fmt.Printf("Fast Gas Price: %s Gwei\n", gasOracle.FastGasPrice)
//	fmt.Printf("Base Fee: %s Gwei\n", gasOracle.SuggestBaseFee)
//	fmt.Printf("Gas Used Ratio: %s\n", gasOracle.GasUsedRatio)
//
// Note:
//   - Gas prices are in Gwei
//   - Post EIP-1559: Safe/Proposed/Fast are Priority Fees
//   - SuggestBaseFee shows the base fee of the next pending block
//   - GasUsedRatio estimates network utilization
func (c *HTTPClient) GetGasOracle(ctx context.Context, opts *GetGasOracleOpts) (*RespGasOracle, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{}

	if opts != nil && opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "gastracker",
		action:          "gasoracle",
		params:          params,
		noFoundReturn:   RespGasOracle{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespGasOracle
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDailyAverageGasLimitOpts contains optional parameters for GetDailyAverageGasLimit
type GetDailyAverageGasLimitOpts struct {
	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by date in ascending order (oldest first)
	//   - "desc": Sort by date in descending order (newest first)
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

// GetDailyAverageGasLimit returns the historical daily average gas limit of the Ethereum network
//
// This endpoint returns the historical daily average gas limit of the Ethereum network
// within a specified date range. Gas limit is the maximum amount of gas that can be
// used in a block, and this data is useful for analyzing network capacity over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startDate: Starting date in yyyy-MM-dd format (e.g., "2019-01-31")
//   - endDate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgGasLimit: List of daily average gas limit data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average gas limit for a date range
//	startDate := "2019-01-31"
//	endDate := "2019-02-28"
//	gasLimits, err := client.GetDailyAverageGasLimit(ctx, startDate, endDate, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, limit := range gasLimits {
//	    fmt.Printf("Date: %s\n", limit.UTCDate)
//	    fmt.Printf("Average Gas Limit: %s\n", limit.AvgGasLimit)
//	}
//
//	// With custom sort order
//	sort := "desc"
//	gasLimits, err := client.GetDailyAverageGasLimit(ctx, startDate, endDate, &GetDailyAverageGasLimitOpts{
//	    Sort: &sort,
//	})
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
//   - Useful for analyzing network capacity over time
func (c *HTTPClient) GetDailyAverageGasLimit(ctx context.Context, startDate, endDate string, opts *GetDailyAverageGasLimitOpts) ([]RespDailyAvgGasLimit, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"startdate": startDate,
		"enddate":   endDate,
		"sort":      "asc",
	}

	if opts != nil && opts.Sort != "" {
		params["sort"] = opts.Sort
	}
	if opts != nil && opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavggaslimit",
		params:          params,
		noFoundReturn:   []RespDailyAvgGasLimit{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgGasLimit
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyTotalGasUsedOpts contains optional parameters for GetDailyTotalGasUsed
type GetDailyTotalGasUsedOpts struct {
	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by date in ascending order (oldest first)
	//   - "desc": Sort by date in descending order (newest first)
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

// GetDailyTotalGasUsed returns the total amount of gas used daily for transactions on the Ethereum network
//
// This endpoint returns the total amount of gas used daily for transactions on the Ethereum
// network within a specified date range. This data is useful for analyzing network activity,
// transaction volume, and gas consumption patterns over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startDate: Starting date in yyyy-MM-dd format (e.g., "2019-01-31")
//   - endDate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyTotalGasUsed: List of daily total gas used data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily total gas used for a date range
//	startDate := "2019-01-31"
//	endDate := "2019-02-28"
//	gasUsed, err := client.GetDailyTotalGasUsed(ctx, startDate, endDate, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, usage := range gasUsed {
//	    fmt.Printf("Date: %s\n", usage.UTCDate)
//	    fmt.Printf("Total Gas Used: %s\n", usage.TotalGasUsed)
//	}
//
//	// With custom sort order
//	sort := "desc"
//	gasUsed, err := client.GetDailyTotalGasUsed(ctx, startDate, endDate, &GetDailyTotalGasUsedOpts{
//	    Sort: &sort,
//	})
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
//   - Useful for analyzing network activity and gas consumption
func (c *HTTPClient) GetDailyTotalGasUsed(ctx context.Context, startDate, endDate string, opts *GetDailyTotalGasUsedOpts) ([]RespDailyTotalGasUsed, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"startdate": startDate,
		"enddate":   endDate,
		"sort":      "asc",
	}

	if opts != nil && opts.Sort != "" {
		params["sort"] = opts.Sort
	}
	if opts != nil && opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailygasused",
		params:          params,
		noFoundReturn:   []RespDailyTotalGasUsed{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyTotalGasUsed
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAverageGasPriceOpts contains optional parameters for GetDailyAverageGasPrice
type GetDailyAverageGasPriceOpts struct {
	// Sort order for the results
	// Options:
	//   - "asc" (default): Sort by date in ascending order (oldest first)
	//   - "desc": Sort by date in descending order (newest first)
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

// GetDailyAverageGasPrice returns the daily average gas price used on the Ethereum network
//
// This endpoint returns the daily average gas price used on the Ethereum network within
// a specified date range. This data is useful for analyzing gas price trends, network
// congestion patterns, and transaction cost analysis over time.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - startDate: Starting date in yyyy-MM-dd format (e.g., "2019-01-31")
//   - endDate: Ending date in yyyy-MM-dd format (e.g., "2019-02-28")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespDailyAvgGasPrice: List of daily average gas price data
//   - error: Error if the request fails
//
// Example:
//
//	// Get daily average gas price for a date range
//	startDate := "2019-01-31"
//	endDate := "2019-02-28"
//	gasPrices, err := client.GetDailyAverageGasPrice(ctx, startDate, endDate, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, price := range gasPrices {
//	    fmt.Printf("Date: %s\n", price.UTCDate)
//	    fmt.Printf("Average Gas Price: %s Gwei\n", price.AvgGasPrice)
//	}
//
//	// With custom sort order
//	sort := "desc"
//	gasPrices, err := client.GetDailyAverageGasPrice(ctx, startDate, endDate, &GetDailyAverageGasPriceOpts{
//	    Sort: &sort,
//	})
//
// Note:
//   - Date format must be yyyy-MM-dd
//   - Returns empty slice if no data found
//   - Gas prices are in Gwei
//   - Useful for analyzing gas price trends and network congestion
func (c *HTTPClient) GetDailyAverageGasPrice(ctx context.Context, startDate, endDate string, opts *GetDailyAverageGasPriceOpts) ([]RespDailyAvgGasPrice, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{
		"startdate": startDate,
		"enddate":   endDate,
		"sort":      "asc",
	}

	if opts != nil && opts.Sort != "" {
		params["sort"] = opts.Sort
	}
	if opts != nil && opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	var onLimitExceeded RateLimitBehavior
	if opts != nil {
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "stats",
		action:          "dailyavggasprice",
		params:          params,
		noFoundReturn:   []RespDailyAvgGasPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgGasPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
