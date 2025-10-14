package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetERC20TokenTransfers(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with address parameter
	transfers, err := config.Client.GetERC20TokenTransfers(ctx, &GetERC20TokenTransfersOpts{
		Address: &TestAddresses.VitalikButerin,
		Page:    &[]int64{1}[0],
		Offset:  &[]int64{10}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC20TokenTransfers failed: %v", err)
	}

	if len(transfers) == 0 {
		t.Log("No ERC20 transfers found for test address")
	} else {
		t.Logf("Found %d ERC20 transfers", len(transfers))
		// Validate first transfer
		transfer := transfers[0]
		if transfer.From == "" {
			t.Error("Transfer From field is empty")
		}
		if transfer.To == "" {
			t.Error("Transfer To field is empty")
		}
		if transfer.Value == "" {
			t.Error("Transfer Value field is empty")
		}
	}
}

func TestGetERC721TokenTransfers(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with address parameter
	transfers, err := config.Client.GetERC721TokenTransfers(ctx, &GetERC721TokenTransfersOpts{
		Address: &TestAddresses.VitalikButerin,
		Page:    &[]int64{1}[0],
		Offset:  &[]int64{10}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC721TokenTransfers failed: %v", err)
	}

	if len(transfers) == 0 {
		t.Log("No ERC721 transfers found for test address")
	} else {
		t.Logf("Found %d ERC721 transfers", len(transfers))
		// Validate first transfer
		transfer := transfers[0]
		if transfer.From == "" {
			t.Error("Transfer From field is empty")
		}
		if transfer.To == "" {
			t.Error("Transfer To field is empty")
		}
		if transfer.TokenID == "" {
			t.Error("Transfer TokenID field is empty")
		}
	}
}

func TestGetERC1155TokenTransfers(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with address parameter
	transfers, err := config.Client.GetERC1155TokenTransfers(ctx, &GetERC1155TokenTransfersOpts{
		Address: &TestAddresses.VitalikButerin,
		Page:    &[]int64{1}[0],
		Offset:  &[]int64{10}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetERC1155TokenTransfers failed: %v", err)
	}

	if len(transfers) == 0 {
		t.Log("No ERC1155 transfers found for test address")
	} else {
		t.Logf("Found %d ERC1155 transfers", len(transfers))
		// Validate first transfer
		transfer := transfers[0]
		if transfer.From == "" {
			t.Error("Transfer From field is empty")
		}
		if transfer.To == "" {
			t.Error("Transfer To field is empty")
		}
		if transfer.TokenID == "" {
			t.Error("Transfer TokenID field is empty")
		}
	}
}

func TestGetAddressFundedBy(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	funding, err := config.Client.GetAddressFundedBy(ctx, TestAddresses.VitalikButerin, &GetAddressFundedByOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetAddressFundedBy failed: %v", err)
	}

	if funding == nil {
		t.Log("No funding information found for test address")
	} else {
		t.Logf("Funding information: %+v", funding)
		if funding.Block == 0 {
			t.Error("Funding Block field is empty")
		}
		if funding.FundingAddress == "" {
			t.Error("Funding FundingAddress field is empty")
		}
	}
}

func TestGetBlocksValidatedByAddress(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	blocks, err := config.Client.GetBlocksValidatedByAddress(ctx, TestAddresses.VitalikButerin, &GetBlocksValidatedByAddressOpts{
		BlockType: &[]string{"blocks"}[0],
		Page:      &[]int64{1}[0],
		Offset:    &[]int64{10}[0],
		ChainID:   &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBlocksValidatedByAddress failed: %v", err)
	}

	if len(blocks) == 0 {
		t.Log("No blocks validated by test address")
	} else {
		t.Logf("Found %d validated blocks", len(blocks))
		// Validate first block
		block := blocks[0]
		if block.BlockNumber == "" {
			t.Error("Block BlockNumber field is empty")
		}
		if block.TimeStamp == "" {
			t.Error("Block TimeStamp field is empty")
		}
	}
}

func TestGetBeaconChainWithdrawals(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	withdrawals, err := config.Client.GetBeaconChainWithdrawals(ctx, TestAddresses.VitalikButerin, &GetBeaconChainWithdrawalsOpts{
		StartBlock: &[]int64{0}[0],
		EndBlock:   &[]int64{999999999999}[0],
		Page:       &[]int64{1}[0],
		Offset:     &[]int64{10}[0],
		Sort:       &[]string{"asc"}[0],
		ChainID:    &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetBeaconChainWithdrawals failed: %v", err)
	}

	if len(withdrawals) == 0 {
		t.Log("No beacon chain withdrawals found for test address")
	} else {
		t.Logf("Found %d beacon chain withdrawals", len(withdrawals))
		// Validate first withdrawal
		withdrawal := withdrawals[0]
		if withdrawal.WithdrawalIndex == "" {
			t.Error("Withdrawal WithdrawalIndex field is empty")
		}
		if withdrawal.ValidatorIndex == "" {
			t.Error("Withdrawal ValidatorIndex field is empty")
		}
	}
}

func TestGetEthBalanceByBlockNumber(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	balance, err := config.Client.GetEthBalanceByBlockNumber(ctx, TestAddresses.VitalikButerin, TestBlocks.RecentBlock, &GetEthBalanceByBlockNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetEthBalanceByBlockNumber failed: %v", err)
	}

	if balance == "" {
		t.Error("Balance is empty")
	} else {
		t.Logf("Balance at block %d: %s", TestBlocks.RecentBlock, balance)
	}
}

func TestGetContractCreatorAndCreation(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	contractAddresses := []string{TestAddresses.USDTContract}
	creations, err := config.Client.GetContractCreatorAndCreation(ctx, contractAddresses, &GetContractCreatorAndCreationOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetContractCreatorAndCreation failed: %v", err)
	}

	if len(creations) == 0 {
		t.Log("No contract creation information found")
	} else {
		t.Logf("Found %d contract creations", len(creations))
		// Validate first creation
		creation := creations[0]
		if creation.ContractAddress == "" {
			t.Error("Creation ContractAddress field is empty")
		}
		if creation.ContractCreator == "" {
			t.Error("Creation ContractCreator field is empty")
		}
	}
}
