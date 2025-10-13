package etherscan

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ============================================================================
// Admin Module - Address Tagging & Labeling
// ============================================================================

// GetAddressTagOpts contains optional parameters
type GetAddressTagOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetAddressTag returns address name tag and metadata
//
// Note: This endpoint is throttled to 2 calls/second regardless of API Pro tier
func (c *HTTPClient) GetAddressTag(ctx context.Context, addresses []string, opts *GetAddressTagOpts) ([]RespAddressTag, error) {
	if len(addresses) > 100 {
		return nil, fmt.Errorf("maximum 100 addresses allowed")
	}

	params := map[string]string{
		"address": strings.Join(addresses, ","),
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = fmt.Sprintf("%d", *opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "nametag",
		action:          "getaddresstag",
		params:          params,
		noFoundReturn:   []RespAddressTag{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespAddressTag
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLabelMasterlistOpts contains optional parameters
type GetLabelMasterlistOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetLabelMasterlist returns the masterlist of available label groupings
func (c *HTTPClient) GetLabelMasterlist(ctx context.Context, opts *GetLabelMasterlistOpts) ([]RespLabelMaster, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = fmt.Sprintf("%d", *opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "nametag",
		action:          "getlabelmasterlist",
		params:          params,
		baseURL:         APIAshx,
		noFoundReturn:   []RespLabelMaster{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespLabelMaster
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExportSpecificLabelCSV returns addresses filtered by a specific label in CSV format
func (c *HTTPClient) ExportSpecificLabelCSV(ctx context.Context, label string) (string, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&label=%s&format=csv&apikey=%s",
		APIAshx, label, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// ExportOFACSanctionedRelatedLabelsCSV returns addresses sanctioned by the U.S. Department of the Treasury's
// Office of Foreign Assets Control's Specially Designated Nationals list in CSV format
func (c *HTTPClient) ExportOFACSanctionedRelatedLabelsCSV(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&label=ofac-sanctioned&format=csv&apikey=%s",
		APIAshx, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// ExportAllAddressTagsCSV exports a complete CSV list of ALL address name tags and/or labels
func (c *HTTPClient) ExportAllAddressTagsCSV(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&format=csv&apikey=%s",
		APIAshx, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetLatestCSVBatchNumberOpts contains optional parameters
type GetLatestCSVBatchNumberOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// GetLatestCSVBatchNumber gets the latest running number for CSV Export
func (c *HTTPClient) GetLatestCSVBatchNumber(ctx context.Context, opts *GetLatestCSVBatchNumberOpts) ([]RespLatestCSVBatchNumber, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = fmt.Sprintf("%d", *opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "nametag",
		action:          "getcurrentbatch",
		params:          params,
		noFoundReturn:   []RespLatestCSVBatchNumber{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespLatestCSVBatchNumber
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ============================================================================
// Admin Module - API Credit & Chain Info
// ============================================================================

// CheckCreditUsageOpts contains optional parameters
type CheckCreditUsageOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// CheckCreditUsage returns information about API credit usage and limits
func (c *HTTPClient) CheckCreditUsage(ctx context.Context, opts *CheckCreditUsageOpts) (*RespCreditUsage, error) {
	params := map[string]string{}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = fmt.Sprintf("%d", *opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "getapilimit",
		action:          "getapilimit",
		params:          params,
		noFoundReturn:   RespCreditUsage{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result RespCreditUsage
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
