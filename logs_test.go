package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetEventLogsByAddress(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract address
	logs, err := config.Client.GetEventLogsByAddress(ctx, TestAddresses.USDTContract, &GetEventLogsByAddressOpts{
		FromBlock: 18000000, // Recent block
		ToBlock:   18000100, // Small range
		Page:      1,
		Offset:    10,
		ChainID:   1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByAddress failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for test contract address")
	} else {
		t.Logf("Found %d event logs", len(logs))
		// Validate first log
		log := logs[0]
		if log.Address == "" {
			t.Error("Log Address field is empty")
		}
		if log.BlockNumber == "" {
			t.Error("Log BlockNumber field is empty")
		}
		if log.TransactionHash == "" {
			t.Error("Log TransactionHash field is empty")
		}
	}
}

func TestGetEventLogsByTopics(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Transfer event topic (ERC20 Transfer)
	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	logs, err := config.Client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
		FromBlock: 18000000, // Recent block
		ToBlock:   18000100, // Small range
		Topic0:    transferTopic,
		Page:      1,
		Offset:    10,
		ChainID:   1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByTopics failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for Transfer topic")
	} else {
		t.Logf("Found %d event logs for Transfer topic", len(logs))
		// Validate first log
		log := logs[0]
		if log.Address == "" {
			t.Error("Log Address field is empty")
		}
		if log.BlockNumber == "" {
			t.Error("Log BlockNumber field is empty")
		}
		if log.TransactionHash == "" {
			t.Error("Log TransactionHash field is empty")
		}
	}
}

func TestGetEventLogsByAddressFilteredByTopics(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract and Transfer event topic
	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	logs, err := config.Client.GetEventLogsByAddressFilteredByTopics(ctx, TestAddresses.USDTContract, &GetEventLogsByAddressFilteredByTopicsOpts{
		FromBlock: 18000000, // Recent block
		ToBlock:   18000100, // Small range
		Topic0:    transferTopic,
		Page:      1,
		Offset:    10,
		ChainID:   1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByAddressFilteredByTopics failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for USDT contract with Transfer topic")
	} else {
		t.Logf("Found %d event logs for USDT contract with Transfer topic", len(logs))
		// Validate first log
		log := logs[0]
		if log.Address == "" {
			t.Error("Log Address field is empty")
		}
		if log.BlockNumber == "" {
			t.Error("Log BlockNumber field is empty")
		}
		if log.TransactionHash == "" {
			t.Error("Log TransactionHash field is empty")
		}
	}
}

func TestGetEventLogsByAddressWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	logs, err := config.Client.GetEventLogsByAddress(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetEventLogsByAddress with nil opts failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for test contract address (with default opts)")
	} else {
		t.Logf("Found %d event logs (with default opts)", len(logs))
	}
}

func TestGetEventLogsByTopicsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	logs, err := config.Client.GetEventLogsByTopics(ctx, nil)
	if err != nil {
		t.Fatalf("GetEventLogsByTopics with nil opts failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found (with default opts)")
	} else {
		t.Logf("Found %d event logs (with default opts)", len(logs))
	}
}

func TestGetEventLogsByAddressFilteredByTopicsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	logs, err := config.Client.GetEventLogsByAddressFilteredByTopics(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetEventLogsByAddressFilteredByTopics with nil opts failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for test contract address (with default opts)")
	} else {
		t.Logf("Found %d event logs (with default opts)", len(logs))
	}
}

func TestGetEventLogsByAddressWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	logs, err := config.Client.GetEventLogsByAddress(ctx, TestAddresses.USDTContract, &GetEventLogsByAddressOpts{
		FromBlock: 40000000, // Recent block on Polygon
		ToBlock:   40000100, // Small range
		Page:      1,
		Offset:    10,
		ChainID:   137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByAddress with Polygon failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found on Polygon for test contract address")
	} else {
		t.Logf("Found %d event logs on Polygon", len(logs))
	}
}

func TestGetEventLogsByTopicsWithMultipleTopics(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Transfer event topic and specific address in topic1
	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// Vitalik's address as topic1 (from address)
	vitalikTopic := "0x000000000000000000000000" + TestAddresses.VitalikButerin[2:] // Remove 0x prefix and pad
	logs, err := config.Client.GetEventLogsByTopics(ctx, &GetEventLogsByTopicsOpts{
		FromBlock: 18000000, // Recent block
		ToBlock:   18000100, // Small range
		Topic0:    transferTopic,
		Topic1:    vitalikTopic,
		Page:      1,
		Offset:    10,
		ChainID:   1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByTopics with multiple topics failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for Transfer topic with Vitalik as sender")
	} else {
		t.Logf("Found %d event logs for Transfer topic with Vitalik as sender", len(logs))
	}
}

func TestGetEventLogsByAddressFilteredByTopicsWithMultipleTopics(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract, Transfer event topic, and specific address in topic1
	transferTopic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// Vitalik's address as topic1 (from address)
	vitalikTopic := "0x000000000000000000000000" + TestAddresses.VitalikButerin[2:] // Remove 0x prefix and pad
	logs, err := config.Client.GetEventLogsByAddressFilteredByTopics(ctx, TestAddresses.USDTContract, &GetEventLogsByAddressFilteredByTopicsOpts{
		FromBlock: 18000000, // Recent block
		ToBlock:   18000100, // Small range
		Topic0:    transferTopic,
		Topic1:    vitalikTopic,
		Page:      1,
		Offset:    10,
		ChainID:   1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEventLogsByAddressFilteredByTopics with multiple topics failed: %v", err)
	}

	if len(logs) == 0 {
		t.Log("No event logs found for USDT contract with Transfer topic and Vitalik as sender")
	} else {
		t.Logf("Found %d event logs for USDT contract with Transfer topic and Vitalik as sender", len(logs))
	}
}
