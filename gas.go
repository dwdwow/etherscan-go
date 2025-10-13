package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Gas Tracker Module
// ============================================================================

// GetConfirmationTimeEstimateOpts contains optional parameters
type GetConfirmationTimeEstimateOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetConfirmationTimeEstimate returns the estimated time, in seconds, for a transaction to be confirmed on the blockchain
func (c *HTTPClient) GetConfirmationTimeEstimate(ctx context.Context, gasprice int, opts *GetConfirmationTimeEstimateOpts) (string, error) {
	params := map[string]string{
		"gasprice": strconv.Itoa(gasprice),
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

// GetGasOracleOpts contains optional parameters
type GetGasOracleOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetGasOracle returns the current Safe, Proposed and Fast gas prices
func (c *HTTPClient) GetGasOracle(ctx context.Context, opts *GetGasOracleOpts) (*RespGasOracle, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// GetDailyAverageGasLimitOpts contains optional parameters
type GetDailyAverageGasLimitOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAverageGasLimit returns the historical daily average gas limit of the Ethereum network
func (c *HTTPClient) GetDailyAverageGasLimit(ctx context.Context, startdate, enddate string, opts *GetDailyAverageGasLimitOpts) ([]RespDailyAvgGasLimit, error) {
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

// GetDailyTotalGasUsedOpts contains optional parameters
type GetDailyTotalGasUsedOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyTotalGasUsed returns the total amount of gas used daily for transactions on the Ethereum network
func (c *HTTPClient) GetDailyTotalGasUsed(ctx context.Context, startdate, enddate string, opts *GetDailyTotalGasUsedOpts) ([]RespDailyTotalGasUsed, error) {
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

// GetDailyAverageGasPriceOpts contains optional parameters
type GetDailyAverageGasPriceOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAverageGasPrice returns the daily average gas price used on the Ethereum network
func (c *HTTPClient) GetDailyAverageGasPrice(ctx context.Context, startdate, enddate string, opts *GetDailyAverageGasPriceOpts) ([]RespDailyAvgGasPrice, error) {
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
