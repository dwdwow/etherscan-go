package etherscan

import (
	"context"
	"fmt"
	"testing"
)

func TestHTTPClient_GetSupportedChains(t *testing.T) {
	client := NewHTTPClient(HTTPClientConfig{})

	chains, err := client.GetSupportedChains(context.Background())
	if err != nil {
		t.Fatalf("GetSupportedChains failed: %v", err)
	}

	for _, chain := range chains.Result {
		fmt.Println(chain.ChainName)
	}
}
