package etherscan

import (
	"context"
	"testing"
	"time"
)

func TestGetAddressTag(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Vitalik's address
	addresses := []string{TestAddresses.VitalikButerin}
	tags, err := config.Client.GetAddressTag(ctx, addresses, &GetAddressTagOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetAddressTag failed: %v", err)
	}

	if len(tags) == 0 {
		t.Log("No address tags found for test address")
	} else {
		t.Logf("Address tags: %+v", tags)
		// Validate first tag
		tag := tags[0]
		if tag.Address == "" {
			t.Error("Tag Address field is empty")
		}
		if tag.Nametag == "" {
			t.Error("Tag Nametag field is empty")
		}
	}
}

func TestGetLabelMasterlist(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	labels, err := config.Client.GetLabelMasterlist(ctx, &GetLabelMasterlistOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetLabelMasterlist failed: %v", err)
	}

	if len(labels) == 0 {
		t.Log("No labels found in masterlist")
	} else {
		t.Logf("Found %d labels in masterlist", len(labels))
		// Validate first label
		label := labels[0]
		if label.LabelName == "" {
			t.Error("Label LabelName field is empty")
		}
		if label.LabelSlug == "" {
			t.Error("Label LabelSlug field is empty")
		}
		t.Logf("First label: %s - %s", label.LabelName, label.LabelSlug)
	}
}

func TestExportSpecificLabelCSV(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with a specific label (e.g., "Exchange")
	label := "Exchange"
	csv, err := config.Client.ExportSpecificLabelCSV(ctx, label)
	if err != nil {
		t.Fatalf("ExportSpecificLabelCSV failed: %v", err)
	}

	if len(csv) == 0 {
		t.Log("No CSV data found for specific label")
	} else {
		t.Logf("CSV data length: %d bytes", len(csv))
		// CSV should contain headers and data
		if len(csv) < 100 {
			t.Error("CSV data seems too short")
		}
	}
}

func TestExportOFACSanctionedRelatedLabelsCSV(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	csv, err := config.Client.ExportOFACSanctionedRelatedLabelsCSV(ctx)
	if err != nil {
		t.Fatalf("ExportOFACSanctionedRelatedLabelsCSV failed: %v", err)
	}

	if len(csv) == 0 {
		t.Log("No OFAC sanctioned related labels CSV data found")
	} else {
		t.Logf("OFAC CSV data length: %d bytes", len(csv))
		// CSV should contain headers and data
		if len(csv) < 100 {
			t.Error("CSV data seems too short")
		}
	}
}

func TestExportAllAddressTagsCSV(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	csv, err := config.Client.ExportAllAddressTagsCSV(ctx)
	if err != nil {
		t.Fatalf("ExportAllAddressTagsCSV failed: %v", err)
	}

	if len(csv) == 0 {
		t.Log("No address tags CSV data found")
	} else {
		t.Logf("Address tags CSV data length: %d bytes", len(csv))
		// CSV should contain headers and data
		if len(csv) < 100 {
			t.Error("CSV data seems too short")
		}
	}
}

func TestGetLatestCSVBatchNumber(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	batchNumbers, err := config.Client.GetLatestCSVBatchNumber(ctx, &GetLatestCSVBatchNumberOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("GetLatestCSVBatchNumber failed: %v", err)
	}

	if len(batchNumbers) == 0 {
		t.Error("Latest CSV batch numbers is empty")
	} else {
		t.Logf("Latest CSV batch numbers: %+v", batchNumbers)
	}
}

func TestCheckCreditUsage(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	usage, err := config.Client.CheckCreditUsage(ctx, &CheckCreditUsageOpts{
		ChainID: &[]int64{1}[0], // Ethereum mainnet
	})
	if err != nil {
		t.Fatalf("CheckCreditUsage failed: %v", err)
	}

	if usage == nil {
		t.Error("Credit usage is nil")
	} else {
		t.Logf("Credit usage: %+v", usage)
		if usage.CreditsUsed == 0 {
			t.Error("CreditsUsed field is empty")
		}
		if usage.CreditsAvailable == 0 {
			t.Error("CreditsAvailable field is empty")
		}
		if usage.CreditLimit == 0 {
			t.Error("CreditLimit field is empty")
		}
	}
}

func TestGetAddressTagWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	addresses := []string{TestAddresses.VitalikButerin}
	tags, err := config.Client.GetAddressTag(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetAddressTag with nil opts failed: %v", err)
	}

	if len(tags) == 0 {
		t.Log("No address tags found for test address (with default opts)")
	} else {
		t.Logf("Address tags (with default opts): %+v", tags)
	}
}

func TestGetLabelMasterlistWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	labels, err := config.Client.GetLabelMasterlist(ctx, nil)
	if err != nil {
		t.Fatalf("GetLabelMasterlist with nil opts failed: %v", err)
	}

	if len(labels) == 0 {
		t.Log("No labels found in masterlist (with default opts)")
	} else {
		t.Logf("Found %d labels in masterlist (with default opts)", len(labels))
	}
}

func TestCheckCreditUsageWithNilOpts(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with nil opts (should use defaults)
	usage, err := config.Client.CheckCreditUsage(ctx, nil)
	if err != nil {
		t.Fatalf("CheckCreditUsage with nil opts failed: %v", err)
	}

	if usage == nil {
		t.Error("Credit usage is nil")
	} else {
		t.Logf("Credit usage (with default opts): %+v", usage)
	}
}

func TestGetAddressTagWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	addresses := []string{TestAddresses.VitalikButerin}
	tags, err := config.Client.GetAddressTag(ctx, addresses, &GetAddressTagOpts{
		ChainID: &[]int64{137}[0], // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetAddressTag with Polygon failed: %v", err)
	}

	if len(tags) == 0 {
		t.Log("No address tags found on Polygon for test address")
	} else {
		t.Logf("Address tags on Polygon: %+v", tags)
	}
}

func TestGetLabelMasterlistWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	labels, err := config.Client.GetLabelMasterlist(ctx, &GetLabelMasterlistOpts{
		ChainID: &[]int64{137}[0], // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("GetLabelMasterlist with Polygon failed: %v", err)
	}

	if len(labels) == 0 {
		t.Log("No labels found in masterlist on Polygon")
	} else {
		t.Logf("Found %d labels in masterlist on Polygon", len(labels))
	}
}

func TestExportSpecificLabelCSVWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	label := "Exchange"
	csv, err := config.Client.ExportSpecificLabelCSV(ctx, label)
	if err != nil {
		t.Fatalf("ExportSpecificLabelCSV with Polygon failed: %v", err)
	}

	if len(csv) == 0 {
		t.Log("No CSV data found for specific label on Polygon")
	} else {
		t.Logf("CSV data length on Polygon: %d bytes", len(csv))
	}
}

func TestCheckCreditUsageWithDifferentChain(t *testing.T) {
	config := GetTestConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test with Polygon mainnet
	usage, err := config.Client.CheckCreditUsage(ctx, &CheckCreditUsageOpts{
		ChainID: &[]int64{137}[0], // Polygon mainnet
	})
	if err != nil {
		t.Fatalf("CheckCreditUsage with Polygon failed: %v", err)
	}

	if usage == nil {
		t.Error("Credit usage is nil on Polygon")
	} else {
		t.Logf("Credit usage on Polygon: %+v", usage)
	}
}
