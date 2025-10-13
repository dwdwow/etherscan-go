# Etherscan Go Client

完整的 Etherscan V2 API Go 客户端，支持所有主要的 API 端点。

## 特性

✅ **完整的 API 覆盖** - 实现了 88+ 个 API 方法  
✅ **Context 支持** - 所有方法都支持 context.Context 用于超时和取消  
✅ **速率限制** - 内置多级速率限制器，支持不同的 API 层级  
✅ **类型安全** - 完整的类型定义和响应结构  
✅ **可选参数** - 使用指针选项模式，方便传递 nil  
✅ **多链支持** - 支持 80+ 条区块链网络  

## 已实现的功能

### 📊 统计信息
- **代码行数**: 6566 行
- **API 方法数**: 88 个公共方法
- **文件数**: 15 个文件
- **模块数**: 10+ 个模块

### 📁 文件组织

| 文件 | 大小 | 说明 |
|------|------|------|
| `http_client.go` | 23KB | 核心客户端和请求处理 |
| `http_resp.go` | 33KB | 所有响应类型定义 |
| `account.go` | 14KB | 账户模块方法 |
| `contract.go` | 8.8KB | 合约模块方法 |
| `block.go` | 6.8KB | 区块模块方法 |
| `transaction.go` | 2.6KB | 交易模块方法 |
| `logs.go` | 6.6KB | 日志模块方法 |
| `proxy.go` | 15KB | RPC 代理方法 |
| `token.go` | 14KB | Token 模块方法 |
| `gas.go` | 5.7KB | Gas 追踪模块方法 |
| `stats.go` | 18KB | 统计模块方法 |
| `layer2.go` | 4.5KB | Layer 2 模块方法 |
| `admin.go` | 6.4KB | 管理模块方法 |
| `ratelimiter.go` | 14KB | 速率限制器实现 |

## 安装

```bash
go get github.com/yourusername/etherscan-go
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    etherscan "github.com/yourusername/etherscan-go"
)

func main() {
    // 创建客户端
    client := etherscan.NewHTTPClient(etherscan.HTTPClientConfig{
        APIKey:         "YOUR_API_KEY",
        DefaultChainID: etherscan.EthereumMainnet,
        APITier:        etherscan.ProPlusTier,
    })
    
    ctx := context.Background()
    
    // 获取 ETH 余额
    balance, err := client.GetEthBalance(ctx, "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance: %s wei\n", balance)
}
```

## API 模块

### 1. Account Module (账户模块)

#### 余额查询
- `GetEthBalance` - 获取 ETH 余额
- `GetEthBalances` - 批量获取 ETH 余额 (最多20个地址)
- `GetEthBalanceByBlockNumber` - 获取指定区块的历史余额

#### 交易查询
- `GetNormalTransactions` - 获取普通交易列表
- `GetInternalTransactionsByAddress` - 获取内部交易 (按地址)
- `GetInternalTransactionsByHash` - 获取内部交易 (按哈希)
- `GetInternalTransactionsByBlockRange` - 获取内部交易 (按区块范围)
- `GetBridgeTransactions` - 获取跨链桥交易

#### Token 转账
- `GetERC20TokenTransfers` - 获取 ERC-20 代币转账记录
- `GetERC721TokenTransfers` - 获取 ERC-721 NFT 转账记录
- `GetERC1155TokenTransfers` - 获取 ERC-1155 代币转账记录

#### 其他
- `GetAddressFundedBy` - 获取地址资金来源
- `GetBlocksValidatedByAddress` - 获取地址验证的区块
- `GetBeaconChainWithdrawals` - 获取信标链提款记录

### 2. Contract Module (合约模块)

- `GetContractABI` - 获取合约 ABI
- `GetContractSourceCode` - 获取合约源代码
- `GetContractCreatorAndCreation` - 获取合约创建者和创建交易
- `VerifySourceCode` - 提交 Solidity 源代码验证
- `VerifyVyperSourceCode` - 提交 Vyper 源代码验证
- `VerifyStylusSourceCode` - 提交 Stylus 源代码验证
- `CheckSourceCodeVerificationStatus` - 检查验证状态

### 3. Transaction Module (交易模块)

- `GetContractExecutionStatus` - 获取合约执行状态
- `GetTransactionReceiptStatus` - 获取交易收据状态

### 4. Block Module (区块模块)

- `GetBlockAndUncleRewards` - 获取区块和叔块奖励
- `GetBlockTransactionsCount` - 获取区块交易数量
- `GetBlockCountdownTime` - 获取区块倒计时
- `GetBlockNumberByTimestamp` - 根据时间戳获取区块号
- `GetDailyAvgBlockSizes` - 获取每日平均区块大小
- `GetDailyBlockCountRewards` - 获取每日区块数量和奖励
- `GetDailyBlockRewards` - 获取每日区块奖励
- `GetDailyAvgBlockTime` - 获取每日平均出块时间
- `GetDailyUncleBlockCountAndRewards` - 获取每日叔块统计

### 5. Logs Module (日志模块)

- `GetEventLogsByAddress` - 根据地址获取事件日志
- `GetEventLogsByTopics` - 根据主题获取事件日志
- `GetEventLogsByAddressFilteredByTopics` - 根据地址和主题过滤事件日志

### 6. Geth/Parity Proxy Module (RPC 代理模块)

#### 区块查询
- `RpcEthBlockNumber` - 获取最新区块号
- `RpcEthBlockByNumber` - 根据区块号获取区块信息
- `RpcEthUncleByBlockNumberAndIndex` - 获取叔块信息
- `RpcEthBlockTransactionCountByNumber` - 获取区块交易数量

#### 交易查询
- `RpcEthTransactionByHash` - 根据哈希获取交易
- `RpcEthTransactionByBlockNumberAndIndex` - 根据区块号和索引获取交易
- `RpcEthTransactionCount` - 获取地址交易数量
- `RpcEthTransactionReceipt` - 获取交易收据

#### 交易发送
- `RpcEthSendRawTransaction` - 发送原始交易

#### 合约调用
- `RpcEthCall` - 执行合约调用
- `RpcEthGetCode` - 获取合约代码
- `RpcEthGetStorageAt` - 获取存储值
- `RpcEthEstimateGas` - 估算 gas 费用

#### Gas 相关
- `RpcEthGetGasPrice` - 获取 gas 价格

### 7. Token Module (代币模块)

#### ERC-20 相关
- `GetERC20TotalSupply` - 获取代币总供应量
- `GetERC20AccountBalance` - 获取代币余额
- `GetERC20HistoricalTotalSupply` - 获取历史总供应量
- `GetERC20HistoricalAccountBalance` - 获取历史余额
- `GetERC20Holders` - 获取代币持有者列表
- `GetERC20HolderCount` - 获取持有者数量
- `GetTopERC20Holders` - 获取代币前N持有者

#### 账户持仓
- `GetTokenInfo` - 获取代币信息
- `GetAccountERC20Holdings` - 获取账户 ERC-20 持仓
- `GetAccountNFTHoldings` - 获取账户 NFT 持仓
- `GetAccountNFTInventories` - 获取账户 NFT 清单

### 8. Gas Tracker Module (Gas 追踪模块)

- `GetConfirmationTimeEstimate` - 估算确认时间
- `GetGasOracle` - 获取 gas 预言机数据
- `GetDailyAverageGasLimit` - 获取每日平均 gas 限制
- `GetDailyTotalGasUsed` - 获取每日总 gas 消耗
- `GetDailyAverageGasPrice` - 获取每日平均 gas 价格

### 9. Stats Module (统计模块)

#### 供应和价格
- `GetTotalEthSupply` - 获取 ETH 总供应量
- `GetTotalEth2Supply` - 获取 ETH2 总供应量
- `GetEthPrice` - 获取 ETH 价格
- `GetEthHistoricalPrices` - 获取历史价格

#### 网络统计
- `GetEthereumNodesSize` - 获取节点大小
- `GetNodeCount` - 获取节点数量
- `GetDailyTxFees` - 获取每日交易费用
- `GetDailyNewAddresses` - 获取每日新地址数
- `GetDailyNetworkUtilizations` - 获取每日网络利用率
- `GetDailyAvgHashrates` - 获取每日平均算力
- `GetDailyTxCounts` - 获取每日交易数
- `GetDailyAvgDifficulties` - 获取每日平均难度

### 10. Layer 2 Module (Layer 2 模块)

- `GetPlasmaDeposits` - 获取 Plasma 存款 (Polygon)
- `GetDepositTxs` - 获取存款交易 (Arbitrum/Optimism)
- `GetWithdrawalTxs` - 获取提款交易 (Arbitrum/Optimism)

### 11. Admin Module (管理模块)

#### 地址标签
- `GetAddressTag` - 获取地址标签
- `GetLabelMasterlist` - 获取标签主列表
- `ExportSpecificLabelCSV` - 导出特定标签的 CSV
- `ExportOFACSanctionedRelatedLabelsCSV` - 导出 OFAC 制裁地址 CSV
- `ExportAllAddressTagsCSV` - 导出所有地址标签 CSV
- `GetLatestCSVBatchNumber` - 获取最新 CSV 批次号

#### API 管理
- `CheckCreditUsage` - 检查 API 额度使用情况

### 12. Chain Info Module (链信息模块)

- `GetSupportedChains` - 获取支持的区块链列表

## 支持的区块链

支持 80+ 条区块链，包括：

- Ethereum Mainnet
- Polygon
- Arbitrum
- Optimism
- Base
- BSC
- Avalanche
- zkSync
- Linea
- Scroll
- 以及更多...

完整列表请查看 `http_client.go` 中的链 ID 常量定义。

## 使用示例

### 获取多个地址余额

```go
addresses := []string{
    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
}

balances, err := client.GetEthBalances(ctx, addresses, nil)
if err != nil {
    log.Fatal(err)
}

for _, bal := range balances {
    fmt.Printf("Address: %s, Balance: %s wei\n", bal.Account, bal.Balance)
}
```

### 获取交易列表

```go
page := 1
offset := 10
sort := "desc"

txs, err := client.GetNormalTransactions(ctx, address, &etherscan.GetNormalTransactionsOpts{
    Page:   &page,
    Offset: &offset,
    Sort:   &sort,
})
```

### 使用 Context 超时

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

balance, err := client.GetEthBalance(ctx, address, nil)
```

### 指定链 ID

```go
chainID := etherscan.PolygonMainnet
balance, err := client.GetEthBalance(ctx, address, &etherscan.GetEthBalanceOpts{
    ChainID: &chainID,
})
```

### 自定义速率限制行为

```go
skipBehavior := etherscan.RateLimitSkip
balance, err := client.GetEthBalance(ctx, address, &etherscan.GetEthBalanceOpts{
    OnLimitExceeded: &skipBehavior,
})
```

## API 层级和速率限制

| 层级 | 每秒请求数 | 每日请求数 |
|------|-----------|----------|
| Free | 5 | 100,000 |
| Standard | 10 | 200,000 |
| Advanced | 20 | 500,000 |
| Professional | 30 | 1,000,000 |
| Pro Plus | 30 | 1,500,000 |

## 速率限制行为

- `RateLimitBlock` - 阻塞等待直到可用 (默认)
- `RateLimitRaise` - 抛出错误
- `RateLimitSkip` - 跳过请求，返回 false

## 错误处理

```go
balance, err := client.GetEthBalance(ctx, address, nil)
if err != nil {
    if err == context.DeadlineExceeded {
        fmt.Println("请求超时")
    } else if err == etherscan.ErrRateLimitExceeded {
        fmt.Println("速率限制")
    } else {
        fmt.Printf("错误: %v\n", err)
    }
    return
}
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行带覆盖率的测试
go test -cover ./...

# 运行速率限制器测试
go test -run TestRateLimiter
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 更新日志

### v1.0.0 (2025-01-14)
- ✅ 实现所有 88+ API 方法
- ✅ 完整的 Context 支持
- ✅ 多级速率限制器
- ✅ 支持 80+ 条区块链
- ✅ 完整的类型定义
- ✅ 详细的文档和示例

## 相关链接

- [Etherscan API 文档](https://docs.etherscan.io/)
- [Etherscan V2 API](https://docs.etherscan.io/v/etherscan-v2/)
