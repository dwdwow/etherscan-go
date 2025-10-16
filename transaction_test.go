package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetNormalTxs(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	transactions, err := config.Client.GetNormalTxs(ctx, TestAddresses.VitalikButerin, &GetNormalTxsOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTxs failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found for test address")
	} else {
		t.Logf("Found %d normal transactions", len(transactions))
		// Validate first transaction
		tx := transactions[0]
		if tx.Hash == "" {
			t.Error("Transaction Hash field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
		if tx.To == "" {
			t.Log("Transaction To field is empty (this is normal for contract creation)")
		}
		if tx.Value == "" {
			t.Error("Transaction Value field is empty")
		}
	}
}

func TestGetInternalTxsByAddress(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	transactions, err := config.Client.GetInternalTxsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTxsByAddressOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTxsByAddress failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test address")
	} else {
		t.Logf("Found %d internal transactions", len(transactions))
		// Validate first transaction
		tx := transactions[0]
		if tx.Hash == "" {
			t.Error("Transaction Hash field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
		if tx.To == "" {
			t.Error("Transaction To field is empty")
		}
		if tx.Value == "" {
			t.Error("Transaction Value field is empty")
		}
	}
}

func TestGetInternalTxsByHash(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	transactions, err := config.Client.GetInternalTxsByHash(ctx, TestTransactions.SampleTxHash, &GetInternalTxsByHashOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTxsByHash failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test transaction hash")
	} else {
		t.Logf("Found %d internal transactions", len(transactions))
		// Validate first transaction
		tx := transactions[0]
		if tx.BlockNumber == "" {
			t.Error("Transaction BlockNumber field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
		if tx.To == "" {
			t.Error("Transaction To field is empty")
		}
	}
}

func TestGetInternalTxsByBlockRange(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a recent block range
	transactions, err := config.Client.GetInternalTxsByBlockRange(ctx, int(TestBlocks.RecentBlock), int(TestBlocks.RecentBlock+10), &GetInternalTxsByBlockRangeOpts{
		Page:    1,
		Offset:  10,
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTxsByBlockRange failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test block range")
	} else {
		t.Logf("Found %d internal transactions", len(transactions))
		// Validate first transaction
		tx := transactions[0]
		if tx.Hash == "" {
			t.Error("Transaction Hash field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
		if tx.To == "" {
			t.Error("Transaction To field is empty")
		}
	}
}

func TestGetContractExecutionStatus(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	status, err := config.Client.GetContractExecutionStatus(ctx, TestTransactions.SampleTxHash, &GetContractExecutionStatusOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetContractExecutionStatus failed: %v", err)
	}

	if status == nil {
		t.Log("Contract execution status not found (this is normal if the transaction doesn't exist)")
	} else {
		t.Logf("Contract execution status: %+v", status)
		if status.IsError == "" {
			t.Error("IsError field is empty")
		}
		// if status.ErrDescription == "" {
		// 	t.Log("ErrDescription field is empty (this is normal for successful transactions)")
		// }
	}
}

func TestGetTxReceiptStatus(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	status, err := config.Client.GetTxReceiptStatus(ctx, TestTransactions.SampleTxHash, &GetTxReceiptStatusOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTxReceiptStatus failed: %v", err)
	}

	if status == nil {
		t.Log("Transaction receipt status not found (this is normal if the transaction doesn't exist)")
	} else {
		t.Logf("Transaction receipt status: %+v", status)
		if status.Status == "" {
			t.Error("Status field is empty")
		}
	}
}

func TestGetNormalTxsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetNormalTxs(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetNormalTxs with nil opts failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d normal transactions (with default opts)", len(transactions))
	}
}

func TestGetInternalTxsByAddressWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetInternalTxsByAddress(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetInternalTxsByAddress with nil opts failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d internal transactions (with default opts)", len(transactions))
	}
}

func TestGetInternalTxsByHashWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetInternalTxsByHash(ctx, TestTransactions.SampleTxHash, nil)
	if err != nil {
		t.Fatalf("GetInternalTxsByHash with nil opts failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test transaction hash (with default opts)")
	} else {
		t.Logf("Found %d internal transactions (with default opts)", len(transactions))
	}
}

func TestGetContractExecutionStatusWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	status, err := config.Client.GetContractExecutionStatus(ctx, TestTransactions.SampleTxHash, nil)
	if err != nil {
		t.Fatalf("GetContractExecutionStatus with nil opts failed: %v", err)
	}

	if status == nil {
		t.Log("Contract execution status not found (with default opts)")
	} else {
		t.Logf("Contract execution status (with default opts): %+v", status)
	}
}

func TestGetTxReceiptStatusWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	status, err := config.Client.GetTxReceiptStatus(ctx, TestTransactions.SampleTxHash, nil)
	if err != nil {
		t.Fatalf("GetTxReceiptStatus with nil opts failed: %v", err)
	}

	if status == nil {
		t.Log("Transaction receipt status not found (with default opts)")
	} else {
		t.Logf("Transaction receipt status (with default opts): %+v", status)
	}
}

func TestGetNormalTxsWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	transactions, err := config.Client.GetNormalTxs(ctx, TestAddresses.VitalikButerin, &GetNormalTxsOpts{
		StartBlock: 40000000, // Recent block on Polygon
		EndBlock:   40000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTxs with Polygon failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d normal transactions on Polygon", len(transactions))
	}
}

func TestGetInternalTxsByAddressWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	transactions, err := config.Client.GetInternalTxsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTxsByAddressOpts{
		StartBlock: 40000000, // Recent block on Polygon
		EndBlock:   40000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTxsByAddress with Polygon failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d internal transactions on Polygon", len(transactions))
	}
}

func TestGetNormalTxsWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	transactions, err := config.Client.GetNormalTxs(ctx, TestAddresses.VitalikButerin, &GetNormalTxsOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "desc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTxs with desc sort failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found for test address (desc sort)")
	} else {
		t.Logf("Found %d normal transactions (desc sort)", len(transactions))
	}
}

func TestGetInternalTxsByAddressWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	transactions, err := config.Client.GetInternalTxsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTxsByAddressOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "desc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTxsByAddress with desc sort failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test address (desc sort)")
	} else {
		t.Logf("Found %d internal transactions (desc sort)", len(transactions))
	}
}
