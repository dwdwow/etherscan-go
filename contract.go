package etherscan

import (
	"context"
	"fmt"
	"strconv"
)

// ============================================================================
// Contract Module
// ============================================================================

// GetContractABIOpts contains optional parameters for GetContractABI
type GetContractABIOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetContractABI returns the Contract Application Binary Interface (ABI) of a verified smart contract
//
// This endpoint returns the ABI (Application Binary Interface) of a verified smart contract.
// The ABI is a JSON array that describes the contract's functions, events, and data structures,
// which is essential for interacting with the contract programmatically.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The contract address that has verified source code
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: The contract ABI as a JSON string
//   - error: Error if the request fails or contract is not verified
//
// Example:
//
//	// Get ABI for a verified contract
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	abi, err := client.GetContractABI(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Parse the ABI JSON string
//	var abiInterface []interface{}
//	if err := json.Unmarshal([]byte(abi), &abiInterface); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use the ABI to interact with the contract
//	for _, item := range abiInterface {
//	    if itemMap, ok := item.(map[string]interface{}); ok {
//	        if name, exists := itemMap["name"]; exists {
//	            fmt.Printf("Function: %s\n", name)
//	        }
//	    }
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	abi, err := client.GetContractABI(ctx, contractAddr, &GetContractABIOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Only works for contracts with verified source code
//   - Returns empty string if contract is not verified
//   - ABI is returned as a JSON string that needs to be parsed
//   - Essential for programmatic contract interaction
func (c *HTTPClient) GetContractABI(ctx context.Context, address string, opts *GetContractABIOpts) (string, error) {
	params := map[string]string{
		"address": address,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "getabi",
		params:          params,
		noFoundReturn:   "",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}

// GetContractSourceCodeOpts contains optional parameters for GetContractSourceCode
type GetContractSourceCodeOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// GetContractSourceCode returns the Solidity source code of a verified smart contract
//
// This endpoint returns the complete source code of a verified smart contract along with
// compilation details, ABI, and metadata. This is useful for code analysis, security audits,
// and understanding contract functionality.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - address: The contract address that has verified source code
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - []RespContractSourceCode: List containing contract source code details
//   - error: Error if the request fails or contract is not verified
//
// Example:
//
//	// Get source code for a verified contract
//	contractAddr := "0xA0b86a33E6441b8C4C5C8C0C0C0C0C0C0C0C0C0"
//	sourceCode, err := client.GetContractSourceCode(ctx, contractAddr, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if len(sourceCode) > 0 {
//	    contract := sourceCode[0]
//	    fmt.Printf("Contract Name: %s\n", contract.ContractName)
//	    fmt.Printf("Compiler Version: %s\n", contract.CompilerVersion)
//	    fmt.Printf("Optimization Used: %s\n", contract.OptimizationUsed)
//	    fmt.Printf("Source Code:\n%s\n", contract.SourceCode)
//	}
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	sourceCode, err := client.GetContractSourceCode(ctx, contractAddr, &GetContractSourceCodeOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Only works for contracts with verified source code
//   - Returns empty slice if contract is not verified
//   - Contains compilation metadata and settings
//   - SourceCode field contains the actual Solidity source
//   - ABI field contains the contract's ABI
func (c *HTTPClient) GetContractSourceCode(ctx context.Context, address string, opts *GetContractSourceCodeOpts) ([]RespContractSourceCode, error) {
	params := map[string]string{
		"address": address,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "getsourcecode",
		params:          params,
		noFoundReturn:   []RespContractSourceCode{},
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return nil, err
	}

	var result []RespContractSourceCode
	if err := unmarshalResponse(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// VerifySourceCodeOpts contains optional parameters for VerifySourceCode
type VerifySourceCodeOpts struct {
	// ConstructorArguments are the constructor arguments used when deploying the contract
	// If the contract has constructor parameters, provide them here
	ConstructorArguments *string

	// CompilerMode specifies the compiler mode (e.g., "solc/zksync" for ZK Stack)
	// Use this for special compilation modes or ZK Stack contracts
	CompilerMode *string

	// ZksolcVersion specifies the zkSolc version for ZK Stack (e.g., "v1.3.14")
	// Required when using ZK Stack compilation
	ZksolcVersion *string

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// VerifySourceCode submits contract source code for verification
//
// This endpoint submits contract source code for verification on Etherscan. After
// successful verification, the contract source code will be publicly visible and
// the ABI will be available for programmatic interaction.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - sourceCode: The Solidity source code to verify
//   - contractAddress: The deployed contract address
//   - contractName: Contract name (e.g., "contracts/Verified.sol:Verified")
//   - compilerVersion: Compiler version (e.g., "v0.8.24+commit.e11b9ed9")
//   - codeFormat: Source code format ("solidity-single-file" or "solidity-standard-json-input")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Verification GUID for tracking verification status
//   - error: Error if the request fails
//
// Example:
//
//	// Verify a simple contract
//	sourceCode := `pragma solidity ^0.8.0;
//	contract MyContract {
//	    uint256 public value;
//	    constructor(uint256 _value) { value = _value; }
//	}`
//	contractAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	contractName := "MyContract.sol:MyContract"
//	compilerVer := "v0.8.24+commit.e11b9ed9"
//	codeFormat := "solidity-single-file"
//
//	guid, err := client.VerifySourceCode(ctx, sourceCode, contractAddr, contractName, compilerVer, codeFormat, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Verification GUID: %s\n", guid)
//
//	// With constructor arguments
//	constructorArgs := "0000000000000000000000000000000000000000000000000000000000000064" // 100 in hex
//	guid, err := client.VerifySourceCode(ctx, sourceCode, contractAddr, contractName, compilerVer, codeFormat, &VerifySourceCodeOpts{
//	    ConstructorArguments: &constructorArgs,
//	})
//
//	// For ZK Stack contracts
//	guid, err := client.VerifySourceCode(ctx, sourceCode, contractAddr, contractName, compilerVer, codeFormat, &VerifySourceCodeOpts{
//	    CompilerMode:  &[]string{"solc/zksync"}[0],
//	    ZksolcVersion: &[]string{"v1.3.14"}[0],
//	})
//
// Note:
//   - Returns GUID for tracking verification status
//   - Use CheckSourceCodeVerificationStatus to check verification progress
//   - Constructor arguments must be ABI-encoded
//   - Source code must match the deployed bytecode exactly
func (c *HTTPClient) VerifySourceCode(ctx context.Context, sourceCode, contractAddress, contractName, compilerVersion, codeFormat string, opts *VerifySourceCodeOpts) (string, error) {
	params := map[string]string{
		"sourceCode":      sourceCode,
		"contractaddress": contractAddress,
		"contractname":    contractName,
		"compilerversion": compilerVersion,
		"codeformat":      codeFormat,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ConstructorArguments != nil {
			params["constructorArguements"] = *opts.ConstructorArguments
		}
		if opts.CompilerMode != nil {
			params["compilermode"] = *opts.CompilerMode
		}
		if opts.ZksolcVersion != nil {
			params["zksolcVersion"] = *opts.ZksolcVersion
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "verifysourcecode",
		params:          params,
		method:          "POST",
		noFoundReturn:   "",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}

// VerifyVyperSourceCodeOpts contains optional parameters for VerifyVyperSourceCode
type VerifyVyperSourceCodeOpts struct {
	// ConstructorArguments are the constructor arguments used when deploying the contract
	// If the contract has constructor parameters, provide them here
	ConstructorArguments *string

	// OptimizationUsed specifies whether optimization was used during compilation
	// Options: 0 (no optimization) or 1 (optimization used)
	OptimizationUsed *int

	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// VerifyVyperSourceCode submits Vyper contract source code for verification
//
// This endpoint submits Vyper contract source code for verification on Etherscan.
// Vyper is a Python-like programming language for Ethereum smart contracts. After
// successful verification, the contract source code will be publicly visible.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - sourceCode: The Vyper source code in JSON format
//   - contractAddress: The deployed contract address
//   - contractName: Contract name (e.g., "contracts/Verified.vy:Verified")
//   - compilerVersion: Vyper compiler version (e.g., "vyper:0.4.0")
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Verification GUID for tracking verification status
//   - error: Error if the request fails
//
// Example:
//
//	// Verify a Vyper contract
//	sourceCode := `{"language": "Vyper", "sources": {"contract.vy": {"content": "@external\ndef hello() -> String[100]:\n    return \"Hello, World!\""}}}`
//	contractAddr := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
//	contractName := "contract.vy:hello"
//	compilerVer := "vyper:0.4.0"
//
//	guid, err := client.VerifyVyperSourceCode(ctx, sourceCode, contractAddr, contractName, compilerVer, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Verification GUID: %s\n", guid)
//
//	// With constructor arguments and optimization
//	constructorArgs := "0000000000000000000000000000000000000000000000000000000000000064"
//	optimization := 1
//	guid, err := client.VerifyVyperSourceCode(ctx, sourceCode, contractAddr, contractName, compilerVer, &VerifyVyperSourceCodeOpts{
//	    ConstructorArguments: &constructorArgs,
//	    OptimizationUsed:     &optimization,
//	})
//
// Note:
//   - Source code must be in JSON format for Vyper
//   - Returns GUID for tracking verification status
//   - Use CheckSourceCodeVerificationStatus to check verification progress
//   - Constructor arguments must be ABI-encoded
func (c *HTTPClient) VerifyVyperSourceCode(ctx context.Context, sourceCode, contractAddress, contractName, compilerVersion string, opts *VerifyVyperSourceCodeOpts) (string, error) {
	params := map[string]string{
		"sourceCode":      sourceCode,
		"contractaddress": contractAddress,
		"contractname":    contractName,
		"compilerversion": compilerVersion,
		"codeformat":      "vyper-json",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ConstructorArguments != nil {
			params["constructorArguments"] = *opts.ConstructorArguments
		}
		if opts.OptimizationUsed != nil {
			params["optimizationUsed"] = strconv.Itoa(*opts.OptimizationUsed)
		}
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "verifysourcecode",
		params:          params,
		method:          "POST",
		noFoundReturn:   "",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}

// VerifyStylusSourceCodeOpts contains optional parameters for Stylus source code verification
type VerifyStylusSourceCodeOpts struct {
	ChainID         *int
	OnLimitExceeded *RateLimitBehavior
}

// VerifyStylusSourceCode submits Stylus contract source code for verification
//
// Args:
//   - sourceCode: The Github link to the source code
//   - contractAddress: The deployed contract address
//   - contractName: Contract name (e.g. "stylus_hello_world")
//   - compilerVersion: Stylus compiler version (e.g. "stylus:0.5.3")
//   - licenseType: Open source license type (e.g. 3 for MIT)
//
// Example:
//
//	guid, err := client.VerifyStylusSourceCode(ctx, githubURL, contractAddr, contractName, compilerVer, 3, nil)
func (c *HTTPClient) VerifyStylusSourceCode(ctx context.Context, sourceCode, contractAddress, contractName, compilerVersion string, licenseType int, opts *VerifyStylusSourceCodeOpts) (string, error) {
	params := map[string]string{
		"sourceCode":      sourceCode,
		"contractaddress": contractAddress,
		"contractname":    contractName,
		"compilerversion": compilerVersion,
		"licenseType":     strconv.Itoa(licenseType),
		"codeformat":      "stylus",
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "verifysourcecode",
		params:          params,
		method:          "POST",
		noFoundReturn:   "",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}

// CheckSourceCodeVerificationStatusOpts contains optional parameters for CheckSourceCodeVerificationStatus
type CheckSourceCodeVerificationStatusOpts struct {
	// ChainID specifies which blockchain network to query
	// If nil, uses the client's default chain ID (EthereumMainnet = 1)
	// Supported chains: EthereumMainnet, PolygonMainnet, ArbitrumOneMainnet, etc.
	ChainID *int

	// OnLimitExceeded specifies behavior when rate limit is exceeded
	// If nil, uses the client's default behavior (RateLimitBlock)
	// Options:
	//   - RateLimitBlock: Wait until a token is available (default)
	//   - RateLimitRaise: Return an error when rate limit is exceeded
	//   - RateLimitSkip: Return false without executing when rate limit is exceeded
	OnLimitExceeded *RateLimitBehavior
}

// CheckSourceCodeVerificationStatus checks the verification status of a submitted source code verification request
//
// This endpoint checks the verification status of a previously submitted source code
// verification request using the GUID returned from VerifySourceCode. This is useful
// for tracking verification progress and determining when verification is complete.
//
// Args:
//   - ctx: Context for request cancellation and timeout
//   - guid: The unique GUID received from the verification request
//   - opts: Optional parameters (can be nil)
//
// Returns:
//   - string: Verification status message
//   - error: Error if the request fails
//
// Example:
//
//	// Check verification status
//	guid := "your-verification-guid-here"
//	status, err := client.CheckSourceCodeVerificationStatus(ctx, guid, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Verification Status: %s\n", status)
//
//	// With custom chain ID
//	chainID := 137 // Polygon
//	status, err := client.CheckSourceCodeVerificationStatus(ctx, guid, &CheckSourceCodeVerificationStatusOpts{
//	    ChainID: &chainID,
//	})
//
// Note:
//   - Use the GUID returned from VerifySourceCode
//   - Verification may take several minutes to complete
//   - Returns status message indicating verification progress
//   - Check periodically until verification is complete
func (c *HTTPClient) CheckSourceCodeVerificationStatus(ctx context.Context, guid string, opts *CheckSourceCodeVerificationStatusOpts) (string, error) {
	params := map[string]string{
		"guid": guid,
	}

	var onLimitExceeded *RateLimitBehavior
	if opts != nil {
		if opts.ChainID != nil {
			params["chainid"] = strconv.Itoa(*opts.ChainID)
		}
		onLimitExceeded = opts.OnLimitExceeded
	}

	data, err := c.request(requestParams{
		ctx:             ctx,
		module:          "contract",
		action:          "checkverifystatus",
		params:          params,
		noFoundReturn:   "",
		onLimitExceeded: onLimitExceeded,
	})
	if err != nil {
		return "", err
	}

	if str, ok := data.(string); ok {
		return str, nil
	}
	return fmt.Sprintf("%v", data), nil
}
