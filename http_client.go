package etherscan

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURL      = "https://api.etherscan.io/v2/api"
	APIAshx      = "https://api-metadata.etherscan.io/v1/api.ashx"
	ChainListURL = "https://api.etherscan.io/v2/chainlist"
)

// Chain IDs
const (
	EthereumMainnet         = 1
	SepoliaTestnet          = 11155111
	HoleskyTestnet          = 17000
	HoodiTestnet            = 560048
	AbstractMainnet         = 2741
	AbstractSepoliaTestnet  = 11124
	ApechainCurtisTestnet   = 33111
	ApechainMainnet         = 33139
	ArbitrumNovaMainnet     = 42170
	ArbitrumOneMainnet      = 42161
	ArbitrumSepoliaTestnet  = 421614
	AvalancheCChain         = 43114
	AvalancheFujiTestnet    = 43113
	BaseMainnet             = 8453
	BaseSepoliaTestnet      = 84532
	BerachainMainnet        = 80094
	BerachainBepoliaTestnet = 80069
	BittorrentChainMainnet  = 199
	BittorrentChainTestnet  = 1029
	BlastMainnet            = 81457
	BlastSepoliaTestnet     = 168587773
	BNBSmartChainMainnet    = 56
	BNBSmartChainTestnet    = 97
	CeloAlfajoresTestnet    = 44787
	CeloMainnet             = 42220
	CronosMainnet           = 25
	FraxtalMainnet          = 252
	FraxtalTestnet          = 2522
	Gnosis                  = 100
	HyperEVMMainnet         = 999
	LineaMainnet            = 59144
	LineaSepoliaTestnet     = 59141
	MantleMainnet           = 5000
	MantleSepoliaTestnet    = 5003
	MemecoRETtestnet        = 43521
	MoonbaseAlphaTestnet    = 1287
	MonadTestnet            = 10143
	MoonbeamMainnet         = 1284
	MoonriverMainnet        = 1285
	OPMainnet               = 10
	OPSepoliaTestnet        = 11155420
	PolygonMainnet          = 137
	PolygonAmoyTestnet      = 80002
	KatanaMainnet           = 747474
	KatanaBokutoTestnet     = 737373
	SeiMainnet              = 1329
	SeiTestnet              = 1328
	ScrollMainnet           = 534352
	ScrollSepoliaTestnet    = 534351
	SonicTestnet            = 14601
	SonicMainnet            = 146
	SophonMainnet           = 50104
	SophonSepoliaTestnet    = 531050104
	SwellchainMainnet       = 1923
	SwellchainTestnet       = 1924
	TaikoMainnet            = 167000
	TaikoHoodiTestnet       = 167012
	UnichainMainnet         = 130
	UnichainSepoliaTestnet  = 1301
	WorldMainnet            = 480
	WorldSepoliaTestnet     = 4801
	XDCApothemTestnet       = 51
	XDCMainnet              = 50
	ZKSyncMainnet           = 324
	ZKSyncSepoliaTestnet    = 300
	OpBNBMainnet            = 204
	OpBNBTestnet            = 5611
)

// API Tiers
const (
	FreeTier         = "free"
	StandardTier     = "standard"
	AdvancedTier     = "advanced"
	ProfessionalTier = "professional"
	ProPlusTier      = "pro_plus"
)

// API rate limits by tier (calls/second)
const (
	FreeTierRateLimit         = 5
	StandardTierRateLimit     = 10
	AdvancedTierRateLimit     = 20
	ProfessionalTierRateLimit = 30
	ProPlusTierRateLimit      = 30
)

// API daily call limits by tier
const (
	FreeTierDailyLimit         = 100_000
	StandardTierDailyLimit     = 200_000
	AdvancedTierDailyLimit     = 500_000
	ProfessionalTierDailyLimit = 1_000_000
	ProPlusTierDailyLimit      = 1_500_000
)

// HTTPClient is a client for the Etherscan V2 API
type HTTPClient struct {
	apiKey          string
	defaultChainID  int
	rateLimiter     *MultiRateLimiter
	onLimitExceeded RateLimitBehavior
	httpClient      *http.Client
}

// HTTPClientConfig represents configuration for HTTPClient
type HTTPClientConfig struct {
	// APIKey is the Etherscan API key (required)
	APIKey string

	// DefaultChainID is the default chain ID to use when not specified in requests
	// Default: EthereumMainnet (1)
	DefaultChainID int

	// APITier specifies the API tier for rate limiting
	// Options: FreeTier, StandardTier, AdvancedTier, ProfessionalTier, ProPlusTier
	// Default: ProPlusTier
	APITier string

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// Default: RateLimitBlock
	OnLimitExceeded RateLimitBehavior

	// HTTPClient allows using a custom HTTP client
	// Default: &http.Client{Timeout: 30 * time.Second}
	HTTPClient *http.Client
}

// NewHTTPClient creates a new Etherscan HTTP client
//
// Example:
//
//	client := NewHTTPClient(HTTPClientConfig{
//	    APIKey: "YOUR_API_KEY",
//	    DefaultChainID: EthereumMainnet,
//	    APITier: ProPlusTier,
//	})
func NewHTTPClient(config HTTPClientConfig) *HTTPClient {
	if config.DefaultChainID == 0 {
		config.DefaultChainID = EthereumMainnet
	}

	if config.APITier == "" {
		config.APITier = ProPlusTier
	}

	if config.OnLimitExceeded == "" {
		config.OnLimitExceeded = RateLimitBlock
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	// Setup rate limiters based on API tier
	var rateLimits []RateLimit
	switch config.APITier {
	case StandardTier:
		rateLimits = []RateLimit{
			{Limit: StandardTierRateLimit, Period: 1 * time.Second},
			{Limit: StandardTierDailyLimit, Period: 24 * time.Hour},
		}
	case AdvancedTier:
		rateLimits = []RateLimit{
			{Limit: AdvancedTierRateLimit, Period: 1 * time.Second},
			{Limit: AdvancedTierDailyLimit, Period: 24 * time.Hour},
		}
	case ProfessionalTier:
		rateLimits = []RateLimit{
			{Limit: ProfessionalTierRateLimit, Period: 1 * time.Second},
			{Limit: ProfessionalTierDailyLimit, Period: 24 * time.Hour},
		}
	case ProPlusTier:
		rateLimits = []RateLimit{
			{Limit: ProPlusTierRateLimit, Period: 1 * time.Second},
			{Limit: ProPlusTierDailyLimit, Period: 24 * time.Hour},
		}
	default: // FreeTier
		rateLimits = []RateLimit{
			{Limit: FreeTierRateLimit, Period: 1 * time.Second},
			{Limit: FreeTierDailyLimit, Period: 24 * time.Hour},
		}
	}

	limiter, err := NewMultiRateLimiter(rateLimits, config.OnLimitExceeded)
	if err != nil {
		// should never happen
		panic(err)
	}

	return &HTTPClient{
		apiKey:          config.APIKey,
		defaultChainID:  config.DefaultChainID,
		rateLimiter:     limiter,
		onLimitExceeded: config.OnLimitExceeded,
		httpClient:      config.HTTPClient,
	}
}

// requestParams contains parameters for internal request method
type requestParams struct {
	ctx             context.Context
	module          string
	action          string
	params          map[string]string
	method          string // "GET" or "POST"
	noFoundReturn   any
	baseURL         string
	onLimitExceeded RateLimitBehavior
	retryCount      int // Track retry attempts for rate limiting
}

// request is the internal method for making API requests
func (c *HTTPClient) request(params requestParams) (any, error) {
	// Use background context if not provided
	if params.ctx == nil {
		params.ctx = context.Background()
	}

	// Determine rate limit behavior
	behavior := c.onLimitExceeded
	if params.onLimitExceeded != "" {
		behavior = params.onLimitExceeded
	}

	// Acquire rate limit token
	acquired, err := c.rateLimiter.Acquire(params.ctx, 1, &behavior)
	if err != nil {
		return nil, err
	}
	if !acquired {
		return nil, errors.New("rate limit exceeded")
	}

	// Set default chain ID if not provided
	if params.params == nil {
		params.params = make(map[string]string)
	}
	if _, ok := params.params["chainid"]; !ok {
		params.params["chainid"] = strconv.Itoa(c.defaultChainID)
	}

	// Remove nil/empty values
	for k, v := range params.params {
		if v == "" {
			delete(params.params, k)
		}
	}

	// Set defaults
	if params.method == "" {
		params.method = "GET"
	}
	if params.baseURL == "" {
		params.baseURL = BaseURL
	}

	// Build request
	var req *http.Request

	switch params.method {
	case "GET":
		// Build query parameters
		queryParams := url.Values{}
		queryParams.Set("module", params.module)
		queryParams.Set("action", params.action)
		queryParams.Set("apikey", c.apiKey)
		for k, v := range params.params {
			queryParams.Set(k, v)
		}

		uri := fmt.Sprintf("%s?%s", params.baseURL, queryParams.Encode())
		req, err = http.NewRequestWithContext(params.ctx, "GET", uri, nil)
	case "POST":
		// Build URL with basic params
		queryParams := url.Values{}
		if chainID, ok := params.params["chainid"]; ok {
			queryParams.Set("chainid", chainID)
			delete(params.params, "chainid")
		}
		queryParams.Set("module", params.module)
		queryParams.Set("action", params.action)
		queryParams.Set("apikey", c.apiKey)

		uri := fmt.Sprintf("%s?%s", params.baseURL, queryParams.Encode())

		// Send remaining params as JSON body
		jsonData, _ := json.Marshal(params.params)
		req, err = http.NewRequestWithContext(params.ctx, "POST", uri, bytes.NewBuffer(jsonData))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if err != nil {
		return nil, err
	}

	// Execute request with retries
	var resp *http.Response
	retryTimes := 3
	for i := 0; i < retryTimes; i++ {
		resp, err = c.httpClient.Do(req)
		if err == nil {
			break
		}

		log.Printf("etherscan: request %s %s failed: %v, retrying %d of %d...", params.module, params.action, err, i+1, retryTimes)
		if i < retryTimes-1 {
			time.Sleep(1 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("etherscan: request %s %s failed after retries: %w", params.module, params.action, err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("etherscan: read response body failed: %w", err)
	}

	// Parse JSON response
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("etherscan: parse %s %s response failed: %v", params.module, params.action, err)
		return params.noFoundReturn, nil
	}

	// Check status
	status := ""
	message := ""
	var data any

	if statusVal, ok := result["status"]; ok {
		status = fmt.Sprintf("%v", statusVal)
	}
	if msgVal, ok := result["message"]; ok {
		message = fmt.Sprintf("%v", msgVal)
	}

	// Check if it's a JSON-RPC response
	if _, hasJSONRPC := result["jsonrpc"]; hasJSONRPC {
		data = result
	} else {
		data = result["result"]
	}

	// Handle HTTP errors
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("etherscan: %s %s failed: %d %s %s %v", params.module, params.action, resp.StatusCode, status, message, data)
	}

	// Handle API errors
	if status == "0" {
		// Check for "No X found" messages
		if strings.HasPrefix(message, "No ") && strings.HasSuffix(message, " found") {
			return params.noFoundReturn, nil
		}

		// Check for rate limit errors
		if strings.Contains(message, "Maximum rate limit reached") ||
			strings.Contains(message, "rate limit") ||
			strings.Contains(message, "Rate limit") {
			// Retry with 1 second delay
			log.Printf("etherscan: rate limit detected for %s %s, retrying in 1 second...", params.module, params.action)
			time.Sleep(1 * time.Second)

			// Recursively retry the request (with a limit to prevent infinite recursion)
			if params.retryCount < 3 {
				params.retryCount++
				return c.request(params)
			}
		}

		return nil, fmt.Errorf("etherscan: %s %s failed: %d %s %s %v", params.module, params.action, resp.StatusCode, status, message, data)
	}

	return data, nil
}

// unmarshalResponse unmarshals the API response into the target type
func unmarshalResponse(data any, target any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, target)
}

// GetSupportedChains returns the list of supported blockchain networks
//
// Example:
//
//	chains, err := client.GetSupportedChains(ctx)
func (c *HTTPClient) GetSupportedChains(ctx context.Context) (*RespSupportedChains, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", ChainListURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RespSupportedChains
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
