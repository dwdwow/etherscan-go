package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetERC20TotalSupply(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	supply, err := config.Client.GetERC20TotalSupply(ctx, TestAddresses.USDTContract, &GetERC20TotalSupplyOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20TotalSupply failed: %v", err)
	}

	if supply == "" {
		t.Error("ERC20 total supply is empty")
	} else {
		t.Logf("USDT total supply: %s", supply)
	}
}

func TestGetERC20AccountBalance(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract and Vitalik's address
	balance, err := config.Client.GetERC20AccountBalance(ctx, TestAddresses.USDTContract, TestAddresses.VitalikButerin, &GetERC20AccountBalanceOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20AccountBalance failed: %v", err)
	}

	if balance == "" {
		t.Error("ERC20 account balance is empty")
	} else {
		t.Logf("USDT balance for %s: %s", TestAddresses.VitalikButerin, balance)
	}
}

func TestGetERC20HistoricalTotalSupply(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract and a recent block
	supply, err := config.Client.GetERC20HistoricalTotalSupply(ctx, TestAddresses.USDTContract, TestBlocks.RecentBlock, &GetERC20HistoricalTotalSupplyOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20HistoricalTotalSupply failed: %v", err)
	}

	if supply == "" {
		t.Error("ERC20 historical total supply is empty")
	} else {
		t.Logf("USDT historical total supply at block %d: %s", TestBlocks.RecentBlock, supply)
	}
}

func TestGetERC20HistoricalAccountBalance(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract, Vitalik's address, and a recent block
	balance, err := config.Client.GetERC20HistoricalAccountBalance(ctx, TestAddresses.USDTContract, TestAddresses.VitalikButerin, TestBlocks.RecentBlock, &GetERC20HistoricalAccountBalanceOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20HistoricalAccountBalance failed: %v", err)
	}

	if balance == "" {
		t.Error("ERC20 historical account balance is empty")
	} else {
		t.Logf("USDT historical balance for %s at block %d: %s", TestAddresses.VitalikButerin, TestBlocks.RecentBlock, balance)
	}
}

func TestGetERC20Holders(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	holders, err := config.Client.GetERC20Holders(ctx, TestAddresses.USDTContract, &GetERC20HoldersOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20Holders failed: %v", err)
	}

	if len(holders) == 0 {
		t.Log("No ERC20 holders found for USDT contract")
	} else {
		t.Logf("Found %d ERC20 holders", len(holders))
		// Validate first holder
		holder := holders[0]
		if holder.TokenHolderAddress == "" {
			t.Error("Holder TokenHolderAddress field is empty")
		}
		if holder.TokenHolderQuantity == "" {
			t.Error("Holder TokenHolderQuantity field is empty")
		}
		t.Logf("First holder: %s, Balance: %s", holder.TokenHolderAddress, holder.TokenHolderQuantity)
	}
}

func TestGetERC20HolderCount(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	count, err := config.Client.GetERC20HolderCount(ctx, TestAddresses.USDTContract, &GetERC20HolderCountOpts{})
	if err != nil {
		t.Fatalf("GetERC20HolderCount failed: %v", err)
	}

	if count == "" {
		t.Error("ERC20 holder count is empty")
	} else {
		t.Logf("USDT holder count: %s", count)
	}
}

func TestGetTopERC20Holders(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	holders, err := config.Client.GetTopERC20Holders(ctx, TestAddresses.USDTContract, 10, &GetTopERC20HoldersOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTopERC20Holders failed: %v", err)
	}

	if len(holders) == 0 {
		t.Log("No top ERC20 holders found for USDT contract")
	} else {
		t.Logf("Found %d top ERC20 holders", len(holders))
		// Validate first holder
		holder := holders[0]
		if holder.TokenHolderAddress == "" {
			t.Error("Holder TokenHolderAddress field is empty")
		}
		if holder.TokenHolderQuantity == "" {
			t.Error("Holder TokenHolderQuantity field is empty")
		}
		t.Logf("Top holder: %s, Balance: %s", holder.TokenHolderAddress, holder.TokenHolderQuantity)
	}
}

func TestGetTokenInfo(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	info, err := config.Client.GetTokenInfo(ctx, TestAddresses.USDTContract, &GetTokenInfoOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTokenInfo failed: %v", err)
	}

	if info == nil {
		t.Error("Token info is nil")
	} else {
		t.Logf("Token info: %+v", info)
		if info.TokenName == "" {
			t.Error("TokenName field is empty")
		}
		if info.Symbol == "" {
			t.Error("Symbol field is empty")
		}
		if info.Divisor == "" {
			t.Error("Divisor field is empty")
		}
	}
}

func TestGetAccountERC20Holdings(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	holdings, err := config.Client.GetAccountERC20Holdings(ctx, TestAddresses.VitalikButerin, &GetAccountERC20HoldingsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetAccountERC20Holdings failed: %v", err)
	}

	if len(holdings) == 0 {
		t.Log("No ERC20 holdings found for test address")
	} else {
		t.Logf("Found %d ERC20 holdings", len(holdings))
		// Validate first holding
		holding := holdings[0]
		if holding.TokenName == "" {
			t.Error("Holding TokenName field is empty")
		}
		if holding.TokenSymbol == "" {
			t.Error("Holding TokenSymbol field is empty")
		}
		if holding.TokenQuantity == "" {
			t.Error("Holding TokenQuantity field is empty")
		}
		t.Logf("First holding: %s (%s), Quantity: %s", holding.TokenName, holding.TokenSymbol, holding.TokenQuantity)
	}
}

func TestGetAccountNFTHoldings(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	holdings, err := config.Client.GetAccountNFTHoldings(ctx, TestAddresses.VitalikButerin, &GetAccountNFTHoldingsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetAccountNFTHoldings failed: %v", err)
	}

	if len(holdings) == 0 {
		t.Log("No NFT holdings found for test address")
	} else {
		t.Logf("Found %d NFT holdings", len(holdings))
		// Validate first holding
		holding := holdings[0]
		if holding.TokenName == "" {
			t.Error("Holding TokenName field is empty")
		}
		if holding.TokenSymbol == "" {
			t.Error("Holding TokenSymbol field is empty")
		}
		if holding.TokenQuantity == "" {
			t.Error("Holding TokenQuantity field is empty")
		}
		t.Logf("First NFT holding: %s (%s), Quantity: %s", holding.TokenName, holding.TokenSymbol, holding.TokenQuantity)
	}
}

func TestGetAccountNFTInventories(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address and a known NFT contract
	// Using a popular NFT contract address (example: Bored Ape Yacht Club)
	nftContract := "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D"
	inventories, err := config.Client.GetAccountNFTInventories(ctx, TestAddresses.VitalikButerin, nftContract, &GetAccountNFTInventoriesOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetAccountNFTInventories failed: %v", err)
	}

	if len(inventories) == 0 {
		t.Log("No NFT inventories found for test address and contract")
	} else {
		t.Logf("Found %d NFT inventories", len(inventories))
		// Validate first inventory
		inventory := inventories[0]
		if inventory.TokenID == "" {
			t.Error("Inventory TokenID field is empty")
		}
		if inventory.TokenAddress == "" {
			t.Error("Inventory TokenAddress field is empty")
		}
		t.Logf("First NFT inventory: TokenID=%s, TokenAddress=%s", inventory.TokenID, inventory.TokenAddress)
	}
}

func TestGetERC20TotalSupplyWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	supply, err := config.Client.GetERC20TotalSupply(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetERC20TotalSupply with nil opts failed: %v", err)
	}

	if supply == "" {
		t.Error("ERC20 total supply is empty")
	} else {
		t.Logf("USDT total supply (with default opts): %s", supply)
	}
}

func TestGetERC20AccountBalanceWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	balance, err := config.Client.GetERC20AccountBalance(ctx, TestAddresses.USDTContract, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetERC20AccountBalance with nil opts failed: %v", err)
	}

	if balance == "" {
		t.Error("ERC20 account balance is empty")
	} else {
		t.Logf("USDT balance (with default opts): %s", balance)
	}
}

func TestGetTokenInfoWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	info, err := config.Client.GetTokenInfo(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetTokenInfo with nil opts failed: %v", err)
	}

	if info == nil {
		t.Error("Token info is nil")
	} else {
		t.Logf("Token info (with default opts): %+v", info)
	}
}

func TestGetERC20TotalSupplyWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	supply, err := config.Client.GetERC20TotalSupply(ctx, TestAddresses.USDTContract, &GetERC20TotalSupplyOpts{
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20TotalSupply with Polygon failed: %v", err)
	}

	if supply == "" {
		t.Error("ERC20 total supply is empty on Polygon")
	} else {
		t.Logf("USDT total supply on Polygon: %s", supply)
	}
}

func TestGetERC20AccountBalanceWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	balance, err := config.Client.GetERC20AccountBalance(ctx, TestAddresses.USDTContract, TestAddresses.VitalikButerin, &GetERC20AccountBalanceOpts{
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20AccountBalance with Polygon failed: %v", err)
	}

	if balance == "" {
		t.Error("ERC20 account balance is empty on Polygon")
	} else {
		t.Logf("USDT balance on Polygon: %s", balance)
	}
}
