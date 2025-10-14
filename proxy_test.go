package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestRpcEthBlockNumber(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	blockNumber, err := config.Client.RpcEthBlockNumber(ctx, &RpcEthBlockNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthBlockNumber failed: %v", err)
	}

	if blockNumber == "" {
		t.Error("Block number is empty")
	} else {
		t.Logf("Current block number: %s", blockNumber)
		// Should be a hex string
		if blockNumber[:2] != "0x" {
			t.Error("Block number should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthBlockByNumber(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with latest block
	block, err := config.Client.RpcEthBlockByNumber(ctx, "latest", &RpcEthBlockByNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthBlockByNumber failed: %v", err)
	}

	if block == nil {
		t.Error("Block is nil")
	} else {
		t.Logf("Block: %+v", block)
		if block.Number == "" {
			t.Error("Block Number field is empty")
		}
		if block.Hash == "" {
			t.Error("Block Hash field is empty")
		}
		if block.ParentHash == "" {
			t.Error("Block ParentHash field is empty")
		}
	}
}

func TestRpcEthBlockByNumberWithFullTxs(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with latest block, returning only transaction hashes
	block, err := config.Client.RpcEthBlockByNumberWithFullTxs(ctx, "latest", &RpcEthBlockByNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthBlockByNumber with transaction hashes failed: %v", err)
	}

	if block == nil {
		t.Error("Block is nil")
	} else {
		t.Logf("Block with transaction hashes: %+v", block)
		if block.Number == "" {
			t.Error("Block Number field is empty")
		}
		if block.Hash == "" {
			t.Error("Block Hash field is empty")
		}
	}
}

func TestRpcEthUncleByBlockNumberAndIndex(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a recent block that might have uncles
	block, err := config.Client.RpcEthUncleByBlockNumberAndIndex(ctx, "latest", "0x0", &RpcEthUncleByBlockNumberAndIndexOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthUncleByBlockNumberAndIndex failed: %v", err)
	}

	if block == nil {
		t.Log("No uncle block found (this is normal for recent blocks)")
	} else {
		t.Logf("Uncle block: %+v", block)
		if block.Number == "" {
			t.Error("Uncle Block Number field is empty")
		}
		if block.Hash == "" {
			t.Error("Uncle Block Hash field is empty")
		}
	}
}

func TestRpcEthBlockTransactionCountByNumber(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := config.Client.RpcEthBlockTransactionCountByNumber(ctx, "latest", &RpcEthBlockTransactionCountByNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthBlockTransactionCountByNumber failed: %v", err)
	}

	if count == "" {
		t.Error("Transaction count is empty")
	} else {
		t.Logf("Transaction count: %s", count)
		// Should be a hex string
		if count[:2] != "0x" {
			t.Error("Transaction count should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthTransactionByHash(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	tx, err := config.Client.RpcEthTransactionByHash(ctx, TestTransactions.SampleTxHash, &RpcEthTransactionByHashOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthTransactionByHash failed: %v", err)
	}

	if tx == nil {
		t.Log("Transaction not found (this is normal if the hash doesn't exist)")
	} else {
		t.Logf("Transaction: %+v", tx)
		if tx.Hash == "" {
			t.Error("Transaction Hash field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
		if tx.To == "" {
			t.Log("Transaction To field is empty (this is normal for contract creation)")
		}
	}
}

func TestRpcEthTransactionByBlockNumberAndIndex(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with latest block and first transaction
	tx, err := config.Client.RpcEthTransactionByBlockNumberAndIndex(ctx, "latest", "0x0", &RpcEthTransactionByBlockNumberAndIndexOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthTransactionByBlockNumberAndIndex failed: %v", err)
	}

	if tx == nil {
		t.Log("Transaction not found (this is normal if the block has no transactions)")
	} else {
		t.Logf("Transaction: %+v", tx)
		if tx.Hash == "" {
			t.Error("Transaction Hash field is empty")
		}
		if tx.From == "" {
			t.Error("Transaction From field is empty")
		}
	}
}

func TestRpcEthTransactionCount(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := config.Client.RpcEthTransactionCount(ctx, TestAddresses.VitalikButerin, "latest", &RpcEthTransactionCountOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthTransactionCount failed: %v", err)
	}

	if count == "" {
		t.Error("Transaction count is empty")
	} else {
		t.Logf("Transaction count for %s: %s", TestAddresses.VitalikButerin, count)
		// Should be a hex string
		if count[:2] != "0x" {
			t.Error("Transaction count should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthTransactionReceipt(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known transaction hash
	receipt, err := config.Client.RpcEthTransactionReceipt(ctx, TestTransactions.SampleTxHash, &RpcEthTransactionReceiptOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthTransactionReceipt failed: %v", err)
	}

	if receipt == nil {
		t.Log("Transaction receipt not found (this is normal if the hash doesn't exist)")
	} else {
		t.Logf("Transaction receipt: %+v", receipt)
		if receipt.TransactionHash == "" {
			t.Error("Receipt TransactionHash field is empty")
		}
		if receipt.BlockNumber == "" {
			t.Error("Receipt BlockNumber field is empty")
		}
		if receipt.GasUsed == "" {
			t.Error("Receipt GasUsed field is empty")
		}
	}
}

func TestRpcEthCall(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract (balanceOf function call)
	// balanceOf(address) function signature: 0x70a08231
	// Vitalik's address as parameter
	callData := "0x70a08231000000000000000000000000" + TestAddresses.VitalikButerin[2:] // Remove 0x prefix
	result, err := config.Client.RpcEthCall(ctx, TestAddresses.USDTContract, callData, &RpcEthCallOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthCall failed: %v", err)
	}

	if result == "" {
		t.Error("Call result is empty")
	} else {
		t.Logf("Call result: %s", result)
		// Should be a hex string
		if result[:2] != "0x" {
			t.Error("Call result should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthGetCode(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract
	code, err := config.Client.RpcEthGetCode(ctx, TestAddresses.USDTContract, &RpcEthGetCodeOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthGetCode failed: %v", err)
	}

	if code == "" {
		t.Error("Contract code is empty")
	} else {
		t.Logf("Contract code length: %d characters", len(code))
		// Should be a hex string
		if code[:2] != "0x" {
			t.Error("Contract code should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthGetStorageAt(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract and storage slot 0
	storage, err := config.Client.RpcEthGetStorageAt(ctx, TestAddresses.USDTContract, "0x0", &RpcEthGetStorageAtOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthGetStorageAt failed: %v", err)
	}

	if storage == "" {
		t.Error("Storage value is empty")
	} else {
		t.Logf("Storage value: %s", storage)
		// Should be a hex string
		if storage[:2] != "0x" {
			t.Error("Storage value should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthGetGasPrice(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gasPrice, err := config.Client.RpcEthGetGasPrice(ctx, &RpcEthGetGasPriceOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthGetGasPrice failed: %v", err)
	}

	if gasPrice == "" {
		t.Error("Gas price is empty")
	} else {
		t.Logf("Gas price: %s", gasPrice)
		// Should be a hex string
		if gasPrice[:2] != "0x" {
			t.Error("Gas price should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthEstimateGas(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a simple transfer call
	estimate, err := config.Client.RpcEthEstimateGas(ctx, TestAddresses.VitalikButerin, "0x", &RpcEthEstimateGasOpts{
		Value:   &[]string{"0x0"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("RpcEthEstimateGas failed: %v", err)
	}

	if estimate == "" {
		t.Error("Gas estimate is empty")
	} else {
		t.Logf("Gas estimate: %s", estimate)
		// Should be a hex string
		if estimate[:2] != "0x" {
			t.Error("Gas estimate should be a hex string starting with 0x")
		}
	}
}

func TestRpcEthBlockNumberWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	blockNumber, err := config.Client.RpcEthBlockNumber(ctx, nil)
	if err != nil {
		t.Fatalf("RpcEthBlockNumber with nil opts failed: %v", err)
	}

	if blockNumber == "" {
		t.Error("Block number is empty")
	} else {
		t.Logf("Current block number (with default opts): %s", blockNumber)
	}
}

func TestRpcEthGetGasPriceWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	gasPrice, err := config.Client.RpcEthGetGasPrice(ctx, nil)
	if err != nil {
		t.Fatalf("RpcEthGetGasPrice with nil opts failed: %v", err)
	}

	if gasPrice == "" {
		t.Error("Gas price is empty")
	} else {
		t.Logf("Gas price (with default opts): %s", gasPrice)
	}
}
