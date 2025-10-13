package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Geth/Parity Proxy Module
// ============================================================================

// RpcEthBlockNumberOpts contains optional parameters
type RpcEthBlockNumberOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthBlockNumber returns the number of the most recent block
func (c *HTTPClient) RpcEthBlockNumber(ctx context.Context, opts *RpcEthBlockNumberOpts) (string, error) {
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

// RpcEthBlockByNumberOpts contains optional parameters
type RpcEthBlockByNumberOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthBlockByNumber returns information about a block by block number
func (c *HTTPClient) RpcEthBlockByNumber(ctx context.Context, tag string, boolean bool, opts *RpcEthBlockByNumberOpts) (*RespEthBlockInfo, error) {
	params := map[string]string{
		"tag":     tag,
		"boolean": fmt.Sprintf("%t", boolean),
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

// RpcEthUncleByBlockNumberAndIndexOpts contains optional parameters
type RpcEthUncleByBlockNumberAndIndexOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthUncleByBlockNumberAndIndex returns information about an uncle block by block number and index
func (c *HTTPClient) RpcEthUncleByBlockNumberAndIndex(ctx context.Context, tag, index string, opts *RpcEthUncleByBlockNumberAndIndexOpts) (*RespEthUncleBlockInfo, error) {
	params := map[string]string{
		"tag":   tag,
		"index": index,
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

// RpcEthBlockTransactionCountByNumberOpts contains optional parameters
type RpcEthBlockTransactionCountByNumberOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthBlockTransactionCountByNumber returns the number of transactions in a block by block number
func (c *HTTPClient) RpcEthBlockTransactionCountByNumber(ctx context.Context, tag string, opts *RpcEthBlockTransactionCountByNumberOpts) (string, error) {
	params := map[string]string{
		"tag": tag,
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
		module:          "proxy",
		action:          "eth_getBlockTxCountByNumber",
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

// RpcEthTransactionByHashOpts contains optional parameters
type RpcEthTransactionByHashOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthTransactionByHash returns information about a transaction by its hash
func (c *HTTPClient) RpcEthTransactionByHash(ctx context.Context, txhash string, opts *RpcEthTransactionByHashOpts) (*RespEthTxInfo, error) {
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
		module:          "proxy",
		action:          "eth_getTxByHash",
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

// RpcEthTransactionByBlockNumberAndIndexOpts contains optional parameters
type RpcEthTransactionByBlockNumberAndIndexOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position
func (c *HTTPClient) RpcEthTransactionByBlockNumberAndIndex(ctx context.Context, tag, index string, opts *RpcEthTransactionByBlockNumberAndIndexOpts) (*RespEthTxInfo, error) {
	params := map[string]string{
		"tag":   tag,
		"index": index,
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
		module:          "proxy",
		action:          "eth_getTxByBlockNumberAndIndex",
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

// RpcEthTransactionCountOpts contains optional parameters
type RpcEthTransactionCountOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthTransactionCount returns the number of transactions performed by an address
func (c *HTTPClient) RpcEthTransactionCount(ctx context.Context, address, tag string, opts *RpcEthTransactionCountOpts) (string, error) {
	params := map[string]string{
		"address": address,
		"tag":     tag,
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
		module:          "proxy",
		action:          "eth_getTxCount",
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

// RpcEthSendRawTransactionOpts contains optional parameters
type RpcEthSendRawTransactionOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthSendRawTransaction submits a pre-signed transaction for broadcast to the Ethereum network
func (c *HTTPClient) RpcEthSendRawTransaction(ctx context.Context, hex string, opts *RpcEthSendRawTransactionOpts) (string, error) {
	params := map[string]string{
		"hex": hex,
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
		module:          "proxy",
		action:          "eth_sendRawTx",
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

// RpcEthTransactionReceiptOpts contains optional parameters
type RpcEthTransactionReceiptOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthTransactionReceipt returns the receipt of a transaction by transaction hash
func (c *HTTPClient) RpcEthTransactionReceipt(ctx context.Context, txhash string, opts *RpcEthTransactionReceiptOpts) (*RespEthTxReceiptInfo, error) {
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
		module:          "proxy",
		action:          "eth_getTxReceipt",
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

// RpcEthCallOpts contains optional parameters
type RpcEthCallOpts struct {
	Tag             *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthCall executes a new message call immediately without creating a transaction on the block chain
func (c *HTTPClient) RpcEthCall(ctx context.Context, to, data string, opts *RpcEthCallOpts) (string, error) {
	params := map[string]string{
		"to":   to,
		"data": data,
		"tag":  "latest",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Tag != nil {
			params["tag"] = *opts.Tag
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// RpcEthGetCodeOpts contains optional parameters
type RpcEthGetCodeOpts struct {
	Tag             *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthGetCode returns code at a given address
func (c *HTTPClient) RpcEthGetCode(ctx context.Context, address string, opts *RpcEthGetCodeOpts) (string, error) {
	params := map[string]string{
		"address": address,
		"tag":     "latest",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Tag != nil {
			params["tag"] = *opts.Tag
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// RpcEthGetStorageAtOpts contains optional parameters
type RpcEthGetStorageAtOpts struct {
	Tag             *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthGetStorageAt returns the value from a storage position at a given address
func (c *HTTPClient) RpcEthGetStorageAt(ctx context.Context, address, position string, opts *RpcEthGetStorageAtOpts) (string, error) {
	params := map[string]string{
		"address":  address,
		"position": position,
		"tag":      "latest",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Tag != nil {
			params["tag"] = *opts.Tag
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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

// RpcEthGetGasPriceOpts contains optional parameters
type RpcEthGetGasPriceOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthGetGasPrice returns the current price per gas in wei
func (c *HTTPClient) RpcEthGetGasPrice(ctx context.Context, opts *RpcEthGetGasPriceOpts) (string, error) {
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

// RpcEthEstimateGasOpts contains optional parameters
type RpcEthEstimateGasOpts struct {
	Value           *string
	Gas             *string
	GasPrice        *string
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// RpcEthEstimateGas makes a call or transaction, which won't be added to the blockchain and returns the used gas
func (c *HTTPClient) RpcEthEstimateGas(ctx context.Context, to, data string, opts *RpcEthEstimateGasOpts) (string, error) {
	params := map[string]string{
		"to":   to,
		"data": data,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.Value != nil {
			params["value"] = *opts.Value
		}
		if opts.Gas != nil {
			params["gas"] = *opts.Gas
		}
		if opts.GasPrice != nil {
			params["gasPrice"] = *opts.GasPrice
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
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
