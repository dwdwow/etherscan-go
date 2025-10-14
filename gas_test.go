package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetConfirmationTimeEstimate(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a reasonable gas price (20 Gwei)
	gasPrice := int64(20000000000) // 20 Gwei in wei
	estimate, err := config.Client.GetConfirmationTimeEstimate(ctx, gasPrice, &GetConfirmationTimeEstimateOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetConfirmationTimeEstimate failed: %v", err)
	}

	if estimate == "" {
		t.Error("Confirmation time estimate is empty")
	} else {
		t.Logf("Confirmation time estimate for %d wei: %s", gasPrice, estimate)
	}
}

func TestGetGasOracle(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	oracle, err := config.Client.GetGasOracle(ctx, &GetGasOracleOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetGasOracle failed: %v", err)
	}

	if oracle == nil {
		t.Error("Gas oracle is nil")
	} else {
		t.Logf("Gas oracle: %+v", oracle)
		if oracle.SafeGasPrice == "" {
			t.Error("SafeGasPrice field is empty")
		}
		if oracle.ProposeGasPrice == "" {
			t.Error("ProposeGasPrice field is empty")
		}
		if oracle.FastGasPrice == "" {
			t.Error("FastGasPrice field is empty")
		}
	}
}

func TestGetDailyAverageGasLimit(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	limits, err := config.Client.GetDailyAverageGasLimit(ctx, startDate, endDate, &GetDailyAverageGasLimitOpts{
		Sort:    &[]string{"asc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAverageGasLimit failed: %v", err)
	}

	if len(limits) == 0 {
		t.Log("No daily average gas limits found for the date range")
	} else {
		t.Logf("Found %d daily average gas limits", len(limits))
		// Validate first entry
		limit := limits[0]
		if limit.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if limit.GasLimit == "" {
			t.Error("GasLimit field is empty")
		}
		t.Logf("First entry: Date=%s, GasLimit=%s", limit.UTCDate, limit.GasLimit)
	}
}

func TestGetDailyTotalGasUsed(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	gasUsed, err := config.Client.GetDailyTotalGasUsed(ctx, startDate, endDate, &GetDailyTotalGasUsedOpts{
		Sort:    &[]string{"asc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyTotalGasUsed failed: %v", err)
	}

	if len(gasUsed) == 0 {
		t.Log("No daily total gas used found for the date range")
	} else {
		t.Logf("Found %d daily total gas used entries", len(gasUsed))
		// Validate first entry
		gas := gasUsed[0]
		if gas.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if gas.GasUsed == "" {
			t.Error("GasUsed field is empty")
		}
		t.Logf("First entry: Date=%s, GasUsed=%s", gas.UTCDate, gas.GasUsed)
	}
}

func TestGetDailyAverageGasPrice(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a date range from 2023
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	prices, err := config.Client.GetDailyAverageGasPrice(ctx, startDate, endDate, &GetDailyAverageGasPriceOpts{
		Sort:    &[]string{"asc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAverageGasPrice failed: %v", err)
	}

	if len(prices) == 0 {
		t.Log("No daily average gas prices found for the date range")
	} else {
		t.Logf("Found %d daily average gas prices", len(prices))
		// Validate first entry
		price := prices[0]
		if price.UTCDate == "" {
			t.Error("UTCDate field is empty")
		}
		if price.AvgGasPriceWei == "" {
			t.Error("AvgGasPriceWei field is empty")
		}
		t.Logf("First entry: Date=%s, AvgGasPriceWei=%s", price.UTCDate, price.AvgGasPriceWei)
	}
}

func TestGetConfirmationTimeEstimateWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	gasPrice := int64(20000000000) // 20 Gwei in wei
	estimate, err := config.Client.GetConfirmationTimeEstimate(ctx, gasPrice, nil)
	if err != nil {
		t.Fatalf("GetConfirmationTimeEstimate with nil opts failed: %v", err)
	}

	if estimate == "" {
		t.Error("Confirmation time estimate is empty")
	} else {
		t.Logf("Confirmation time estimate with default opts: %s", estimate)
	}
}

func TestGetGasOracleWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	oracle, err := config.Client.GetGasOracle(ctx, nil)
	if err != nil {
		t.Fatalf("GetGasOracle with nil opts failed: %v", err)
	}

	if oracle == nil {
		t.Error("Gas oracle is nil")
	} else {
		t.Logf("Gas oracle with default opts: %+v", oracle)
	}
}

func TestGetDailyAverageGasLimitWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	limits, err := config.Client.GetDailyAverageGasLimit(ctx, startDate, endDate, &GetDailyAverageGasLimitOpts{
		Sort:    &[]string{"desc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAverageGasLimit with desc sort failed: %v", err)
	}

	if len(limits) == 0 {
		t.Log("No daily average gas limits found for the date range")
	} else {
		t.Logf("Found %d daily average gas limits (desc order)", len(limits))
	}
}

func TestGetDailyTotalGasUsedWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	gasUsed, err := config.Client.GetDailyTotalGasUsed(ctx, startDate, endDate, &GetDailyTotalGasUsedOpts{
		Sort:    &[]string{"desc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyTotalGasUsed with desc sort failed: %v", err)
	}

	if len(gasUsed) == 0 {
		t.Log("No daily total gas used found for the date range")
	} else {
		t.Logf("Found %d daily total gas used entries (desc order)", len(gasUsed))
	}
}

func TestGetDailyAverageGasPriceWithDescSort(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with descending sort
	startDate := "2023-01-01"
	endDate := "2023-01-07"
	prices, err := config.Client.GetDailyAverageGasPrice(ctx, startDate, endDate, &GetDailyAverageGasPriceOpts{
		Sort:    &[]string{"desc"}[0],
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetDailyAverageGasPrice with desc sort failed: %v", err)
	}

	if len(prices) == 0 {
		t.Log("No daily average gas prices found for the date range")
	} else {
		t.Logf("Found %d daily average gas prices (desc order)", len(prices))
	}
}
