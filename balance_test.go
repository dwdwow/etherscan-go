package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetEthBalance(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test single address balance
	balance, err := config.Client.GetEthBalance(ctx, TestAddresses.VitalikButerin, &GetEthBalanceOpts{
		Tag:     "latest",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalance failed: %v", err)
	}

	if balance == "" {
		t.Error("Balance is empty")
	} else {
		t.Logf("Balance for %s: %s", TestAddresses.VitalikButerin, balance)
	}
}

func TestGetEthBalances(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test multiple addresses balance
	addresses := []string{
		TestAddresses.VitalikButerin,
		TestAddresses.USDTContract,
		TestAddresses.USDCContract,
	}

	balances, err := config.Client.GetEthBalances(ctx, addresses, &GetEthBalancesOpts{
		Tag:     "latest",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalances failed: %v", err)
	}

	if len(balances) == 0 {
		t.Error("No balances returned")
	} else {
		t.Logf("Found %d balances", len(balances))
		// Validate first balance
		balance := balances[0]
		if balance.Account == "" {
			t.Error("Balance Account field is empty")
		}
		if balance.Balance == "" {
			t.Error("Balance Balance field is empty")
		}
		t.Logf("Balance for %s: %s", balance.Account, balance.Balance)
	}
}

func TestGetEthBalanceWithPendingTag(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with pending tag
	balance, err := config.Client.GetEthBalance(ctx, TestAddresses.VitalikButerin, &GetEthBalanceOpts{
		Tag:     "pending",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalance with pending tag failed: %v", err)
	}

	if balance == "" {
		t.Error("Balance is empty")
	} else {
		t.Logf("Pending balance for %s: %s", TestAddresses.VitalikButerin, balance)
	}
}

func TestGetEthBalanceWithLatestTag(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with earliest tag
	balance, err := config.Client.GetEthBalance(ctx, TestAddresses.VitalikButerin, &GetEthBalanceOpts{
		Tag:     "latest",
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalance with latest tag failed: %v", err)
	}

	if balance == "" {
		t.Error("Balance is empty")
	} else {
		t.Logf("Earliest balance for %s: %s", TestAddresses.VitalikButerin, balance)
	}
}

func TestGetEthBalancesWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	addresses := []string{TestAddresses.VitalikButerin}
	balances, err := config.Client.GetEthBalances(ctx, addresses, &GetEthBalancesOpts{
		Tag:     "latest",
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalances with Polygon failed: %v", err)
	}

	if len(balances) == 0 {
		t.Error("No balances returned for Polygon")
	} else {
		t.Logf("Found %d balances on Polygon", len(balances))
		balance := balances[0]
		t.Logf("Polygon balance for %s: %s", balance.Account, balance.Balance)
	}
}

func TestGetEthBalanceWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	balance, err := config.Client.GetEthBalance(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetEthBalance with nil opts failed: %v", err)
	}

	if balance == "" {
		t.Error("Balance is empty")
	} else {
		t.Logf("Default balance for %s: %s", TestAddresses.VitalikButerin, balance)
	}
}

func TestGetEthBalancesWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	addresses := []string{TestAddresses.VitalikButerin}
	balances, err := config.Client.GetEthBalances(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetEthBalances with nil opts failed: %v", err)
	}

	if len(balances) == 0 {
		t.Error("No balances returned")
	} else {
		t.Logf("Found %d default balances", len(balances))
		balance := balances[0]
		t.Logf("Default balance for %s: %s", balance.Account, balance.Balance)
	}
}
