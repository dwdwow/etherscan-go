package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetNormalTransactions(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	transactions, err := config.Client.GetNormalTransactions(ctx, TestAddresses.VitalikButerin, &GetNormalTransactionsOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTransactions failed: %v", err)
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

func TestGetInternalTransactionsByAddress(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	transactions, err := config.Client.GetInternalTransactionsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTransactionsByAddressOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTransactionsByAddress failed: %v", err)
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

func TestGetInternalTransactionsByHash(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	transactions, err := config.Client.GetInternalTransactionsByHash(ctx, TestTransactions.SampleTxHash, &GetInternalTransactionsByHashOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTransactionsByHash failed: %v", err)
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

func TestGetInternalTransactionsByBlockRange(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a recent block range
	transactions, err := config.Client.GetInternalTransactionsByBlockRange(ctx, int(TestBlocks.RecentBlock), int(TestBlocks.RecentBlock+10), &GetInternalTransactionsByBlockRangeOpts{
		Page:    1,
		Offset:  10,
		Sort:    "asc",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTransactionsByBlockRange failed: %v", err)
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
		if status.ErrDescription == "" {
			t.Log("ErrDescription field is empty (this is normal for successful transactions)")
		}
	}
}

func TestGetTransactionReceiptStatus(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	status, err := config.Client.GetTransactionReceiptStatus(ctx, TestTransactions.SampleTxHash, &GetTransactionReceiptStatusOpts{
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetTransactionReceiptStatus failed: %v", err)
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

func TestGetNormalTransactionsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetNormalTransactions(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetNormalTransactions with nil opts failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d normal transactions (with default opts)", len(transactions))
	}
}

func TestGetInternalTransactionsByAddressWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetInternalTransactionsByAddress(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetInternalTransactionsByAddress with nil opts failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d internal transactions (with default opts)", len(transactions))
	}
}

func TestGetInternalTransactionsByHashWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	transactions, err := config.Client.GetInternalTransactionsByHash(ctx, TestTransactions.SampleTxHash, nil)
	if err != nil {
		t.Fatalf("GetInternalTransactionsByHash with nil opts failed: %v", err)
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

func TestGetTransactionReceiptStatusWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	status, err := config.Client.GetTransactionReceiptStatus(ctx, TestTransactions.SampleTxHash, nil)
	if err != nil {
		t.Fatalf("GetTransactionReceiptStatus with nil opts failed: %v", err)
	}

	if status == nil {
		t.Log("Transaction receipt status not found (with default opts)")
	} else {
		t.Logf("Transaction receipt status (with default opts): %+v", status)
	}
}

func TestGetNormalTransactionsWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	transactions, err := config.Client.GetNormalTransactions(ctx, TestAddresses.VitalikButerin, &GetNormalTransactionsOpts{
		StartBlock: 40000000, // Recent block on Polygon
		EndBlock:   40000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTransactions with Polygon failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d normal transactions on Polygon", len(transactions))
	}
}

func TestGetInternalTransactionsByAddressWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	transactions, err := config.Client.GetInternalTransactionsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTransactionsByAddressOpts{
		StartBlock: 40000000, // Recent block on Polygon
		EndBlock:   40000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "asc",
		ChainID:    137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTransactionsByAddress with Polygon failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d internal transactions on Polygon", len(transactions))
	}
}

func TestGetNormalTransactionsWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	transactions, err := config.Client.GetNormalTransactions(ctx, TestAddresses.VitalikButerin, &GetNormalTransactionsOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "desc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetNormalTransactions with desc sort failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No normal transactions found for test address (desc sort)")
	} else {
		t.Logf("Found %d normal transactions (desc sort)", len(transactions))
	}
}

func TestGetInternalTransactionsByAddressWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	transactions, err := config.Client.GetInternalTransactionsByAddress(ctx, TestAddresses.VitalikButerin, &GetInternalTransactionsByAddressOpts{
		StartBlock: 18000000, // Recent block
		EndBlock:   18000100, // Small range
		Page:       1,
		Offset:     10,
		Sort:       "desc",
		ChainID:    1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetInternalTransactionsByAddress with desc sort failed: %v", err)
	}

	if len(transactions) == 0 {
		t.Log("No internal transactions found for test address (desc sort)")
	} else {
		t.Logf("Found %d internal transactions (desc sort)", len(transactions))
	}
}
