package etherscan

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestGetContractABI(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract (should be verified)
	abi, err := config.Client.GetContractABI(ctx, TestAddresses.USDTContract, &GetContractABIOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetContractABI failed: %v", err)
	}

	if abi == "" {
		t.Log("Contract ABI is empty (contract may not be verified)")
	} else {
		t.Logf("Contract ABI length: %d characters", len(abi))
		// ABI should be a JSON string
		if abi[0] != '[' && abi[0] != '{' {
			t.Error("ABI should be a JSON string")
		}
	}
}

func TestGetContractSourceCode(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with USDT contract (should be verified)
	sourceCodes, err := config.Client.GetContractSourceCode(ctx, TestAddresses.USDTContract, &GetContractSourceCodeOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetContractSourceCode failed: %v", err)
	}

	if len(sourceCodes) == 0 {
		t.Log("No source code found (contract may not be verified)")
	} else {
		t.Logf("Found %d source code entries", len(sourceCodes))
		// Validate first entry
		sourceCode := sourceCodes[0]
		if sourceCode.ContractName == "" {
			t.Log("ContractName is empty")
		}
		if sourceCode.CompilerVersion == "" {
			t.Log("CompilerVersion is empty")
		}
		if sourceCode.SourceCode == "" {
			t.Log("SourceCode is empty")
		}
		t.Logf("Contract: %s, Compiler: %s", sourceCode.ContractName, sourceCode.CompilerVersion)
	}
}

func TestGetContractABIWithUnverifiedContract(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with an unverified contract (random address)
	unverifiedAddress := "0x1234567890123456789012345678901234567890"
	abi, err := config.Client.GetContractABI(ctx, unverifiedAddress, &GetContractABIOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		if strings.Contains(err.Error(), "Contract source code not verified") {
			t.Skip("Contract source code not verified")
		}
		t.Fatalf("GetContractABI failed: %v", err)
	}

	if abi == "" {
		t.Log("Contract ABI is empty (expected for unverified contract)")
	} else {
		t.Logf("Unexpected ABI found for unverified contract: %s", abi)
	}
}

func TestGetContractSourceCodeWithUnverifiedContract(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with an unverified contract (random address)
	unverifiedAddress := "0x1234567890123456789012345678901234567890"
	sourceCodes, err := config.Client.GetContractSourceCode(ctx, unverifiedAddress, &GetContractSourceCodeOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetContractSourceCode failed: %v", err)
	}

	if len(sourceCodes) == 0 {
		t.Log("No source code found (expected for unverified contract)")
	} else {
		t.Logf("Unexpected source code found for unverified contract: %d entries", len(sourceCodes))
	}
}

func TestGetContractABIWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	abi, err := config.Client.GetContractABI(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetContractABI with nil opts failed: %v", err)
	}

	if abi == "" {
		t.Log("Contract ABI is empty (contract may not be verified)")
	} else {
		t.Logf("Contract ABI with default opts: %d characters", len(abi))
	}
}

func TestGetContractSourceCodeWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	sourceCodes, err := config.Client.GetContractSourceCode(ctx, TestAddresses.USDTContract, nil)
	if err != nil {
		t.Fatalf("GetContractSourceCode with nil opts failed: %v", err)
	}

	if len(sourceCodes) == 0 {
		t.Log("No source code found (contract may not be verified)")
	} else {
		t.Logf("Source code with default opts: %d entries", len(sourceCodes))
	}
}

func TestGetContractABIWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	abi, err := config.Client.GetContractABI(ctx, TestAddresses.USDTContract, &GetContractABIOpts{
		ChainID: &[]int64{137}[0], // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetContractABI with Polygon failed: %v", err)
	}

	if abi == "" {
		t.Log("Contract ABI is empty on Polygon (contract may not be verified)")
	} else {
		t.Logf("Contract ABI on Polygon: %d characters", len(abi))
	}
}

func TestGetContractSourceCodeWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	sourceCodes, err := config.Client.GetContractSourceCode(ctx, TestAddresses.USDTContract, &GetContractSourceCodeOpts{
		ChainID: &[]int64{137}[0], // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetContractSourceCode with Polygon failed: %v", err)
	}

	if len(sourceCodes) == 0 {
		t.Log("No source code found on Polygon (contract may not be verified)")
	} else {
		t.Logf("Source code on Polygon: %d entries", len(sourceCodes))
	}
}

// Note: The following tests are for contract verification methods
// These are typically not tested in integration tests as they require
// actual contract deployment and verification process

func TestVerifySourceCode(t *testing.T) {
	t.Skip("Skipping VerifySourceCode test - requires actual contract deployment")
}

func TestVerifyVyperSourceCode(t *testing.T) {
	t.Skip("Skipping VerifyVyperSourceCode test - requires actual contract deployment")
}

func TestVerifyStylusSourceCode(t *testing.T) {
	t.Skip("Skipping VerifyStylusSourceCode test - requires actual contract deployment")
}

func TestCheckSourceCodeVerificationStatus(t *testing.T) {
	t.Skip("Skipping CheckSourceCodeVerificationStatus test - requires verification GUID")
}
