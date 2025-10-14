package etherscan

import (
	"context"
	"strconv"
	"testing"
	"time"
)

func TestGetBlockAndUncleRewards(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a recent block number
	rewards, err := config.Client.GetBlockAndUncleRewards(ctx, TestBlocks.RecentBlock, &GetBlockAndUncleRewardsOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlockAndUncleRewards failed: %v", err)
	}

	if rewards == nil {
		t.Error("Block rewards is nil")
	} else {
		t.Logf("Block rewards: %+v", rewards)
		if rewards.BlockNumber == "" {
			t.Error("BlockNumber field is empty")
		}
		if rewards.BlockReward == "" {
			t.Error("BlockReward field is empty")
		}
	}
}

func TestGetBlockTransactionsCount(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a recent block number
	count, err := config.Client.GetBlockTransactionsCount(ctx, TestBlocks.RecentBlock, &GetBlockTransactionsCountOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlockTransactionsCount failed: %v", err)
	}

	if count == nil {
		t.Error("Block transactions count is nil")
	} else {
		t.Logf("Block transactions count: %+v", count)
		if count.Block == 0 {
			t.Error("Block field is empty")
		}
		if count.TxsCount == 0 {
			t.Log("Block has no transactions (this is normal for some blocks)")
		}
	}
}

func TestGetBlockCountdownTime(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Use a known block number instead of calling RpcEthBlockNumber
	blockNo := "0x1a2b3c" // Use a hardcoded block number for testing

	// Use a known block number instead of calling RpcEthBlockNumber
	blockNumber := blockNo // Example hex block number

	if blockNumber == "" {
		t.Error("Block number is empty")
	} else {
		t.Logf("Block number: %s", blockNumber)
	}

	// Convert hex block number to int64
	currentBlock, err := strconv.ParseInt(blockNumber[2:], 16, 64)
	if err != nil {
		t.Fatalf("Failed to parse block number: %v", err)
	}

	// Use current block + 1000 as future block
	futureBlock := currentBlock + 1000

	countdown, err := config.Client.GetBlockCountdownTime(ctx, futureBlock, &GetBlockCountdownTimeOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlockCountdownTime failed: %v", err)
	}

	if countdown == nil {
		t.Error("Block countdown is nil")
	} else {
		t.Logf("Block countdown: %+v", countdown)
		if countdown.CurrentBlock == "" {
			t.Error("CurrentBlock field is empty")
		}
		if countdown.CountdownBlock == "" {
			t.Error("CountdownBlock field is empty")
		}
	}
}

func TestGetBlockNumberByTimestamp(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a timestamp from 2023
	timestamp := int64(1672531200) // 2023-01-01 00:00:00 UTC
	blockNumber, err := config.Client.GetBlockNumberByTimestamp(ctx, timestamp, "before", &GetBlockNumberByTimestampOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlockNumberByTimestamp failed: %v", err)
	}

	if blockNumber == 0 {
		t.Error("Block number is empty")
	} else {
		t.Logf("Block number for timestamp %d: %d", timestamp, blockNumber)
	}
}

func TestGetBlockNumberByTimestampAfter(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with "after" parameter
	timestamp := int64(1672531200) // 2023-01-01 00:00:00 UTC
	blockNumber, err := config.Client.GetBlockNumberByTimestamp(ctx, timestamp, "after", &GetBlockNumberByTimestampOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlockNumberByTimestamp with 'after' failed: %v", err)
	}

	if blockNumber == 0 {
		t.Error("Block number is empty")
	} else {
		t.Logf("Block number after timestamp %d: %d", timestamp, blockNumber)
	}
}

func TestGetDailyAvgBlockSizes(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	sizes, err := config.Client.GetDailyAvgBlockSizes(ctx, startDate, endDate, &GetDailyAvgBlockSizesOpts{
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAvgBlockSizes failed: %v", err)
	}

	if len(sizes) == 0 {
		t.Log("No daily average block sizes found for the date range")
	} else {
		t.Logf("Found %d daily average block sizes", len(sizes))
		// Validate first entry
		size := sizes[0]
		if size.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if size.BlockSizeBytes == 0 {
			t.Error("BlockSizeBytes field is empty")
		}
		t.Logf("First entry: Date=%s, Size=%d bytes", size.UTCDate, size.BlockSizeBytes)
	}
}

func TestGetDailyAvgBlockSizesWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	sizes, err := config.Client.GetDailyAvgBlockSizes(ctx, startDate, endDate, &GetDailyAvgBlockSizesOpts{
		Sort:    "desc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAvgBlockSizes with desc sort failed: %v", err)
	}

	if len(sizes) == 0 {
		t.Log("No daily average block sizes found for the date range")
	} else {
		t.Logf("Found %d daily average block sizes (desc order)", len(sizes))
	}
}

func TestGetBlockAndUncleRewardsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	rewards, err := config.Client.GetBlockAndUncleRewards(ctx, TestBlocks.RecentBlock, nil)
	if err != nil {
		t.Fatalf("GetBlockAndUncleRewards with nil opts failed: %v", err)
	}

	if rewards == nil {
		t.Error("Block rewards is nil")
	} else {
		t.Logf("Block rewards with default opts: %+v", rewards)
	}
}

func TestGetBlockTransactionsCountWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	count, err := config.Client.GetBlockTransactionsCount(ctx, TestBlocks.RecentBlock, nil)
	if err != nil {
		t.Fatalf("GetBlockTransactionsCount with nil opts failed: %v", err)
	}

	if count == nil {
		t.Error("Block transactions count is nil")
	} else {
		t.Logf("Block transactions count with default opts: %+v", count)
	}
}
