package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetDailyBlockCountRewards(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	rewards, err := config.Client.GetDailyBlockCountRewards(ctx, startDate, endDate, &GetDailyBlockCountRewardsOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyBlockCountRewards failed: %v", err)
	}

	if len(rewards) == 0 {
		t.Log("No daily block count rewards found for the date range")
	} else {
		t.Logf("Found %d daily block count rewards", len(rewards))
		// Validate first entry
		reward := rewards[0]
		if reward.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if reward.BlockCount == 0 {
			t.Error("BlockCount field is empty")
		}
		if reward.BlockRewardsEth == "" {
			t.Error("BlockRewardsEth field is empty")
		}
		t.Logf("First entry: Date=%s, BlockCount=%d, Rewards=%s ETH", reward.UTCDate, reward.BlockCount, reward.BlockRewardsEth)
	}
}

func TestGetDailyBlockRewards(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	rewards, err := config.Client.GetDailyBlockRewards(ctx, startDate, endDate, &GetDailyBlockRewardsOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyBlockRewards failed: %v", err)
	}

	if len(rewards) == 0 {
		t.Log("No daily block rewards found for the date range")
	} else {
		t.Logf("Found %d daily block rewards", len(rewards))
		// Validate first entry
		reward := rewards[0]
		if reward.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if reward.BlockRewardsEth == "" {
			t.Error("BlockRewardsEth field is empty")
		}
		t.Logf("First entry: Date=%s, Rewards=%s ETH", reward.UTCDate, reward.BlockRewardsEth)
	}
}

func TestGetDailyAvgBlockTime(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	times, err := config.Client.GetDailyAvgBlockTime(ctx, startDate, endDate, &GetDailyAvgBlockTimeOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAvgBlockTime failed: %v", err)
	}

	if len(times) == 0 {
		t.Log("No daily average block times found for the date range")
	} else {
		t.Logf("Found %d daily average block times", len(times))
		// Validate first entry
		time := times[0]
		if time.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if time.BlockTimeSec == "" {
			t.Error("BlockTimeSec field is empty")
		}
		t.Logf("First entry: Date=%s, BlockTime=%s seconds", time.UTCDate, time.BlockTimeSec)
	}
}

func TestGetDailyUncleBlockCountAndRewards(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	uncles, err := config.Client.GetDailyUncleBlockCountAndRewards(ctx, startDate, endDate, &GetDailyUncleBlockCountAndRewardsOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyUncleBlockCountAndRewards failed: %v", err)
	}

	if len(uncles) == 0 {
		t.Log("No daily uncle block count and rewards found for the date range")
	} else {
		t.Logf("Found %d daily uncle block count and rewards", len(uncles))
		// Validate first entry
		uncle := uncles[0]
		if uncle.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if uncle.UncleBlockCount == 0 {
			t.Error("UncleBlockCount field is empty")
		}
		if uncle.UncleBlockRewardsEth == "" {
			t.Error("UncleBlockRewardsEth field is empty")
		}
		t.Logf("First entry: Date=%s, UncleCount=%d, Rewards=%s ETH", uncle.UTCDate, uncle.UncleBlockCount, uncle.UncleBlockRewardsEth)
	}
}

func TestGetTotalEthSupply(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	supply, err := config.Client.GetTotalEthSupply(ctx, &GetTotalEthSupplyOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTotalEthSupply failed: %v", err)
	}

	if supply == "" {
		t.Error("Total ETH supply is empty")
	} else {
		t.Logf("Total ETH supply: %s", supply)
	}
}

func TestGetTotalEth2Supply(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	supply, err := config.Client.GetTotalEth2Supply(ctx, &GetTotalEth2SupplyOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTotalEth2Supply failed: %v", err)
	}

	if supply == "" {
		t.Error("Total ETH2 supply is empty")
	} else {
		t.Logf("Total ETH2 supply: %s", supply)
	}
}

func TestGetEthPrice(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	price, err := config.Client.GetEthPrice(ctx, &GetEthPriceOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthPrice failed: %v", err)
	}

	if price == nil {
		t.Error("ETH price is nil")
	} else {
		t.Logf("ETH price: %+v", price)
		if price.EthBTC == "" {
			t.Error("EthBTC field is empty")
		}
		if price.EthUSD == "" {
			t.Error("EthUSD field is empty")
		}
	}
}

func TestGetEthHistoricalPrices(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	prices, err := config.Client.GetEthHistoricalPrices(ctx, startDate, endDate, &GetEthHistoricalPricesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthHistoricalPrices failed: %v", err)
	}

	if len(prices) == 0 {
		t.Log("No ETH historical prices found for the date range")
	} else {
		t.Logf("Found %d ETH historical prices", len(prices))
		// Validate first entry
		price := prices[0]
		if price.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if price.Value == "" {
			t.Error("Value field is empty")
		}
		t.Logf("First entry: Date=%s, Value=%s", price.UTCDate, price.Value)
	}
}

func TestGetEthereumNodesSize(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	sizes, err := config.Client.GetEthereumNodesSize(ctx, startDate, endDate, "geth", "default", "asc", &GetEthereumNodesSizeOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthereumNodesSize failed: %v", err)
	}

	if len(sizes) == 0 {
		t.Log("No Ethereum nodes size data found for the date range")
	} else {
		t.Logf("Found %d Ethereum nodes size entries", len(sizes))
		// Validate first entry
		size := sizes[0]
		if size.BlockNumber == "" {
			t.Error("BlockNumber field is empty")
		}
		if size.ChainSize == "" {
			t.Error("ChainSize field is empty")
		}
		t.Logf("First entry: BlockNumber=%s, ChainSize=%s", size.BlockNumber, size.ChainSize)
	}
}

func TestGetNodeCount(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := config.Client.GetNodeCount(ctx, &GetNodeCountOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetNodeCount failed: %v", err)
	}

	if count == nil {
		t.Error("Node count is nil")
	} else {
		t.Logf("Node count: %+v", count)
	}
}

func TestGetDailyTxFees(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	fees, err := config.Client.GetDailyTxFees(ctx, startDate, endDate, &GetDailyTxFeesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyTxFees failed: %v", err)
	}

	if len(fees) == 0 {
		t.Log("No daily transaction fees found for the date range")
	} else {
		t.Logf("Found %d daily transaction fees", len(fees))
		// Validate first entry
		fee := fees[0]
		if fee.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if fee.TransactionFeeEth == "" {
			t.Error("TransactionFeeEth field is empty")
		}
		t.Logf("First entry: Date=%s, TransactionFee=%s ETH", fee.UTCDate, fee.TransactionFeeEth)
	}
}

func TestGetDailyNewAddresses(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	addresses, err := config.Client.GetDailyNewAddresses(ctx, startDate, endDate, &GetDailyNewAddressesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyNewAddresses failed: %v", err)
	}

	if len(addresses) == 0 {
		t.Log("No daily new addresses found for the date range")
	} else {
		t.Logf("Found %d daily new addresses", len(addresses))
		// Validate first entry
		addr := addresses[0]
		if addr.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if addr.NewAddressCount == 0 {
			t.Error("NewAddressCount field is empty")
		}
		t.Logf("First entry: Date=%s, NewAddresses=%d", addr.UTCDate, addr.NewAddressCount)
	}
}

func TestGetDailyNetworkUtilizations(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	utilizations, err := config.Client.GetDailyNetworkUtilizations(ctx, startDate, endDate, &GetDailyNetworkUtilizationsOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyNetworkUtilizations failed: %v", err)
	}

	if len(utilizations) == 0 {
		t.Log("No daily network utilizations found for the date range")
	} else {
		t.Logf("Found %d daily network utilizations", len(utilizations))
		// Validate first entry
		util := utilizations[0]
		if util.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if util.NetworkUtilization == "" {
			t.Error("NetworkUtilization field is empty")
		}
		t.Logf("First entry: Date=%s, NetworkUtilization=%s", util.UTCDate, util.NetworkUtilization)
	}
}

func TestGetDailyAvgHashrates(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	hashrates, err := config.Client.GetDailyAvgHashrates(ctx, startDate, endDate, &GetDailyAvgHashratesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAvgHashrates failed: %v", err)
	}

	if len(hashrates) == 0 {
		t.Log("No daily average hashrates found for the date range")
	} else {
		t.Logf("Found %d daily average hashrates", len(hashrates))
		// Validate first entry
		hashrate := hashrates[0]
		if hashrate.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if hashrate.NetworkHashRate == "" {
			t.Error("NetworkHashRate field is empty")
		}
		t.Logf("First entry: Date=%s, NetworkHashRate=%s", hashrate.UTCDate, hashrate.NetworkHashRate)
	}
}

func TestGetDailyTxCounts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	counts, err := config.Client.GetDailyTxCounts(ctx, startDate, endDate, &GetDailyTxCountsOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyTxCounts failed: %v", err)
	}

	if len(counts) == 0 {
		t.Log("No daily transaction counts found for the date range")
	} else {
		t.Logf("Found %d daily transaction counts", len(counts))
		// Validate first entry
		count := counts[0]
		if count.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if count.TransactionCount == 0 {
			t.Error("TransactionCount field is empty")
		}
		t.Logf("First entry: Date=%s, TransactionCount=%d", count.UTCDate, count.TransactionCount)
	}
}

func TestGetDailyAvgDifficulties(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	difficulties, err := config.Client.GetDailyAvgDifficulties(ctx, startDate, endDate, &GetDailyAvgDifficultiesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAvgDifficulties failed: %v", err)
	}

	if len(difficulties) == 0 {
		t.Log("No daily average difficulties found for the date range")
	} else {
		t.Logf("Found %d daily average difficulties", len(difficulties))
		// Validate first entry
		difficulty := difficulties[0]
		if difficulty.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if difficulty.NetworkDifficulty == "" {
			t.Error("NetworkDifficulty field is empty")
		}
		t.Logf("First entry: Date=%s, NetworkDifficulty=%s", difficulty.UTCDate, difficulty.NetworkDifficulty)
	}
}

func TestGetTotalEthSupplyWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	supply, err := config.Client.GetTotalEthSupply(ctx, nil)
	if err != nil {
		t.Fatalf("GetTotalEthSupply with nil opts failed: %v", err)
	}

	if supply == "" {
		t.Error("Total ETH supply is empty")
	} else {
		t.Logf("Total ETH supply (with default opts): %s", supply)
	}
}

func TestGetEthPriceWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	price, err := config.Client.GetEthPrice(ctx, nil)
	if err != nil {
		t.Fatalf("GetEthPrice with nil opts failed: %v", err)
	}

	if price == nil {
		t.Error("ETH price is nil")
	} else {
		t.Logf("ETH price (with default opts): %+v", price)
	}
}
