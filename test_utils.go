package etherscan

import (
	"os"
	"testing"
)

// TestConfig holds test configuration
type TestConfig struct {
	APIKey string
	Client *HTTPClient
}

// GetTestConfig returns test configuration with API key from environment
func GetTestConfig(t *testing.T) *TestConfig {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping test: ETHERSCAN_API_KEY environment variable not set")
	}

	client := NewHTTPClient(HTTPClientConfig{
		APIKey:  apiKey,
		APITier: StandardTier,
	})

	return &TestConfig{
		APIKey: apiKey,
		Client: client,
	}
}

// TestAddresses contains commonly used test addresses
var TestAddresses = struct {
	VitalikButerin string
	USDTContract   string
	USDCContract   string
	WETHContract   string
}{
	VitalikButerin: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", // Vitalik's address
	USDTContract:   "0xdAC17F958D2ee523a2206206994597C13D831ec7", // USDT contract on Ethereum
	USDCContract:   "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC contract on Ethereum
	WETHContract:   "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH contract on Ethereum
}

// TestTransactions contains commonly used test transaction hashes
var TestTransactions = struct {
	SampleTxHash string
}{
	SampleTxHash: "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44",
}

// TestBlocks contains commonly used test block numbers
var TestBlocks = struct {
	GenesisBlock int64
	RecentBlock  int64
}{
	GenesisBlock: 0,
	RecentBlock:  18500000, // A recent block number
}
