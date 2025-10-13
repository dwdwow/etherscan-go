package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Stats Module - Daily Statistics
// ============================================================================

// GetDailyBlockCountRewardsOpts contains optional parameters
type GetDailyBlockCountRewardsOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyBlockCountRewards returns daily block count and rewards within a date range
func (c *HTTPClient) GetDailyBlockCountRewards(ctx context.Context, startdate, enddate string, opts *GetDailyBlockCountRewardsOpts) ([]RespDailyBlockCountReward, error) {
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
		action:          "dailyblkcount",
		params:          params,
		noFoundReturn:   []RespDailyBlockCountReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyBlockCountReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyBlockRewardsOpts contains optional parameters
type GetDailyBlockRewardsOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyBlockRewards returns daily block rewards distributed to miners within a date range
func (c *HTTPClient) GetDailyBlockRewards(ctx context.Context, startdate, enddate string, opts *GetDailyBlockRewardsOpts) ([]RespDailyBlockReward, error) {
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
		action:          "dailyblockrewards",
		params:          params,
		noFoundReturn:   []RespDailyBlockReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyBlockReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgBlockTimeOpts contains optional parameters
type GetDailyAvgBlockTimeOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgBlockTime returns daily average time for a block to be included in the blockchain
func (c *HTTPClient) GetDailyAvgBlockTime(ctx context.Context, startdate, enddate string, opts *GetDailyAvgBlockTimeOpts) ([]RespDailyAvgTimeBlockMined, error) {
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
		action:          "dailyavgblocktime",
		params:          params,
		noFoundReturn:   []RespDailyAvgTimeBlockMined{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgTimeBlockMined
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyUncleBlockCountAndRewardsOpts contains optional parameters
type GetDailyUncleBlockCountAndRewardsOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyUncleBlockCountAndRewards returns daily uncle block count and rewards
func (c *HTTPClient) GetDailyUncleBlockCountAndRewards(ctx context.Context, startdate, enddate string, opts *GetDailyUncleBlockCountAndRewardsOpts) ([]RespDailyUncleBlockCountAndReward, error) {
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
		action:          "dailyuncleblkcount",
		params:          params,
		noFoundReturn:   []RespDailyUncleBlockCountAndReward{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyUncleBlockCountAndReward
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Stats Module - Supply and Price
// ============================================================================

// GetTotalEthSupplyOpts contains optional parameters
type GetTotalEthSupplyOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetTotalEthSupply returns the current amount of Eth in circulation excluding ETH2 Staking rewards and EIP1559 burnt fees
func (c *HTTPClient) GetTotalEthSupply(ctx context.Context, opts *GetTotalEthSupplyOpts) (string, error) {
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
		module:          "stats",
		action:          "ethsupply",
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

// GetTotalEth2SupplyOpts contains optional parameters
type GetTotalEth2SupplyOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetTotalEth2Supply returns the current amount of Eth in circulation, ETH2 Staking rewards, EIP1559 burnt fees
func (c *HTTPClient) GetTotalEth2Supply(ctx context.Context, opts *GetTotalEth2SupplyOpts) (string, error) {
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
		module:          "stats",
		action:          "ethsupply2",
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

// GetEthPriceOpts contains optional parameters
type GetEthPriceOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetEthPrice returns the latest price of the native/gas token
func (c *HTTPClient) GetEthPrice(ctx context.Context, opts *GetEthPriceOpts) (*RespEthPrice, error) {
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
		module:          "stats",
		action:          "ethprice",
		params:          params,
		noFoundReturn:   RespEthPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespEthPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEthHistoricalPricesOpts contains optional parameters
type GetEthHistoricalPricesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetEthHistoricalPrices returns the historical price of 1 ETH
func (c *HTTPClient) GetEthHistoricalPrices(ctx context.Context, startdate, enddate string, opts *GetEthHistoricalPricesOpts) ([]RespEthHistoricalPrice, error) {
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
		action:          "ethdailyprice",
		params:          params,
		noFoundReturn:   []RespEthHistoricalPrice{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEthHistoricalPrice
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Stats Module - Network Statistics
// ============================================================================

// GetEthereumNodesSizeOpts contains optional parameters
type GetEthereumNodesSizeOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetEthereumNodesSize returns the size of the Ethereum blockchain, in bytes, over a date range
func (c *HTTPClient) GetEthereumNodesSize(ctx context.Context, startdate, enddate, clienttype, syncmode, sort string, opts *GetEthereumNodesSizeOpts) ([]RespEtheumNodeSize, error) {
	params := map[string]string{
		"startdate":  startdate,
		"enddate":    enddate,
		"clienttype": clienttype,
		"syncmode":   syncmode,
		"sort":       sort,
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
		module:          "stats",
		action:          "chainsize",
		params:          params,
		noFoundReturn:   []RespEtheumNodeSize{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespEtheumNodeSize
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeCountOpts contains optional parameters
type GetNodeCountOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetNodeCount returns the total number of discoverable Ethereum nodes
func (c *HTTPClient) GetNodeCount(ctx context.Context, opts *GetNodeCountOpts) (*RespNodeCount, error) {
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
		module:          "stats",
		action:          "nodecount",
		params:          params,
		noFoundReturn:   RespNodeCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespNodeCount
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDailyTxFeesOpts contains optional parameters
type GetDailyTxFeesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyTxFees returns the amount of transaction fees paid to miners per day
func (c *HTTPClient) GetDailyTxFees(ctx context.Context, startdate, enddate string, opts *GetDailyTxFeesOpts) ([]RespDailyTxFee, error) {
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
		action:          "dailytxnfee",
		params:          params,
		noFoundReturn:   []RespDailyTxFee{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyTxFee
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyNewAddressesOpts contains optional parameters
type GetDailyNewAddressesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyNewAddresses returns the number of new Ethereum addresses created per day
func (c *HTTPClient) GetDailyNewAddresses(ctx context.Context, startdate, enddate string, opts *GetDailyNewAddressesOpts) ([]RespDailyNewAddress, error) {
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
		action:          "dailynewaddress",
		params:          params,
		noFoundReturn:   []RespDailyNewAddress{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyNewAddress
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyNetworkUtilizationsOpts contains optional parameters
type GetDailyNetworkUtilizationsOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyNetworkUtilizations returns the daily average gas used over gas limit, in percentage
func (c *HTTPClient) GetDailyNetworkUtilizations(ctx context.Context, startdate, enddate string, opts *GetDailyNetworkUtilizationsOpts) ([]RespDailyNetworkUtilization, error) {
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
		action:          "dailynetutilization",
		params:          params,
		noFoundReturn:   []RespDailyNetworkUtilization{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyNetworkUtilization
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgHashratesOpts contains optional parameters
type GetDailyAvgHashratesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgHashrates returns the historical measure of processing power of the Ethereum network
func (c *HTTPClient) GetDailyAvgHashrates(ctx context.Context, startdate, enddate string, opts *GetDailyAvgHashratesOpts) ([]RespDailyAvgHashrate, error) {
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
		action:          "dailyavghashrate",
		params:          params,
		noFoundReturn:   []RespDailyAvgHashrate{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgHashrate
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyTxCountsOpts contains optional parameters
type GetDailyTxCountsOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyTxCounts returns the number of transactions performed on the Ethereum blockchain per day
func (c *HTTPClient) GetDailyTxCounts(ctx context.Context, startdate, enddate string, opts *GetDailyTxCountsOpts) ([]RespDailyTxCount, error) {
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
		action:          "dailytx",
		params:          params,
		noFoundReturn:   []RespDailyTxCount{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyTxCount
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDailyAvgDifficultiesOpts contains optional parameters
type GetDailyAvgDifficultiesOpts struct {
	Sort            *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetDailyAvgDifficulties returns the historical mining difficulty of the Ethereum network
func (c *HTTPClient) GetDailyAvgDifficulties(ctx context.Context, startdate, enddate string, opts *GetDailyAvgDifficultiesOpts) ([]RespDailyAvgDifficulty, error) {
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
		action:          "dailyavgnetdifficulty",
		params:          params,
		noFoundReturn:   []RespDailyAvgDifficulty{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespDailyAvgDifficulty
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
