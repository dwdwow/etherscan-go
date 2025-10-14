package etherscan

// Account Module Response Types

type RespGetEthBalance string

// RespBridgeTx represents a bridge transaction
// Example:
//
//	{
//	    "hash": "0x79eaf9951f474d5fd78bae5e3d6e089b54e61b81d272ce6f98e8ac6b56ec0f93",
//	    "blockNumber": "42107370",
//	    "timeStamp": "1757792370",
//	    "from": "0xfa98b60e02a61b6590f073cad56e68326652d094",
//	    "address": "0x1545c4ccf40a5e89ac1482a32485f62d369560d3",
//	    "amount": "1000000000000000000",
//	    "tokenName": "",
//	    "symbol": "",
//	    "contractAddress": "0x7301cfa0e1756b71869e93d4e4dca5c7d0eb0aa6",
//	    "divisor": ""
//	}
type RespBridgeTx struct {
	Hash            string `json:"hash"`
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	From            string `json:"from"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	TokenName       string `json:"tokenName"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contractAddress"`
	Divisor         string `json:"divisor"`
}

type RespGetBridgeTxs []RespBridgeTx

// RespEthBalanceEntry represents an Ethereum balance entry
// Example:
//
//	{
//	    "account": "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
//	    "balance": "40891626854930000000000"
//	}
type RespEthBalanceEntry struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

type RespGetEthBalances []RespEthBalanceEntry

// RespNormalTx represents a normal transaction
// Example:
//
//	{
//	    "blockNumber": "14923678",
//	    "timeStamp": "1654646411",
//	    "hash": "0xc52783ad354aecc04c670047754f062e3d6d04e8f5b24774472651f9c3882c60",
//	    ...
//	}
type RespNormalTx struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

type RespGetNormalTxs []RespNormalTx

// RespInternalTxByAddress represents an internal transaction by address
type RespInternalTxByAddress struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceID         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

type RespGetInternalTxsByAddress []RespInternalTxByAddress

// RespInternalTxByHash represents an internal transaction by hash
type RespInternalTxByHash struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

type RespGetInternalTxsByHash []RespInternalTxByHash

// RespInternalTxByBlockRange represents an internal transaction by block range
type RespInternalTxByBlockRange struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceID         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

type RespGetInternalTxsByBlockRange []RespInternalTxByBlockRange

// RespERC20TokenTransfer represents an ERC-20 token transfer
type RespERC20TokenTransfer struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

type RespGetERC20TokenTransfers []RespERC20TokenTransfer

// RespERC721TokenTransfer represents an ERC-721 token transfer
type RespERC721TokenTransfer struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

type RespGetERC721TokenTransfers []RespERC721TokenTransfer

// RespERC1155TokenTransfer represents an ERC-1155 token transfer
type RespERC1155TokenTransfer struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	From              string `json:"from"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenValue        string `json:"tokenValue"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	Confirmations     string `json:"confirmations"`
}

type RespGetERC1155TokenTransfers []RespERC1155TokenTransfer

// RespAddressFundedBy represents address funding information
type RespAddressFundedBy struct {
	Block          int64  `json:"block"`
	TimeStamp      string `json:"timeStamp"`
	FundingAddress string `json:"fundingAddress"`
	FundingTxn     string `json:"fundingTxn"`
	Value          string `json:"value"`
}

// RespBlockValidated represents a validated block
type RespBlockValidated struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
	BlockReward string `json:"blockReward"`
}

type RespBlocksValidatedByAddress []RespBlockValidated

// RespBeaconChainWithdrawal represents a beacon chain withdrawal
type RespBeaconChainWithdrawal struct {
	WithdrawalIndex string `json:"withdrawalIndex"`
	ValidatorIndex  string `json:"validatorIndex"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	BlockNumber     string `json:"blockNumber"`
	Timestamp       string `json:"timestamp"`
}

type RespBeaconChainWithdrawals []RespBeaconChainWithdrawal

type RespEthBalanceByBlockNumber string

// Contract Module Response Types

type RespContractABI string

// RespContractSourceCode represents contract source code information
type RespContractSourceCode struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	OptimizationUsed     string `json:"OptimizationUsed"`
	Runs                 string `json:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
	SimilarMatch         string `json:"SimilarMatch"`
}

type RespContractSourceCodes []RespContractSourceCode

// RespContractCreationAndCreation represents contract creation information
type RespContractCreationAndCreation struct {
	ContractAddress  string `json:"contractAddress"`
	ContractCreator  string `json:"contractCreator"`
	TxHash           string `json:"txHash"`
	BlockNumber      string `json:"blockNumber"`
	Timestamp        string `json:"timestamp"`
	ContractFactory  string `json:"contractFactory"`
	CreationBytecode string `json:"creationBytecode"`
}

type RespContractCreationAndCreations []RespContractCreationAndCreation

type RespVerifySourceCode string

type RespVerifyVyperSourceCode string

type RespVerifyStylusSourceCode string

type RespCheckSourceCodeVerificationStatus string

// RespContractExecutionStatus represents contract execution status
type RespContractExecutionStatus struct {
	IsError        string `json:"isError"`
	ErrDescription string `json:"errDescription"`
}

// RespCheckTxReceiptStatus represents transaction receipt status
type RespCheckTxReceiptStatus struct {
	Status string `json:"status"`
}

// Block Module Response Types

// UncleReward represents an uncle block reward
type UncleReward struct {
	Miner         string `json:"miner"`
	UnclePosition string `json:"unclePosition"`
	Blockreward   string `json:"blockreward"`
}

// RespBlockReward represents block reward information
type RespBlockReward struct {
	BlockNumber          string        `json:"blockNumber"`
	TimeStamp            string        `json:"timeStamp"`
	BlockMiner           string        `json:"blockMiner"`
	BlockReward          string        `json:"blockReward"`
	Uncles               []UncleReward `json:"uncles"`
	UncleInclusionReward string        `json:"uncleInclusionReward"`
}

// RespBlockTxsCountByBlockNo represents transaction count by block number
type RespBlockTxsCountByBlockNo struct {
	Block            int64 `json:"block"`
	TxsCount         int64 `json:"txsCount"`
	InternalTxsCount int64 `json:"internalTxsCount"`
	ERC20TxsCount    int64 `json:"erc20TxsCount"`
	ERC721TxsCount   int64 `json:"erc721TxsCount"`
	ERC1155TxsCount  int64 `json:"erc1155TxsCount"`
}

// RespEstimateBlockCountdownTimeByBlockNo represents estimated block countdown time
type RespEstimateBlockCountdownTimeByBlockNo struct {
	CurrentBlock      string `json:"CurrentBlock"`
	CountdownBlock    string `json:"CountdownBlock"`
	RemainingBlock    string `json:"RemainingBlock"`
	EstimateTimeInSec string `json:"EstimateTimeInSec"`
}

type RespBlockNumber int

// RespDailyAvgBlockSize represents daily average block size
type RespDailyAvgBlockSize struct {
	UTCDate        string `json:"UTCDate"`
	UnixTimeStamp  string `json:"unixTimeStamp"`
	BlockSizeBytes int64  `json:"blockSize_bytes"`
}

type RespDailyAvgBlockSizes []RespDailyAvgBlockSize

// RespDailyBlockCountReward represents daily block count and reward
type RespDailyBlockCountReward struct {
	UTCDate         string `json:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp"`
	BlockCount      int64  `json:"blockCount"`
	BlockRewardsEth string `json:"blockRewards_Eth"`
}

type RespDailyBlockCountRewards []RespDailyBlockCountReward

// RespDailyBlockReward represents daily block reward
type RespDailyBlockReward struct {
	UTCDate         string `json:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp"`
	BlockRewardsEth string `json:"blockRewards_Eth"`
}

type RespDailyBlockRewards []RespDailyBlockReward

// RespDailyAvgTimeBlockMined represents daily average time for block mining
type RespDailyAvgTimeBlockMined struct {
	UTCDate       string `json:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	BlockTimeSec  string `json:"blockTime_sec"`
}

type RespDailyAvgTimeBlockMineds []RespDailyAvgTimeBlockMined

// RespDailyUncleBlockCountAndReward represents daily uncle block count and reward
type RespDailyUncleBlockCountAndReward struct {
	UTCDate              string `json:"UTCDate"`
	UnixTimeStamp        string `json:"unixTimeStamp"`
	UncleBlockCount      int64  `json:"uncleBlockCount"`
	UncleBlockRewardsEth string `json:"uncleBlockRewards_Eth"`
}

type RespDailyUncleBlockCountAndRewards []RespDailyUncleBlockCountAndReward

// Logs Module Response Types

// RespEventLogByAddress represents an event log by address
type RespEventLogByAddress struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type RespEventLogsByAddress []RespEventLogByAddress

// RespEventLogByTopics represents an event log by topics
type RespEventLogByTopics struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type RespEventLogsByTopics []RespEventLogByTopics

// RespEventLogByAddressFilteredByTopics represents an event log by address filtered by topics
type RespEventLogByAddressFilteredByTopics struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type RespEventLogsByAddressFilteredByTopics []RespEventLogByAddressFilteredByTopics

// Geth/Parity Proxy Module Response Types

// RespJsonRpc represents a generic JSON-RPC response
type RespJsonRpc[Result any] struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  Result `json:"result"`
}

type RespEthBlockNumberHex = RespJsonRpc[string]

// RespEthBlockInfo represents Ethereum block information
type RespEthBlockInfo struct {
	BaseFeePerGas    string   `json:"baseFeePerGas"`
	Difficulty       string   `json:"difficulty"`
	ExtraData        string   `json:"extraData"`
	GasLimit         string   `json:"gasLimit"`
	GasUsed          string   `json:"gasUsed"`
	Hash             string   `json:"hash"`
	LogsBloom        string   `json:"logsBloom"`
	Miner            string   `json:"miner"`
	MixHash          string   `json:"mixHash"`
	Nonce            string   `json:"nonce"`
	Number           string   `json:"number"`
	ParentHash       string   `json:"parentHash"`
	ReceiptsRoot     string   `json:"receiptsRoot"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	Size             string   `json:"size"`
	StateRoot        string   `json:"stateRoot"`
	Timestamp        string   `json:"timestamp"`
	TotalDifficulty  string   `json:"totalDifficulty"`
	Transactions     []string `json:"transactions"`
	TransactionsRoot string   `json:"transactionsRoot"`
	Uncles           []string `json:"uncles"`
}

type RespEthBlock = RespJsonRpc[RespEthBlockInfo]

type RespEthBlockInfoWithFullTxs struct {
	BaseFeePerGas    string          `json:"baseFeePerGas"`
	Difficulty       string          `json:"difficulty"`
	ExtraData        string          `json:"extraData"`
	GasLimit         string          `json:"gasLimit"`
	GasUsed          string          `json:"gasUsed"`
	Hash             string          `json:"hash"`
	LogsBloom        string          `json:"logsBloom"`
	Miner            string          `json:"miner"`
	MixHash          string          `json:"mixHash"`
	Nonce            string          `json:"nonce"`
	Number           string          `json:"number"`
	ParentHash       string          `json:"parentHash"`
	ReceiptsRoot     string          `json:"receiptsRoot"`
	Sha3Uncles       string          `json:"sha3Uncles"`
	Size             string          `json:"size"`
	StateRoot        string          `json:"stateRoot"`
	Timestamp        string          `json:"timestamp"`
	TotalDifficulty  string          `json:"totalDifficulty"`
	Transactions     []RespEthTxInfo `json:"transactions"`
	TransactionsRoot string          `json:"transactionsRoot"`
	Uncles           []string        `json:"uncles"`
}

type RespEthBlockWithFullTxs = RespJsonRpc[RespEthBlockInfoWithFullTxs]

// RespEthUncleBlockInfo represents Ethereum uncle block information
type RespEthUncleBlockInfo struct {
	BaseFeePerGas    string   `json:"baseFeePerGas"`
	Difficulty       string   `json:"difficulty"`
	ExtraData        string   `json:"extraData"`
	GasLimit         string   `json:"gasLimit"`
	GasUsed          string   `json:"gasUsed"`
	Hash             string   `json:"hash"`
	LogsBloom        string   `json:"logsBloom"`
	Miner            string   `json:"miner"`
	MixHash          string   `json:"mixHash"`
	Nonce            string   `json:"nonce"`
	Number           string   `json:"number"`
	ParentHash       string   `json:"parentHash"`
	ReceiptsRoot     string   `json:"receiptsRoot"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	Size             string   `json:"size"`
	StateRoot        string   `json:"stateRoot"`
	Timestamp        string   `json:"timestamp"`
	TransactionsRoot string   `json:"transactionsRoot"`
	Uncles           []string `json:"uncles"`
}

type RespEthUncleBlock = RespJsonRpc[RespEthUncleBlockInfo]

type RespEthBlockTxCount = RespJsonRpc[string]

// RespEthTxInfo represents Ethereum transaction information
type RespEthTxInfo struct {
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	TransactionIndex     string        `json:"transactionIndex"`
	Value                string        `json:"value"`
	Type                 string        `json:"type"`
	AccessList           []interface{} `json:"accessList"`
	ChainID              string        `json:"chainId"`
	V                    string        `json:"v"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
}

type RespEthTx = RespJsonRpc[RespEthTxInfo]

type RespEthTxCount = RespJsonRpc[string]

type RespEthSendRawTx = RespJsonRpc[string]

// RespEthTxReceiptLog represents a transaction receipt log
type RespEthTxReceiptLog struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

// RespEthTxReceiptInfo represents Ethereum transaction receipt information
type RespEthTxReceiptInfo struct {
	BlockHash         string                `json:"blockHash"`
	BlockNumber       string                `json:"blockNumber"`
	ContractAddress   *string               `json:"contractAddress"`
	CumulativeGasUsed string                `json:"cumulativeGasUsed"`
	EffectiveGasPrice string                `json:"effectiveGasPrice"`
	From              string                `json:"from"`
	GasUsed           string                `json:"gasUsed"`
	Logs              []RespEthTxReceiptLog `json:"logs"`
	LogsBloom         string                `json:"logsBloom"`
	Status            string                `json:"status"`
	To                string                `json:"to"`
	TransactionHash   string                `json:"transactionHash"`
	TransactionIndex  string                `json:"transactionIndex"`
	Type              string                `json:"type"`
}

type RespEthTxReceipt = RespJsonRpc[RespEthTxReceiptInfo]

type RespEthCall = RespJsonRpc[string]

type RespEthGetCode = RespJsonRpc[string]

type RespEthGetStorageAt = RespJsonRpc[string]

type RespEthGetGasPrice = RespJsonRpc[string]

type RespEthEstimateGas = RespJsonRpc[string]

// Token Module Response Types

type RespERC20TotalSupply string

type RespERC20AccountBalance string

type RespERC20HistoricalTotalSupply string

type RespERC20HistoricalAccountBalance string

// RespERC20HolderInfo represents ERC-20 token holder information
type RespERC20HolderInfo struct {
	TokenHolderAddress  string `json:"TokenHolderAddress"`
	TokenHolderQuantity string `json:"TokenHolderQuantity"`
}

type RespERC20Holders []RespERC20HolderInfo

type RespERC20HolderCount string

// RespTopTokenHolder represents top token holder information
type RespTopTokenHolder struct {
	TokenHolderAddress     string `json:"TokenHolderAddress"`
	TokenHolderQuantity    string `json:"TokenHolderQuantity"`
	TokenHolderAddressType string `json:"TokenHolderAddressType"`
}

type RespTopTokenHolders []RespTopTokenHolder

// RespTokenInfo represents token information
type RespTokenInfo struct {
	ContractAddress string `json:"contractAddress"`
	TokenName       string `json:"tokenName"`
	Symbol          string `json:"symbol"`
	Divisor         string `json:"divisor"`
	TokenType       string `json:"tokenType"`
	TotalSupply     string `json:"totalSupply"`
	BlueCheckmark   string `json:"blueCheckmark"`
	Description     string `json:"description"`
	Website         string `json:"website"`
	Email           string `json:"email"`
	Blog            string `json:"blog"`
	Reddit          string `json:"reddit"`
	Slack           string `json:"slack"`
	Facebook        string `json:"facebook"`
	Twitter         string `json:"twitter"`
	Bitcointalk     string `json:"bitcointalk"`
	Github          string `json:"github"`
	Telegram        string `json:"telegram"`
	Wechat          string `json:"wechat"`
	Linkedin        string `json:"linkedin"`
	Discord         string `json:"discord"`
	Whitepaper      string `json:"whitepaper"`
	TokenPriceUSD   string `json:"tokenPriceUSD"`
}

// RespERC20Holding represents an ERC-20 token holding
type RespERC20Holding struct {
	TokenAddress  string `json:"TokenAddress"`
	TokenName     string `json:"TokenName"`
	TokenSymbol   string `json:"TokenSymbol"`
	TokenQuantity string `json:"TokenQuantity"`
	TokenDivisor  string `json:"TokenDivisor"`
}

type RespERC20Holdings []RespERC20Holding

// RespNFTHolding represents an NFT holding
type RespNFTHolding struct {
	TokenAddress  string `json:"TokenAddress"`
	TokenName     string `json:"TokenName"`
	TokenSymbol   string `json:"TokenSymbol"`
	TokenQuantity string `json:"TokenQuantity"`
}

type RespNFTHoldings []RespNFTHolding

// RespNFTTokenInventory represents an NFT token inventory item
type RespNFTTokenInventory struct {
	TokenAddress string `json:"TokenAddress"`
	TokenID      string `json:"TokenId"`
}

type RespNFTTokenInventories []RespNFTTokenInventory

// Gas Tracker Module Response Types

type RespConfirmationTimeEstimation string

// RespGasOracle represents gas oracle information
type RespGasOracle struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

// RespDailyAvgGasLimit represents daily average gas limit
type RespDailyAvgGasLimit struct {
	UTCDate       string `json:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	GasLimit      string `json:"gasLimit"`
}

type RespDailyAvgGasLimits []RespDailyAvgGasLimit

// RespDailyTotalGasUsed represents daily total gas used
type RespDailyTotalGasUsed struct {
	UTCDate       string `json:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	GasUsed       string `json:"gasUsed"`
}

type RespDailyTotalGasUseds []RespDailyTotalGasUsed

// RespDailyAvgGasPrice represents daily average gas price
type RespDailyAvgGasPrice struct {
	UTCDate        string `json:"UTCDate"`
	UnixTimeStamp  string `json:"unixTimeStamp"`
	MaxGasPriceWei string `json:"maxGasPrice_Wei"`
	MinGasPriceWei string `json:"minGasPrice_Wei"`
	AvgGasPriceWei string `json:"avgGasPrice_Wei"`
}

type RespDailyAvgGasPrices []RespDailyAvgGasPrice

// Stats Module Response Types

type RespTotalEthSupply string

type RespTotalEth2Supply string

// RespEthPrice represents Ethereum price information
type RespEthPrice struct {
	EthBTC          string `json:"ethbtc"`
	EthBTCTimestamp string `json:"ethbtc_timestamp"`
	EthUSD          string `json:"ethusd"`
	EthUSDTimestamp string `json:"ethusd_timestamp"`
}

// RespEtheumNodeSize represents Ethereum node size information
type RespEtheumNodeSize struct {
	BlockNumber    string `json:"blockNumber"`
	ChainTimeStamp string `json:"chainTimeStamp"`
	ChainSize      string `json:"chainSize"`
	ClientType     string `json:"clientType"`
	SyncMode       string `json:"syncMode"`
}

type RespEtheumNodesSize []RespEtheumNodeSize

// RespNodeCount represents node count information
type RespNodeCount struct {
	UTCDate        string `json:"UTCDate"`
	TotalNodeCount string `json:"TotalNodeCount"`
}

// RespDailyTxFee represents daily transaction fee
type RespDailyTxFee struct {
	UTCDate           string `json:"UTCDate"`
	UnixTimeStamp     string `json:"unixTimeStamp"`
	TransactionFeeEth string `json:"transactionFee_Eth"`
}

type RespDailyTxFees []RespDailyTxFee

// RespDailyNewAddress represents daily new address count
type RespDailyNewAddress struct {
	UTCDate         string `json:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp"`
	NewAddressCount int64  `json:"newAddressCount"`
}

type RespDailyNewAddresses []RespDailyNewAddress

// RespDailyNetworkUtilization represents daily network utilization
type RespDailyNetworkUtilization struct {
	UTCDate            string `json:"UTCDate"`
	UnixTimeStamp      string `json:"unixTimeStamp"`
	NetworkUtilization string `json:"networkUtilization"`
}

type RespDailyNetworkUtilizations []RespDailyNetworkUtilization

// RespDailyAvgHashrate represents daily average hashrate
type RespDailyAvgHashrate struct {
	UTCDate         string `json:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp"`
	NetworkHashRate string `json:"networkHashRate"`
}

type RespDailyAvgHashrates []RespDailyAvgHashrate

// RespDailyTxCount represents daily transaction count
type RespDailyTxCount struct {
	UTCDate          string `json:"UTCDate"`
	UnixTimeStamp    string `json:"unixTimeStamp"`
	TransactionCount int64  `json:"transactionCount"`
}

type RespDailyTxCounts []RespDailyTxCount

// RespDailyAvgDifficulty represents daily average difficulty
type RespDailyAvgDifficulty struct {
	UTCDate           string `json:"UTCDate"`
	UnixTimeStamp     string `json:"unixTimeStamp"`
	NetworkDifficulty string `json:"networkDifficulty"`
}

type RespDailyAvgDifficulties []RespDailyAvgDifficulty

// RespEthHistoricalPrice represents Ethereum historical price
type RespEthHistoricalPrice struct {
	UTCDate       string `json:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	Value         string `json:"value"`
}

type RespEthHistoricalPrices []RespEthHistoricalPrice

// Layer 2 Module Response Types

// RespPlasmaDeposit represents a plasma deposit
type RespPlasmaDeposit struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
	BlockReward string `json:"blockReward"`
}

type RespPlasmaDeposits []RespPlasmaDeposit

// RespDepositTx represents a deposit transaction
type RespDepositTx struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	BlockHash         string `json:"blockHash"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	Input             string `json:"input"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	IsError           string `json:"isError"`
	ErrDescription    string `json:"errDescription"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	QueueIndex        string `json:"queueIndex"`
	L1TransactionHash string `json:"L1transactionhash"`
	L1TxOrigin        string `json:"L1TxOrigin"`
	TokenAddress      string `json:"tokenAddress"`
	TokenSentFrom     string `json:"tokenSentFrom"`
	TokenSentTo       string `json:"tokenSentTo"`
	TokenValue        string `json:"tokenValue"`
}

type RespDepositTxs []RespDepositTx

// RespWithdrawalTx represents a withdrawal transaction
type RespWithdrawalTx struct {
	BlockNumber            string `json:"blockNumber"`
	TimeStamp              string `json:"timeStamp"`
	BlockHash              string `json:"blockHash"`
	Hash                   string `json:"hash"`
	Nonce                  string `json:"nonce"`
	From                   string `json:"from"`
	To                     string `json:"to"`
	Value                  string `json:"value"`
	Gas                    string `json:"gas"`
	GasPrice               string `json:"gasPrice"`
	Input                  string `json:"input"`
	CumulativeGasUsed      string `json:"cumulativeGasUsed"`
	GasUsed                string `json:"gasUsed"`
	IsError                string `json:"isError"`
	ErrDescription         string `json:"errDescription"`
	TxReceiptStatus        string `json:"txreceipt_status"`
	Message                string `json:"message"`
	MessageNonce           string `json:"messageNonce"`
	Status                 string `json:"status"`
	L1TransactionHash      string `json:"L1transactionhash"`
	TokenAddress           string `json:"tokenAddress"`
	WithdrawalType         string `json:"withdrawalType"`
	TokenValue             string `json:"tokenValue"`
	L1TransactionHashProve string `json:"L1transactionhashProve"`
}

type RespWithdrawalTxs []RespWithdrawalTx

// API Credit & Chain Info Module Response Types

// RespCreditUsage represents credit usage information
type RespCreditUsage struct {
	CreditsUsed            int64  `json:"creditsUsed"`
	CreditsAvailable       int64  `json:"creditsAvailable"`
	CreditLimit            int64  `json:"creditLimit"`
	LimitInterval          string `json:"limitInterval"`
	IntervalExpiryTimespan string `json:"intervalExpiryTimespan"`
}

// RespSupportedChain represents a supported chain
type RespSupportedChain struct {
	ChainName     string `json:"chainname"`
	ChainID       string `json:"chainid"`
	BlockExplorer string `json:"blockexplorer"`
	APIURL        string `json:"apiurl"`
	Status        int64  `json:"status"`
}

// RespSupportedChains represents supported chains information
type RespSupportedChains struct {
	TotalCount int64                `json:"totalcount"`
	Result     []RespSupportedChain `json:"result"`
}

// Address Tagging & Labeling Module Response Types

// RespAddressTag represents address tag information
type RespAddressTag struct {
	Address              string   `json:"address"`
	Nametag              string   `json:"nametag"`
	InternalNametag      string   `json:"internal_nametag"`
	URL                  string   `json:"url"`
	ShortDescription     string   `json:"shortdescription"`
	Notes1               string   `json:"notes_1"`
	Notes2               string   `json:"notes_2"`
	Labels               []string `json:"labels"`
	LabelsSlug           []string `json:"labels_slug"`
	Reputation           int64    `json:"reputation"`
	OtherAttributes      []string `json:"other_attributes"`
	LastUpdatedTimestamp int64    `json:"lastupdatedtimestamp"`
}

type RespAddressTags []RespAddressTag

// RespLabelMaster represents label master information
type RespLabelMaster struct {
	LabelName            string `json:"labelname"`
	LabelSlug            string `json:"labelslug"`
	ShortDescription     string `json:"shortdescription"`
	Notes                string `json:"notes"`
	LastUpdatedTimestamp int64  `json:"lastupdatedtimestamp"`
}

type RespLabelMasterlist []RespLabelMaster

// RespLatestCSVBatchNumber represents nametag batch information
type RespLatestCSVBatchNumber struct {
	Nametag              string `json:"nametag"`
	Batch                string `json:"batch"`
	LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
}

type RespLatestCSVBatchNumbers []RespLatestCSVBatchNumber
