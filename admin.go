package etherscan

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// ============================================================================
// Admin Module - Address Tagging & Labeling
// ============================================================================

// GetAddressTagOpts contains optional parameters for GetAddressTag
type GetAddressTagOpts struct {
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

// GetAddressTag returns address name tag and metadata
//
// This endpoint returns name tags and metadata for addresses. Name tags are human-readable
// labels assigned to addresses by Etherscan, such as exchange names, contract names, etc.
// This is useful for identifying known addresses and their purposes.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - addresses: List of addresses to get tags for (maximum 100 addresses)
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespAddressTag: List of address tags with metadata
//   - error: Error if the request fails or more than 100 addresses provided
//
// Example:
//
//	// Get tags for multiple addresses
//	addresses := []string{
//	    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
//	    "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0",
//	}
//	tags, err := client.GetAddressTag(ctx, addresses, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, tag := range tags {
//	    fmt.Printf("Address: %s\n", tag.Address)
//	    fmt.Printf("Name Tag: %s\n", tag.NameTag)
//	    fmt.Printf("Tag Type: %s\n", tag.TagType)
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	tags, err := client.GetAddressTag(ctx, addresses, &GetAddressTagOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - This endpoint is throttled to 2 calls/second regardless of API Pro tier
//   - Maximum 100 addresses per call
//   - Returns empty slice if no tags found
//   - Useful for identifying known addresses and their purposes
func (c *HTTPClient) GetAddressTag(ctx context.Context, addresses []string, opts *GetAddressTagOpts) ([]RespAddressTag, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	if len(addresses) > 100 {
		return nil, fmt.Errorf("maximum 100 addresses allowed")
	}

	params := map[string]string{
		"address": strings.Join(addresses, ","),
	}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

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

// GetLabelMasterlistOpts contains optional parameters for GetLabelMasterlist
type GetLabelMasterlistOpts struct {
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

// GetLabelMasterlist returns the masterlist of available label groupings
//
// This endpoint returns the masterlist of available label groupings that can be used
// for filtering addresses. Labels are categories used to group addresses by their
// purpose or type (e.g., exchanges, contracts, etc.).
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespLabelMaster: List of available label groupings
//   - error: Error if the request fails
//
// Example:
//
//	// Get all available label groupings
//	labels, err := client.GetLabelMasterlist(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, label := range labels {
//	    fmt.Printf("Label: %s\n", label.Label)
//	    fmt.Printf("Description: %s\n", label.Description)
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	labels, err := client.GetLabelMasterlist(ctx, &GetLabelMasterlistOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Returns empty slice if no labels found
//   - Labels can be used with ExportSpecificLabelCSV to filter addresses
//   - Useful for discovering available address categories
func (c *HTTPClient) GetLabelMasterlist(ctx context.Context, opts *GetLabelMasterlistOpts) ([]RespLabelMaster, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

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
//
// This endpoint exports addresses that belong to a specific label category in CSV format.
// This is useful for bulk analysis of addresses by category (e.g., all exchange addresses).
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - label: The label category to filter by (e.g., "exchange", "contract")
//
// Returns:
//   - string: CSV data containing addresses and their metadata
//   - error: Error if the request fails
//
// Example:
//
//	// Export all exchange addresses
//	csvData, err := client.ExportSpecificLabelCSV(ctx, "exchange")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("CSV Data:\n%s\n", csvData)
//
//	// Export all contract addresses
//	csvData, err := client.ExportSpecificLabelCSV(ctx, "contract")
//
// Note:
//   - Returns CSV data as string
//   - Use GetLabelMasterlist to see available labels
//   - Useful for bulk address analysis by category
func (c *HTTPClient) ExportSpecificLabelCSV(ctx context.Context, label string) ([]byte, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&label=%s&format=csv&apikey=%s",
		APIAshx, label, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

// ExportOFACSanctionedRelatedLabelsCSV returns addresses sanctioned by the U.S. Department of the Treasury's
// Office of Foreign Assets Control's Specially Designated Nationals list in CSV format
//
// This endpoint exports addresses that are sanctioned by OFAC (Office of Foreign Assets Control)
// in CSV format. These are addresses associated with entities on the SDN (Specially Designated Nationals) list.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//
// Returns:
//   - string: CSV data containing OFAC-sanctioned addresses and their metadata
//   - error: Error if the request fails
//
// Example:
//
//	// Export OFAC-sanctioned addresses
//	csvData, err := client.ExportOFACSanctionedRelatedLabelsCSV(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("OFAC CSV Data:\n%s\n", csvData)
//
// Note:
//   - Returns CSV data as string
//   - Contains addresses sanctioned by OFAC
//   - Useful for compliance and risk assessment
func (c *HTTPClient) ExportOFACSanctionedRelatedLabelsCSV(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&label=ofac-sanctioned&format=csv&apikey=%s",
		APIAshx, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

// ExportAllAddressTagsCSV exports a complete CSV list of ALL address name tags and/or labels
//
// This endpoint exports all address name tags and labels in CSV format. This includes
// all categorized addresses with their metadata and labels.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//
// Returns:
//   - string: CSV data containing all address tags and labels
//   - error: Error if the request fails
//
// Example:
//
//	// Export all address tags
//	csvData, err := client.ExportAllAddressTagsCSV(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("All Address Tags CSV:\n%s\n", csvData)
//
// Note:
//   - Returns CSV data as string
//   - Contains all address tags and labels
//   - Large dataset - may take time to download
func (c *HTTPClient) ExportAllAddressTagsCSV(ctx context.Context) ([]byte, error) {
	url := fmt.Sprintf("%s?module=nametag&action=exportaddresstags&format=csv&apikey=%s",
		APIAshx, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

// GetLatestCSVBatchNumberOpts contains optional parameters for GetLatestCSVBatchNumber
type GetLatestCSVBatchNumberOpts struct {
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

// GetLatestCSVBatchNumber gets the latest running number for CSV Export
//
// This endpoint returns the latest batch number for CSV exports. This is useful for
// tracking the version of exported data and ensuring you have the most recent dataset.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespLatestCSVBatchNumber: List containing the latest batch number information
//   - error: Error if the request fails
//
// Example:
//
//	// Get the latest CSV batch number
//	batchInfo, err := client.GetLatestCSVBatchNumber(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if len(batchInfo) > 0 {
//	    fmt.Printf("Latest batch number: %s\n", batchInfo[0].BatchNumber)
//	}
//
// Note:
//   - Returns empty slice if no batch info found
//   - Useful for tracking CSV export versions
func (c *HTTPClient) GetLatestCSVBatchNumber(ctx context.Context, opts *GetLatestCSVBatchNumberOpts) ([]RespLatestCSVBatchNumber, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

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

// CheckCreditUsageOpts contains optional parameters for CheckCreditUsage
type CheckCreditUsageOpts struct {
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

// CheckCreditUsage returns information about API credit usage and limits
//
// This endpoint returns information about your API credit usage, including daily limits,
// current usage, and remaining credits. This is useful for monitoring API usage and
// ensuring you don't exceed your plan limits.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - *RespCreditUsage: Credit usage information including limits and current usage
//   - error: Error if the request fails
//
// Example:
//
//	// Check API credit usage
//	usage, err := client.CheckCreditUsage(ctx, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Daily Limit: %s\n", usage.DailyLimit)
//	fmt.Printf("Daily Used: %s\n", usage.DailyUsed)
//	fmt.Printf("Daily Remaining: %s\n", usage.DailyRemaining)
//	fmt.Printf("Monthly Limit: %s\n", usage.MonthlyLimit)
//	fmt.Printf("Monthly Used: %s\n", usage.MonthlyUsed)
//	fmt.Printf("Monthly Remaining: %s\n", usage.MonthlyRemaining)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	usage, err := client.CheckCreditUsage(ctx, &CheckCreditUsageOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Returns nil if no usage info found
//   - Useful for monitoring API usage and limits
//   - Helps prevent exceeding plan limits
func (c *HTTPClient) CheckCreditUsage(ctx context.Context, opts *CheckCreditUsageOpts) (*RespCreditUsage, error) {
	// Apply default values from struct tags
	if err := ApplyDefaults(opts); err != nil {
		return nil, err
	}

	params := map[string]string{}

	if opts.ChainID != 0 {
		params["chainid"] = strconv.FormatInt(opts.ChainID, 10)
	}

	// Handle rate limiting
	onLimitExceeded := opts.OnLimitExceeded

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
