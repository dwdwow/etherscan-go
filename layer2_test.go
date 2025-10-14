package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetPlasmaDeposits(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known address
	deposits, err := config.Client.GetPlasmaDeposits(ctx, TestAddresses.VitalikButerin, &GetPlasmaDepositsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetPlasmaDeposits failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No plasma deposits found for test address")
	} else {
		t.Logf("Found %d plasma deposits", len(deposits))
		// Validate first deposit
		deposit := deposits[0]
		if deposit.BlockNumber == "" {
			t.Error("Deposit BlockNumber field is empty")
		}
		if deposit.TimeStamp == "" {
			t.Error("Deposit TimeStamp field is empty")
		}
		if deposit.BlockReward == "" {
			t.Error("Deposit BlockReward field is empty")
		}
	}
}

func TestGetDepositTxs(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known address
	deposits, err := config.Client.GetDepositTxs(ctx, TestAddresses.VitalikButerin, &GetDepositTxsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDepositTxs failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No deposit transactions found for test address")
	} else {
		t.Logf("Found %d deposit transactions", len(deposits))
		// Validate first deposit
		deposit := deposits[0]
		if deposit.From == "" {
			t.Error("Deposit From field is empty")
		}
		if deposit.To == "" {
			t.Error("Deposit To field is empty")
		}
		if deposit.Value == "" {
			t.Error("Deposit Value field is empty")
		}
	}
}

func TestGetWithdrawalTxs(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a known address
	withdrawals, err := config.Client.GetWithdrawalTxs(ctx, TestAddresses.VitalikButerin, &GetWithdrawalTxsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 1, // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetWithdrawalTxs failed: %v", err)
	}

	if len(withdrawals) == 0 {
		t.Log("No withdrawal transactions found for test address")
	} else {
		t.Logf("Found %d withdrawal transactions", len(withdrawals))
		// Validate first withdrawal
		withdrawal := withdrawals[0]
		if withdrawal.From == "" {
			t.Error("Withdrawal From field is empty")
		}
		if withdrawal.To == "" {
			t.Error("Withdrawal To field is empty")
		}
		if withdrawal.Value == "" {
			t.Error("Withdrawal Value field is empty")
		}
	}
}

func TestGetPlasmaDepositsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	deposits, err := config.Client.GetPlasmaDeposits(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetPlasmaDeposits with nil opts failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No plasma deposits found for test address (with default opts)")
	} else {
		t.Logf("Found %d plasma deposits (with default opts)", len(deposits))
	}
}

func TestGetDepositTxsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	deposits, err := config.Client.GetDepositTxs(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetDepositTxs with nil opts failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No deposit transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d deposit transactions (with default opts)", len(deposits))
	}
}

func TestGetWithdrawalTxsWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	withdrawals, err := config.Client.GetWithdrawalTxs(ctx, TestAddresses.VitalikButerin, nil)
	if err != nil {
		t.Fatalf("GetWithdrawalTxs with nil opts failed: %v", err)
	}

	if len(withdrawals) == 0 {
		t.Log("No withdrawal transactions found for test address (with default opts)")
	} else {
		t.Logf("Found %d withdrawal transactions (with default opts)", len(withdrawals))
	}
}

func TestGetPlasmaDepositsWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	deposits, err := config.Client.GetPlasmaDeposits(ctx, TestAddresses.VitalikButerin, &GetPlasmaDepositsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetPlasmaDeposits with Polygon failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No plasma deposits found on Polygon for test address")
	} else {
		t.Logf("Found %d plasma deposits on Polygon", len(deposits))
	}
}

func TestGetDepositTxsWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	deposits, err := config.Client.GetDepositTxs(ctx, TestAddresses.VitalikButerin, &GetDepositTxsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetDepositTxs with Polygon failed: %v", err)
	}

	if len(deposits) == 0 {
		t.Log("No deposit transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d deposit transactions on Polygon", len(deposits))
	}
}

func TestGetWithdrawalTxsWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	withdrawals, err := config.Client.GetWithdrawalTxs(ctx, TestAddresses.VitalikButerin, &GetWithdrawalTxsOpts{
		Page:    1,
		Offset:  10,
		ChainID: 137, // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetWithdrawalTxs with Polygon failed: %v", err)
	}

	if len(withdrawals) == 0 {
		t.Log("No withdrawal transactions found on Polygon for test address")
	} else {
		t.Logf("Found %d withdrawal transactions on Polygon", len(withdrawals))
	}
}
